import { BrowserWindow, dialog, OpenDialogReturnValue, shell } from "electron";
import fs from "fs";
import { error, info, log } from "../shared/debugLog";
import { FileFolderOptions } from "../../renderer/ipc/types";
import { exec } from "child_process";
import { BasicDirHelper } from "../shared/basicHelper";

class FileController {

    folderSelect(
        manWindow: BrowserWindow,
        options: FileFolderOptions,
    ): Promise<OpenDialogReturnValue> {
        return new Promise((resolve, reject) => {
            dialog
                .showOpenDialog(manWindow, options)
                .then(value => {
                    log("当前选中的文件目录", value);
                    resolve(value);
                })
                .catch(e => {
                    error("当前选中的文件目录", e);
                    reject(e);
                });
        });
    }

    /**
     * 创建文件夹
     * @param filePath
     */
    createFolder(filePath: string) {
        return new Promise((resolve, reject) => {
            try {
                if (!fs.existsSync(filePath)) {
                    fs.mkdirSync(filePath, { recursive: true });
                    log("创建文件夹成功", filePath);
                    resolve(true);
                } else {
                    log("创建文件夹成功1", filePath);
                    resolve(true);
                }
            } catch (e) {
                reject(`创建目录失败${e?.toString()}`);
            }
        });
    }

    /**
     * 异步读取文件
     * @param filePath
     * @param options
     */
    readFile(filePath, options?: { encoding: BufferEncoding, flag?: string | undefined } | BufferEncoding): Promise<string> {
        return new Promise((resolve, reject) => {
            if (!fs.existsSync(filePath)) {
                reject(null);
            } else {
                fs.readFile(filePath, options ?? "utf8", (error, data) => {
                    if (error) {
                        reject(error);
                    } else {
                        resolve(data);
                    }
                });
            }
        });
    }

    /**
     * 同步读取文件
     * @param filePath
     * @param options
     */
    readFileSync(filePath, options?: { encoding: BufferEncoding, flag?: string | undefined } | BufferEncoding) {
        const newOptions = options ?? "utf8";
        return fs.readFileSync(filePath, newOptions);
    }

    /**
     * 打开文件所在目录
     * @param filePath
     */
    openFilePath(filePath): Promise<string> {
        return new Promise<string>((resolve, reject) => {
            this.checkWritePermissions(filePath).then(() => {
                shell.showItemInFolder(filePath);
                resolve("打开成功");
            }).catch(reject);
        });
    }

    /**
     * 校验文件是否存在写操作
     * @param folderPath
     * @param constants
     */
    async checkWritePermissions(folderPath: string, constants: number = fs.constants.F_OK): Promise<boolean> {
        return new Promise((resolve, reject) => {
            try {
                fs.access(folderPath, constants, errorMessage => {
                    if (errorMessage) {
                        error(`读取文件信息失败信息: ${constants} ${errorMessage}`);
                        reject(errorMessage);
                    } else {
                        log(`读取文件权限成功信息: ${constants}`, folderPath);
                        // 有写入权限
                        resolve(true);
                    }
                });
            } catch (errorMessage) {
                error(`读取失败信息:${constants} ${errorMessage}`);
                reject(errorMessage);
            }
        });
    }

    /**
     * 授权文件
     * @param directory
     */
    async execAuthorityFile(directory: string) {
        return new Promise((resolve, reject) => {
            const transformPath = BasicDirHelper.transformMacWindowPath(directory, process.platform == "win32");
            exec(`chmod +x ${transformPath}`, (errorMessage) => {
                if (errorMessage) {
                    error("授权失败信息", errorMessage, `chmod +x ${transformPath}`);
                    reject(errorMessage);
                    return;
                }
                log("授权成功信息", `chmod +x ${transformPath}`);
                resolve(true);
            });
        });
    }

    /**
     * 校验文件
     * @param filePath
     * @param config
     * 1. 校验文件是否存在，并且是否存在读写等权限
     * 2. 如无权限则调用原生dialog
     * 3. 通过用户选择去进行操作
     */
    staticCheckFilePermissions(filePath: string, config?: {
        /**
         * fs.constants 默认为是否可执行、是否可写、是否存在
         */
        fileConstants?: number;
        /**
         * 提示消息
         */
        messageBoxTitle?: string;
        /**
         * 提示type
         */
        messageBoxType?: Electron.MessageBoxOptions["type"];
        /**
         * 校验文件成功
         * @param success
         */
        constantsSuccessCallback?(success?: boolean): void;
        /**
         * 校验文件失败
         * @param error
         */
        constantsCatchCallback?(error?: string): void;
        /**
         * 弹窗事件
         * @param confirm
         * @param 1 拒绝
         * @param 0 同意
         */
        messageBoxSuccessCallback?(confirm: number): void;
        /**
         * 弹窗错误处理
         * @param error
         */
        messageBoxCatchCallback?(error?: string): void;
    }) {
        if (!filePath) {
            error("当前文件路径不存在");
            throw Error("当前文件路径不存在");
        }
        this.checkWritePermissions(filePath, config?.fileConstants ?? fs.constants.X_OK | fs.constants.R_OK | fs.constants.F_OK).then((success) => {
            config?.constantsSuccessCallback?.(success);
        }).catch((err) => {
            config?.constantsCatchCallback?.();
            error("当前启动程序没有执行权限", err);
            // 没有权限的话让用户授权
            dialog.showMessageBox({
                type: config?.messageBoxType ?? "warning",
                message: config?.messageBoxTitle ?? "软件正常运行需要同意相关授权,是否同意授权?",
                buttons: [
                    "是",
                    "否"
                ],
                defaultId: 0,
                cancelId: 1,
            }).then((response) => {
                config?.messageBoxSuccessCallback?.(response.response);
            }).catch((e) => {
                error("授权失败", e);
                config?.messageBoxCatchCallback?.(e?.toString());
                // 如果拒绝或者错误
            });
        });
    }

    /**
     * 授权文件
     * @param directory
     */
    async sudoAuthorityFile(directory: string) {
        return new Promise((resolve, reject) => {
            // 首先去校验当前文件是否存在
            this.checkWritePermissions(directory, fs.constants.F_OK).then(() => {
                // 如果文件存在则去校验是否存在可执行标志
                this.checkWritePermissions(directory, fs.constants.X_OK).then(() => {
                    // 如果存在并且存在执行标志
                    info("当前存在执行标志的文件路径", directory);
                    resolve(true);
                }).catch(() => {
                    this.execAuthorityFile(directory).then(resolve).catch(reject);
                });
            }).catch((errorMessage) => {
                error("文件读取失败信息", errorMessage, "当前文件不存在");
                reject(errorMessage);
            });
        });
    }
}

export default new FileController();
