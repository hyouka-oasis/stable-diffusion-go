import { ipcMain, IpcMainEvent } from 'electron';
import { BasicRendererIpcPreload } from "./types";
import { IpcConst } from "../shared/ipcConst";

abstract class IpcBasicAbs {
    abstract success(data: any, id: number, message?: string): BasicRendererIpcPreload;

    abstract fail(data: any, id: number, message?: string): BasicRendererIpcPreload;

    abstract ipcSet<T extends string>(ipcConst: IpcConst, callback: (event, args: T) => void): void;
}

class BasicIpcResponse implements IpcBasicAbs {
    ipcMap: Map<IpcConst, (event: IpcMainEvent, args) => void> = new Map();

    fail<T = null>(data: T, id: number, message?: string): BasicRendererIpcPreload {
        return {
            id,
            success: false,
            message: message || 'failed',
            data: data || null,
        };
    }

    success<T = null>(data: T, id: number, message?: string): BasicRendererIpcPreload {
        return {
            id,
            success: true,
            message: message || 'success',
            data,
        };
    }

    init() {
        this.ipcMap?.forEach((value, key) => {
            ipcMain.on(key as any, value);
        });
    }

    ipcSet<T>(ipcConst: IpcConst, callback: (event, args: T) => void): BasicIpcResponse {
        this.ipcMap.set(ipcConst, (event, args) => callback(event, args));
        return this;
    }
}


export {
    BasicIpcResponse,
};
