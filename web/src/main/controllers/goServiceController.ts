import { ChildProcessWithoutNullStreams, spawn } from "child_process";
import { PlatformHelper } from "../shared/PlatformHelper";
import { join } from "path";
import { app, net } from "electron";
import fileController from "../controllers/FileController";
import FileController from "../controllers/FileController";
import sudo from "../shared/super";
import { BUILDER_NAME } from "../../../.erb/electron-builder/system-config";
import { resourcePath } from "../shared/pathHelper";
import fs from "fs";
import { error, log } from "../shared/debugLog";
import { BasicDirHelper } from "../shared/basicHelper";
import FileHelper from "../shared/fileHelper";
import { ProjectDirHelper } from "../shared/projectDirHelper";


/**
 * @desc go 服务的可执行文件
 * @author feihan
 * @date 2023/1/12 11:49
 */
class GoServiceController {
    goChildProcess: ChildProcessWithoutNullStreams | undefined;
    serverPort?: string;

    /**
     * 拿到环境变量路径以及转换
     * @param isWindows
     */
    getOasisExecutePath = (isWindows?: boolean): string => {
        const userDataPath = BasicDirHelper.getUserDataPath();
        const path = join(userDataPath.trim(), "server");
        if (isWindows) {
            return BasicDirHelper.transformWindowPath(path);
        }
        return BasicDirHelper.transformMacPath(path);
    };

    /**
     * 启动go
     * @param cmd
     * @param isWindows
     */
    spawnEncapsulation(cmd, isWindows?: boolean) {
        const oasis_execute_path = this.getOasisExecutePath(isWindows);
        const spawnCmd = BasicDirHelper.transformMacWindowPath(cmd, isWindows);
        const upperLayerPath = join(spawnCmd, "..");
        let spawnShell;
        if (isWindows) {
            spawnShell = `cd ${upperLayerPath}&& ${PlatformHelper.getGoServerName()} -execute_path=${oasis_execute_path}`;
        } else {
            spawnShell = `${spawnCmd} -execute_path=${oasis_execute_path}`;
        }
        log({
            "最终启动的命令": spawnShell,
        });
        return spawn(spawnShell, {
            shell: true,
        });
    }

    /**
     * 1.校验go启动程序是否有执行权限如果没有则让用户授权
     */
    startGoProgram(): Promise<string> {
        if (!app.isPackaged) {
            return Promise.resolve("8889");
        }
        const goServicePath = ProjectDirHelper.getGoServicePath();
        return new Promise((resolve, reject) => {
            fileController.staticCheckFilePermissions(goServicePath, {
                fileConstants: fs.constants.X_OK | fs.constants.R_OK,
                constantsSuccessCallback: (success?: boolean) => {
                    this.checkStartServer().then(resolve).catch(reject);
                },
                constantsCatchCallback(error?: string) {
                    log("这边校验到没有权限");
                },
                messageBoxSuccessCallback: (confirm: number) => {
                    if (confirm == 1) {
                        error("用户不允许授权退出程序");
                        app.quit();
                        return;
                    }
                    if (confirm == 0) {
                        // 由于Linux下一定的输入密码才行所以强制让他输密码
                        if (process.platform === "linux") {
                            sudo.exec(`chmod +x ${goServicePath}`, {
                                name: BUILDER_NAME,
                                icns: join(resourcePath, "icon.icns")
                            }, (err) => {
                                error("输入密码后的错误信息", err);
                                if (err) {
                                    app.quit();
                                    return;
                                }
                                this.checkStartServer().then(resolve).catch(reject);
                            });
                        } else {
                            // 这边是mac或者window去授权go程序
                            FileController.sudoAuthorityFile(goServicePath).then(() => {
                                log("启动程序授权成功");
                                this.checkStartServer().then(resolve).catch(reject);
                            }).catch(errorMessage => {
                                error("启动程序授权失败", errorMessage);
                                app.quit();
                            });
                        }
                    }
                },
                messageBoxCatchCallback(errorMessage?: string) {
                    error("授权失败", errorMessage);
                    // 如果拒绝或者错误
                    app.quit();
                }
            });
        });
    }

    /**
     * 启动child_process
     * @param goServicePath
     * @param isWindow
     */
    checkChildProcess(goServicePath: string, isWindow?: boolean): Promise<string> {
        return new Promise((resolve, reject) => {
            const goPortPath = BasicDirHelper.getClientBuilderResourceFilePath("port.json");
            log({
                "后端启动地址": `chmod +x ${goServicePath}`,
                "后端端口地址": goPortPath,
            });
            const child_process = this.spawnEncapsulation(goServicePath, isWindow);
            log(child_process.pid, "进程pid");
            this.goChildProcess = child_process;
            child_process.stdout?.on("error", (data) => {
                this.structReset();
                reject(data);
                error("程序启动失败", data);
            });
            child_process.stdout?.on("data", (data) => {
                const reg = "Using port:(.*?)\\.\\.\\.";
                const res = data.toString().trim().match(reg);
                if (res?.[1]) {
                    FileHelper.createFile(goPortPath, JSON.stringify({
                        port: res?.[1]?.trim()
                    }));
                    this.serverPort = res?.[1]?.trim();
                    resolve(res?.[1]?.trim());
                }
                log(data.toString().trim());
            });
            child_process.on("close", (code, data) => {
                this.structReset();
                error("进程被关闭", code, data);
                reject(data);
            });
        });
    }

    checkStartServer(): Promise<string> {
        return new Promise((resolve, reject) => {
            const goServicePath = ProjectDirHelper.getGoServicePath();
            const { isWindow } = PlatformHelper.validatePlatform();
            this.checkChildProcess(goServicePath, isWindow).then(resolve).catch(reject);
        });
    }

    structReset() {
        this.goChildProcess = undefined;
        this.serverPort = undefined;
    }

    endGoProgram() {
        if (this.goChildProcess && this.serverPort) {
            const goChildProcess = this.goChildProcess;
            const port = this.serverPort;
            net.fetch(`http://127.0.0.1:${port}/basic/exit`);
            goChildProcess.stdin.end(() => {
                log("终止程序");
            });
            this.structReset();
        }
    }

    /**
     * 终止服务
     */
    terminateGoProgram() {
        // const name = ProjectDirHelper.getGoServiceName();
        // if (!name) {
        //     error("没有支持的shell命令");
        //     return null;
        // }
        // const pids = ProcessHelper.getProgramPid(name);
        // if (pids) {
        //     ProcessHelper.killProgramByPid(pids);
        // } else {
        //     error("进程id不合法");
        // }
    }
}

export default GoServiceController;
