import { BrowserWindowConstructorOptions } from "electron";
import { BROWSER_WINDOW_KEY } from "../shared/ipcConst";

export interface ReadyToShowBrowserWindowProps {
    port?: string;
    origin?: string;
    id?: BROWSER_WINDOW_KEY;
    url?: string;
    electronConfig?: MainBrowserWindowConfigProps;
}

export interface MainBrowserWindowConfigProps extends BrowserWindowConstructorOptions {
    /**
     * 图标地址
     */
    icon?: string;
    /**
     * preload加载地址
     */
    preload?: string;
}

export class ElectronBrowserWindowDefaultConfig {
    /**
     * 在electron中只有右和下
     */
    static MAX_RIGHT = 1280;
    static MAX_BOTTOM = 724;
    static START_PAGE_TIME = 3000;
    /**
     * 配置主window
     * @param config
     */
    static mainBrowserWindowConfig = (config?: MainBrowserWindowConfigProps): BrowserWindowConstructorOptions => {
        return {
            show: false,
            width: this.MAX_RIGHT,
            frame: false,
            height: this.MAX_BOTTOM,
            minWidth: this.MAX_RIGHT,
            minHeight: this.MAX_BOTTOM,
            icon: config?.icon,
            title: "推文助手",
            webPreferences: {
                webviewTag: true,
                webSecurity: false,
                nodeIntegration: true,
                plugins: true,
            },
            ...config,
        };
    };
}

