import { IpcConst } from "main/shared/ipcConst";

export interface BasicMainIpcPreload<T = unknown> {
    key?: IpcConst,
    data: T
    id: number;
}

export interface BasicRendererIpcPreload<T = unknown> extends Omit<BasicMainIpcPreload<T>, 'key' | 'args'> {
    id: number;
    success: boolean;
    message: string;
    data: T;
}
