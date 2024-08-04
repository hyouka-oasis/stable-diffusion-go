import { OpenDialogReturnValue } from "electron";
import { BasicRendererIpcTransferStationAdapter } from "renderer/ipc/BasicRendererIpcTransferStationAdapter";
import { BasicRendererIpcPreload } from "main/ipc/types";
import { IpcConst } from "main/shared/ipcConst";
import { FileFolderOptions } from "renderer/ipc/types";

class FileIpcAdapterImpl extends BasicRendererIpcTransferStationAdapter {

    constructor() {
        super();
    }

    onFolderSelect(options: FileFolderOptions): Promise<BasicRendererIpcPreload<OpenDialogReturnValue>> {
        return new Promise(
            (resolve, reject) => {
                this.ipcSendMessage(resolve, reject, IpcConst.FOLDER_READ, options);
            },
        );
    }

    onReadFile(filePath: string, options?: { encoding: BufferEncoding, flag?: string | undefined } | BufferEncoding): Promise<BasicRendererIpcPreload<string>> {
        return new Promise(
            (resolve, reject) => {
                this.ipcSendMessage(resolve, reject, IpcConst.FILE_READ, { filePath, options });
            },
        );
    }

    onOpenFilePath(filePath: string): Promise<BasicRendererIpcPreload<string>> {
        return new Promise(
            (resolve, reject) => {
                this.ipcSendMessage(resolve, reject, IpcConst.FILE_OPEN, filePath);
            },
        );
    }
}

export {
    FileIpcAdapterImpl,
};
