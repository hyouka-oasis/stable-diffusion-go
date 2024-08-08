import path from "path";
import { PlatformHelper } from "./PlatformHelper";
import { BasicDirHelper, isDevelopment } from "./BasicHelper";
import { app } from "electron";
import FileController from "../../main/controllers/FileController";

export class ProjectDirHelper {

    /**
     * @desc 获取go服务目录
     * @author feihan
     * @date 2023/1/13 9:17
     */
    static getGoServicePath(): string {
        const { isWindow, isLinux, isMac } = PlatformHelper.validatePlatform();
        let suffix = "";
        const goServerName = PlatformHelper.getGoServerName();
        if (!goServerName) {
            throw Error("不存在当前系统");
        }
        if (isMac) {
            suffix = `${BasicDirHelper.getDevelopmentGoServerPath("mac")}/${goServerName}`;
        }
        if (isWindow) {
            suffix = `${BasicDirHelper.getDevelopmentGoServerPath("win")}/${goServerName}`;
        }
        if (isLinux) {
            suffix = `${BasicDirHelper.getDevelopmentGoServerPath("linux")}/${goServerName}`;
        }
        if (!suffix) {
            return suffix;
        }
        console.log(suffix, "suffix");
        const filePath = path.join(suffix);
        return isWindow ? filePath.replace(/\//g, "//") : filePath;
    }

    /**
     * 获取应用数据存放目录
     * */
    static getAppDataDir(): string {
        // Todo 不同系统存放位置不一样，举例：
        //  windows C:\Users\fh\AppData\Local\Temp\animate-client-data
        //  linux 待定
        return "";
    }

    /**
     * 废弃该接口的实现，logs目录直接使用getAppLogsDir获取即可
     * 设置应用操作日志的存放目录
     * */
    static setAppLogsDir() {
        const { devRootDir, prodRootDir } = BasicDirHelper.getRootDir();
        const devDir = path.join(devRootDir, "logs");
        const prodDir = path.join(prodRootDir, "logs");
        // 开发环境和生产环境下 日志目录均是 C:\Users\fh\AppData\Roaming\Electron\logs，与devDir、prodDir的值不符
        console.log("devDir ", devDir, "prodDir ", prodDir);
        const dir = isDevelopment ? devDir : prodDir;
        // app.setPath("logs", dir);
        app.setAppLogsPath(dir);
        // 获取logs
        // app.getPath("logs");
    }

    /**
     * 获取应用操作日志的存放目录
     * */
    static getAppLogsDir(): string {
        const { devRootDir, prodRootDir } = BasicDirHelper.getRootDir();
        const devDir = path.join(devRootDir, "logs");
        const prodDir = path.join(prodRootDir, "logs");
        return isDevelopment ? devDir : prodDir;
    }

    /**
     * 获取数据处理中间件服务的日志存放目录
     * */
    static async getDataProcessCenterLogDir(): Promise<string> {
        const userDataPath = BasicDirHelper.getUserDataPath();
        const processLogPath = path.join(userDataPath, "oasis-server/log/process-center-log");
        await FileController.createFolder(processLogPath);
        return processLogPath;
    }
}
