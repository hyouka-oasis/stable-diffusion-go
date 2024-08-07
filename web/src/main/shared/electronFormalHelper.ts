import { app, BrowserWindow, dialog } from "electron";
import { ProcessHelper } from "../shared/processHelper";
import { error, log } from "../shared/debugLog";

/**
 * 不允许打开devtools
 * @param mainWindow
 */
export const closeDevTools = (mainWindow: BrowserWindow | null) => {
    if (!mainWindow) {
        return;
    }
    if (process.env.NODE_ENV === "development") {
        return;
    }
    if (process.env.CLOSE_DEVTOOLS === "true") {
        return;
    }
    mainWindow?.webContents?.on("devtools-opened", () => {
        // 关闭开发者工具控制台
        mainWindow?.webContents?.closeDevTools();
    });
};
/**
 * 关闭客户端的时候进行的操作
 * @param mainWindow
 */
export const closedBrowserWindowHandler = (mainWindow: BrowserWindow | null) => {
    if (!mainWindow) {
        return;
    }
    mainWindow?.on("closed", () => {
        mainWindow = null;
    });
};

/**
 * browserWindow准备就绪的动作
 * @param mainWindow
 * @param callback
 */
export const readyToShowBrowserWindowHandler = (mainWindow: BrowserWindow | null, callback?: (mainWindow: BrowserWindow | null) => void) => {
    if (!mainWindow) {
        return;
    }
    if (!mainWindow) {
        throw new Error("\"mainWindow\" is not defined");
    }
    mainWindow?.on("ready-to-show", () => {
        if (process.env.START_MINIMIZED) {
            mainWindow.minimize();
        } else {
            // mac全屏聚焦太难受了
            if (process.platform === "darwin") {
                mainWindow.showInactive();
            } else {
                mainWindow.show();
            }
        }
        callback?.(mainWindow);
    });
};

interface KillRelaunchServerServer {
    /**
     * 展示图标
     */
    icon?: string;
}

/**
 * 异常状态杀死
 */
export const killRelaunchServer = (options?: KillRelaunchServerServer) => {
    dialog.showMessageBox({
        message: "客户端启动失败，请尝试重新启动客户端。",
        buttons: [ "退出程序", "重新启动" ],
        icon: options?.icon,
    }).then(response => {
        const pids = ProcessHelper.getOasisServerPids();
        if (!pids) {
            error("关闭进程失败", pids);
        } else {
            ProcessHelper.killProgramByPid(pids);
            log("拿到的所有pid进程", pids);
        }
        if (response?.response === 0) {
            app.quit();
        }
        if (response?.response === 1) {
            app.relaunch();
            app.quit();
        }
    }).catch(err => {
        app.quit();
        error(err, "启动错误后用户手动行为");
    });
};
