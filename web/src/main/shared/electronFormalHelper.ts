import { BrowserWindow } from "electron";

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
