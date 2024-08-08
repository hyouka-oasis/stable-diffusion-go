import path, { join } from "path";
import { app } from "electron";
import { PlatformHelper } from "./PlatformHelper";
import fileController from "../controllers/FileController";
import { error } from "./debugLog";

// 应用打包后import.meta.env.DEV为true，正常应该为false，故使用!app.isPackaged替代
// export const isDevelopment = import.meta.env.DEV;
export const isDevelopment = !app.isPackaged;

export class BasicDirHelper {
    /**
     * 应用程序目录
     */
    static getAppPath() {
        return app.getAppPath();
    }

    /**
     * @desc 获取项目根目录
     * @author feihan
     * @date 2023/1/13 9:17
     */
    static getRootDir() {
        // appPath表示当前应用的目录，其值如下：
        // 开发环境：格式：项目代码存放目录logs； exp: D:\work\reality3deditorclient
        // BasicDirHelper.getUserDataPath表示当前缓存存放目录
        // 生产环境：与缓存是同一目录； exp: C:\Users\[user name]\AppData\Roaming\[客户端名称]
        /**
         * 在macOs当中 打包成app后如果electron-build配置的是extraResources那么在打包完成后所拷贝的resources文件在对应的包/Content/Resources/下
         * 如果使用的是extraFiles则在包/Content/MacOS/下
         */
        let appPath: string = BasicDirHelper.getUserDataPath();
        if (isDevelopment) {
            appPath = app.getAppPath();
        }
        const devRootDir = appPath;
        const prodRootDir = appPath;
        return { devRootDir, prodRootDir };
    }

    /**
     * 应用程序缓存目录
     */
    static getUserDataPath() {
        return app.getPath("userData");
    }

    /**
     * 应用程序缓存目录
     */
    static getAppDataPath() {
        return app.getPath("appData");
    }


    static getDevelopmentGoServerPath(name: string): string {
        /**
         * todo 后续目录需要修改 注意
         */
        const { isWindow, isLinux, isMac } = PlatformHelper.validatePlatform();
        let suffix = "";
        if (process.env.NODE_ENV === "development") {
            return path.join(process.cwd(), "../", "server/tmp/");
        }
        if (isMac || isLinux) {
            suffix = `${path.join(BasicDirHelper.getAppPath(), "../", "oasis-server")}`;
        }
        if (isWindow) {
            suffix = path.join(process.cwd(), "./resources", "oasis-server");
        }
        if (!suffix) {
            return suffix;
        }
        return suffix;
    }

    /**
     * 拿到指定路径
     */
    static getClientBuilderResourceFilePath(name?: string): string {
        let suffix = "";
        // 如果是开发模式则拿到当前工作目录下的port;
        if (process.env.NODE_ENV === "development") {
            return `${path.join(BasicDirHelper.getAppPath(), name ?? "")}`;
        }
        // mac下通过二进制文件启动的话是在~/Desktop/../../.app/Content/
        // 如果是通过app启动的话则是在.asar内
        // 目前Linux还不确定是否为这样
        suffix = path.join(BasicDirHelper.getAppPath(), "../", "oasis-server", name ? `/${name}` : "");
        if (!suffix) {
            // 获取不到port的话不应该启动
            throw Error("程序不能正常启动, 请联系工作人员");
        }
        return suffix;
    }

    /**
     * 获取数据处理中间件程序地址
     * @param binPathName
     */
    static getDataProcessCenterServerPath(binPathName: string): string {
        const { isWindow, isLinux, isMac } = PlatformHelper.validatePlatform();
        let suffix = "";
        const serverName = binPathName.replace("bin/", "").replace(".exe", "");
        if (process.env.NODE_ENV === "development") {
            if (isMac) {
                suffix = path.join(process.cwd(), `node_modules/@oasis/${serverName}/${binPathName}`);
            }
            if (isLinux) {
                suffix = path.join(process.cwd(), `node_modules/@oasis/${serverName}/${binPathName}`);
            }
            if (isWindow) {
                suffix = path.join(process.cwd(), `node_modules/@oasis/${serverName}/${binPathName}`);
            }
            return suffix;
        }
        if (isMac || isLinux) {
            suffix = `${path.join(BasicDirHelper.getAppPath(), "../", `oasis-server/bin/${serverName}`)}`;
        }
        if (isWindow) {
            suffix = path.join(process.cwd(), "./resources", `oasis-server/bin/${serverName}.exe`);
        }
        if (!suffix) {
            return suffix;
        }
        return suffix;
    }

    /**
     * 转换Linux或者window路径
     * @param path
     * @param isWindows
     */
    static transformMacWindowPath(path?: string, isWindows?: boolean) {
        if (!path || !path?.trim()?.length) {
            return "";
        }
        if (isWindows) {
            return this.transformWindowPath(path);
        }
        return this.transformMacPath(path);
    }

    /**
     * 转换为window cmd适用的格式
     * @param path
     */
    static transformWindowPath = (path: string) => {
        if (!path || !path?.trim()?.length) {
            return "";
        }
        return path.replace(" ", "^ ");
    };

    /**
     * 转换为window cmd适用的格式
     * @param path
     */
    static transformMacPath = (path: string) => {
        if (!path || !path?.trim()?.length) {
            return "";
        }
        return path.replace(" ", "\\ ");
    };

    /**
     * 创建主题目录
     */
    static initThemeFolder() {
        const appPath = BasicDirHelper.getUserDataPath();
        const themePath = join(appPath, "theme");
        fileController.createFolder(themePath).catch(e => {
            error("创建文件夹失败", e);
        });
    }
}
