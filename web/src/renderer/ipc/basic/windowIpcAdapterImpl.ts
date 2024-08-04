import { BasicRendererIpcTransferStationAdapter } from "renderer/ipc/BasicRendererIpcTransferStationAdapter";
import { BROWSER_WINDOW_KEY, IpcConst } from "main/shared/ipcConst";
import { BasicRendererIpcPreload } from "main/ipc/types";
import { BasicElectronWindowBoundsProps, BasicElectronWindowConfigType } from "renderer/ipc/types";


class WindowIpcAdapterImpl extends BasicRendererIpcTransferStationAdapter {
    maxElectronWindow() {
        return new Promise<BasicElectronWindowBoundsProps>((resolve, reject) => {
            this.ipcSendMessage(resolve, reject, IpcConst.MAIN_BROWSER_WINDOW_SETTING, {
                windowControllerType: "max",
            });
        });
    }

    minElectronWindow() {
        return new Promise<BasicElectronWindowBoundsProps>((resolve, reject) => {
            this.ipcSendMessage(resolve, reject, IpcConst.MAIN_BROWSER_WINDOW_SETTING, {
                windowControllerType: "min",
            });
        });
    }

    setElectronWindowConfig(config: BasicElectronWindowBoundsProps) {
        return new Promise<BasicElectronWindowBoundsProps>((resolve, reject) => {
            this.ipcSendMessage(resolve, reject, IpcConst.MAIN_BROWSER_WINDOW_SETTING, config);
        });
    }

    getElectronWindowConfig(type: BasicElectronWindowConfigType, key?: BROWSER_WINDOW_KEY) {
        return new Promise<BasicRendererIpcPreload<boolean>>((resolve, reject) => {
            this.ipcSendMessage(resolve, reject, IpcConst.MAIN_BROWSER_GET_WINDOW_SETTING, { type, key });
        });
    }

    /**
     *  删除browser
     * @param id
     */
    closeBrowserWindow(id: BROWSER_WINDOW_KEY | BROWSER_WINDOW_KEY[]): Promise<BasicRendererIpcPreload<null>> {
        return new Promise((resolve, reject) => {
            this.ipcSendMessage(resolve, reject, IpcConst.CLOSE_NEW_BROWSER_WINDOW, id);
        });
    }

    /**
     * 关闭所有窗口
     */
    closeAllBrowserWindow(): Promise<BasicRendererIpcPreload<null>> {
        return new Promise((resolve, reject) => {
            this.ipcSendMessage(resolve, reject, IpcConst.CLOSE_ALL_BROWSER_WINDOW);
        });
    }
}

export {
    WindowIpcAdapterImpl,
};
