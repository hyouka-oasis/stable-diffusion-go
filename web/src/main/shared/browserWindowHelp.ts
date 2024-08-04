import { BrowserWindow, BrowserWindowConstructorOptions } from "electron";
import { preloadUrl } from "../shared/pathHelper";
import { BROWSER_WINDOW_KEY } from "../shared/ipcConst";

class BrowserWindowHelp {
    private browserWindowMap: Map<BROWSER_WINDOW_KEY, BrowserWindow | null> = new Map();

    /**
     * 创建browserWindows
     * @param id
     * @param config
     */
    initBrowserWindow(id: BROWSER_WINDOW_KEY, config?: BrowserWindowConstructorOptions & {
        callback?(): void;
    }): BrowserWindow | null {
        if (!id) {
            console.error("必须传入browserId才可创建");
            return null;
        }
        const newBrowserWindow = new BrowserWindow({
            ...config,
            webPreferences: {
                ...config?.webPreferences,
                preload: config?.webPreferences?.preload ? config?.webPreferences?.preload : preloadUrl
            }
        });
        this.setBrowserWindow(id, newBrowserWindow);
        config?.callback?.();
        return newBrowserWindow;
    }

    /**
     * 设置browserWindows
     * @param id
     * @param browserWindow
     */
    setBrowserWindow(id: BROWSER_WINDOW_KEY, browserWindow: BrowserWindow | null) {
        this.browserWindowMap.set(id, browserWindow);
    }

    /**
     * 获取browserWindows
     * @param id
     */
    getBrowserWindow(id: BROWSER_WINDOW_KEY): BrowserWindow | null | undefined {
        return this.browserWindowMap.get(id);
    }

    /**
     * 获取所有browserWindows
     */
    getAllBrowserWindow(): Map<BROWSER_WINDOW_KEY, BrowserWindow | null> {
        return this.browserWindowMap;
    }

    /**
     * 删除browser以及ipc
     * @param id
     */
    deleteBrowserWindow(id: BROWSER_WINDOW_KEY) {
        const browserWindow = this.getBrowserWindow(id);
        if (browserWindow) {
            browserWindow.destroy();
            this.setBrowserWindow(id, null);
            this.browserWindowMap.delete(id);
        }
    }

    /**
     * 置空browser
     */
    destroyBrowserWindow() {
        this.browserWindowMap.forEach((browser, id) => {
            this.deleteBrowserWindow(id);
        });
        this.browserWindowMap.clear();
    }
}

export default new BrowserWindowHelp();
