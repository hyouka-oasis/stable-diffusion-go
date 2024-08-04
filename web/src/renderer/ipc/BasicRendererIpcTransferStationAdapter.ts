import { BasicRendererStaticIpcAdapter } from "renderer/ipc/BasicRendererStaticIpcAdapter";
import { BasicIpcConst, IpcConst } from "main/shared/ipcConst";

const { ipcRenderer } = window.electron ?? {};

class BasicRendererIpcTransferStationAdapter {
    getPrimaryKey() {
        return BasicRendererStaticIpcAdapter.getPrimaryKey();
    }

    /**
     * DATA, RESOLVE, REJECT 会自动依赖Promise类型
     * @param resolve
     * @param reject
     * @param eventKey
     * @param data
     */
    ipcSendMessage<DATA, RESOLVE = unknown, REJECT = unknown>(
        resolve: (value: RESOLVE) => void,
        reject: (value: REJECT) => void,
        eventKey: IpcConst,
        data?: DATA,
    ) {
        const primaryKey = this.getPrimaryKey();
        BasicRendererStaticIpcAdapter.setIpcRendererResolverMap(primaryKey, resolve);
        BasicRendererStaticIpcAdapter.setIpcRendererRejectMap(primaryKey, reject);
        ipcRenderer?.sendMessage<DATA>(BasicIpcConst.BASIC_MAIN_IPC_KEY, {
            key: eventKey,
            data: data as DATA,
            id: primaryKey,
        });
    }
}

export {
    BasicRendererIpcTransferStationAdapter,
};
