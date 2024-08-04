import { ipcMain } from "electron";
import { BasicIpcConst, BasicIpcKeys } from "../shared/ipcConst";
import { BasicIpcResponse } from "../ipc/basicIpcResponse";
import { error } from "../shared/debugLog";
import { BasicMainIpcPreload } from "../ipc/types";
import { windowIpc } from "../ipc/windowIpc";
import { browserWindowIpc } from "../ipc/browserWindowIpc";
import { fileIpc } from "../ipc/fileIpc";


class BasicIpc extends BasicIpcResponse {
    private readonly _staticIpc;

    constructor() {
        super();
        this.init();
        this._staticIpc = {
            ...windowIpc,
            ...browserWindowIpc,
            ...fileIpc
        };
    }

    init() {
        ipcMain.on(BasicIpcConst.BASIC_MAIN_IPC_KEY, async (event, preload: BasicMainIpcPreload) => {
            if (!preload.key) {
                error("当前不存在preloadKey");
                return;
            }
            if (!BasicIpcKeys[preload.key]) {
                error(`当前传入key:${preload.key},不在指定范围内`);
                return;
            }
            if (!this._staticIpc[preload.key]) {
                error(`当前传入key:${preload.key},没有实例方法`);
                return;
            }
            const { reply } = event;
            try {
                const ipcResponse = await this._staticIpc[preload.key](preload.data);
                reply(BasicIpcConst.BASIC_RENDERER_IPC_KEY, this.success(ipcResponse, preload.id));
            } catch (e) {
                error(e, "ipc通讯出错");
                reply(BasicIpcConst.BASIC_RENDERER_IPC_KEY, this.fail(null, preload.id, e?.toString()));
            }
        });
    }

    replace() {
        ipcMain.removeAllListeners();
    }
}

export {
    BasicIpc,
};
