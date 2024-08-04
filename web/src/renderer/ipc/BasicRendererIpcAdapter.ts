import { BasicRendererStaticIpcAdapter } from "renderer/ipc/BasicRendererStaticIpcAdapter";
import { WindowIpcAdapterImpl } from "renderer/ipc/basic/windowIpcAdapterImpl";
import { FileIpcAdapterImpl } from "renderer/ipc/basic/fileIpcAdapterImpl";

export interface IpcApi {
    windowAdapter: WindowIpcAdapterImpl;
    fileAdapter: FileIpcAdapterImpl;
}

/**
 * BasicRendererStaticIpcAdapter请勿多次继承
 */
class BasicRendererIpcAdapter extends BasicRendererStaticIpcAdapter {
    ipcApi: IpcApi;

    constructor() {
        super();
        this.ipcApi = {
            windowAdapter: new WindowIpcAdapterImpl(),
            fileAdapter: new FileIpcAdapterImpl(),
        };
    }

    getIpcApi() {
        return this.ipcApi;
    }
}


export const ipcApi = new BasicRendererIpcAdapter().getIpcApi();
