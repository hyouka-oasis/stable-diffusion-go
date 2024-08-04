import { BrowserWindow, Rectangle } from "electron";
import { PlatformHelper } from "../shared/PlatformHelper";

class BrowserWindowController {
    /**
     * 关闭窗口
     * @param browserWindow
     */
    closeBrowserWindow(browserWindow: BrowserWindow | null | undefined) {
        browserWindow?.close();
    }

    /**
     * 最大化
     * @param browserWindow
     */
    setBrowserWindowFullScreen(browserWindow: BrowserWindow | null | undefined) {
        const { isWindow } = PlatformHelper.validatePlatform();
        if (isWindow) {
            if (browserWindow?.isMaximized()) {
                browserWindow?.unmaximize();
            } else {
                browserWindow?.maximize();
            }
        } else {
            browserWindow?.setFullScreen(!browserWindow?.isFullScreen());
        }
    }

    /**
     * 最小化
     * @param browserWindow
     */
    setBrowserWindowMinimize(browserWindow: BrowserWindow | null | undefined) {
        browserWindow?.minimize();
    }

    /**
     * 是否允许放大放小
     * @param resizable
     * @param browserWindow
     */
    setBrowserWindowResizable(resizable: boolean, browserWindow: BrowserWindow | null | undefined) {
        browserWindow?.setResizable(resizable);
    }

    /**
     * 设置窗体属性
     * @param bounds
     * @param browserWindow
     */
    setBrowserWindowBounds(bounds: Partial<Rectangle> & { boundsAnimation?: boolean }, browserWindow: BrowserWindow | null | undefined) {
        browserWindow?.setBounds(bounds);
    }

    /**
     * 设置是否允许最大化
     * @param maximizable
     * @param browserWindow
     */
    setBrowserWindowMaximizable(maximizable: boolean, browserWindow: BrowserWindow | null | undefined) {
        browserWindow?.setMaximizable(maximizable);
    }

    /**
     * 设置是否居中
     * @param browserWindow
     */
    setBrowserWindowCenter(browserWindow: BrowserWindow | null | undefined) {
        browserWindow?.center();
    }

    /**
     * 设置是否显示隐藏
     * @param browserWindow
     */
    setBrowserWindowVisible(browserWindow: BrowserWindow | null | undefined, visible: boolean) {
        if (visible) {
            browserWindow?.show();
        } else {
            browserWindow?.hide();
        }
    }

    /**
     * 设置聚焦
     * @param browserWindow
     * @param focus
     */
    setBrowserWindowFocus(browserWindow: BrowserWindow | null | undefined, focus: boolean) {
        if (focus) {
            browserWindow?.focus();
        } else {
            browserWindow?.blur();
        }
    }
}

export default new BrowserWindowController();
