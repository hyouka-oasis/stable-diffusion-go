import { BROWSER_WINDOW_KEY } from "main/shared/ipcConst";
import { OpenDialogOptions } from "electron";

export interface BasicElectronWindowBoundsProps {
    /**
     * 窗口实体名
     */
    browserWindowKey?: BROWSER_WINDOW_KEY;
    windowControllerType?: "close" | "min" | "max";
    resizable?: boolean;
    bounds?: {
        x?: number;
        y?: number;
        width?: number;
        height?: number;
    };
    boundsAnimation?: boolean;
    maximizable?: boolean;
    center?: boolean;
    hide?: boolean;
    show?: boolean;
    focus?: boolean;
}

/**
 * 具体参数参考一下链接
 * @link https://www.electronjs.org/zh/docs/latest/api/browser-window#wingetnormalbounds
 */
export type BasicElectronWindowConfigType =
    | "getBounds"
    | "getNormalBounds"
    | "isEnabled"
    | "getMaximumSize"
    | "getMinimumSize"
    | "isFullScreenable"
    | "isMaximized"
    | "isFullScreen";


export interface FileFolderOptions extends OpenDialogOptions {
}

export interface FileFolderCreateOptions {
    /**
     * 文件夹路径
     */
    folderPath: string;
    isFilePath?: boolean;
}
