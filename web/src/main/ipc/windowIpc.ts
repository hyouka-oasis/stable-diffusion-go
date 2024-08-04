/**
 * @author hyouka
 * @description 窗口通讯
 */
import { BrowserWindow } from "electron";
import { BasicElectronWindowBoundsProps, BasicElectronWindowConfigType } from "../../renderer/ipc/types";
import browserWindowController from "../controllers/browserWindowController";
import { BROWSER_WINDOW_KEY, IpcConst } from "../shared/ipcConst";
import browserWindowHelp from "../shared/browserWindowHelp";

const WINDOW_SETTING_HANDLER = (config: BasicElectronWindowBoundsProps, browserWindow: BrowserWindow | null | undefined) => {
    const windowControllerType = config.windowControllerType;
    const resizable = config.resizable;
    const bounds = config.bounds;
    const boundsAnimation = config.boundsAnimation;
    const maximizable = config.maximizable;
    const center = config.center;
    const hide = config.hide;
    const show = config.show;
    const focus = config.focus;
    if (windowControllerType !== undefined) {
        if (windowControllerType === "close") {
            browserWindowController.closeBrowserWindow(browserWindow);
        } else if (windowControllerType === "max") {
            browserWindowController?.setBrowserWindowFullScreen(browserWindow);
        } else if (windowControllerType === "min") {
            browserWindowController?.setBrowserWindowMinimize(browserWindow);
        }
    }
    if (resizable !== undefined) {
        browserWindowController?.setBrowserWindowResizable(resizable, browserWindow);
    }
    if (bounds !== undefined) {
        browserWindowController?.setBrowserWindowBounds({ ...bounds, boundsAnimation }, browserWindow);
    }
    if (maximizable !== undefined) {
        browserWindowController?.setBrowserWindowMaximizable(maximizable, browserWindow);
    }
    if (center !== undefined) {
        browserWindowController?.setBrowserWindowCenter(browserWindow);
    }
    if (hide !== undefined) {
        browserWindowController?.setBrowserWindowVisible(browserWindow, hide);
    }
    if (show !== undefined) {
        browserWindowController?.setBrowserWindowVisible(browserWindow, show);
    }
    if (focus !== undefined) {
        browserWindowController?.setBrowserWindowFocus(browserWindow, focus);
    }
};

export const windowIpc = {
    [IpcConst.MAIN_BROWSER_WINDOW_SETTING](config: BasicElectronWindowBoundsProps) {
        const browserWindowKey = config?.browserWindowKey ?? BROWSER_WINDOW_KEY.MAIN_BROWSER;
        const mainBrowserWindow = browserWindowHelp.getBrowserWindow(browserWindowKey);
        if (!mainBrowserWindow) {
            return new Promise((resolve, reject) => {
                reject(`mainBrowserWindow获取异常, key:${browserWindowKey}`);
            });
        }
        WINDOW_SETTING_HANDLER(config, mainBrowserWindow);
        return new Promise((resolve) => {
            resolve("成功");
        });
    },
    [IpcConst.MAIN_BROWSER_GET_WINDOW_SETTING](config: { type: BasicElectronWindowConfigType, key: BROWSER_WINDOW_KEY }) {
        const mainBrowserWindow = browserWindowHelp.getBrowserWindow(config?.key ?? BROWSER_WINDOW_KEY.MAIN_BROWSER);
        return new Promise((resolve, reject) => {
            let data;
            if (mainBrowserWindow && mainBrowserWindow[config.type]) {
                data = mainBrowserWindow?.[config.type]();
                resolve(data);
            } else {
                reject({});
            }
        });
    },
};
