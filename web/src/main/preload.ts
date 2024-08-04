/**
 * TODO 这边是向全局（windows）注入electron方法
 * @link https://www.electronjs.org/zh/docs/latest/tutorial/process-model#preload-%E8%84%9A%E6%9C%AC
 */
import { clipboard, contextBridge, ipcRenderer, IpcRendererEvent } from 'electron';
import { BasicIpcConst } from "./shared/ipcConst";
import { BasicMainIpcPreload, BasicRendererIpcPreload } from "./ipc/types";
import { PlatformHelper } from "./shared/platformHelper";

const electronHandler = {
    ipcRenderer: {
        /**
         * 向主进程(main)发送讯息
         * @param channel
         * @param args
         */
        sendMessage<T>(channel: BasicIpcConst, args?: BasicMainIpcPreload<T>) {
            ipcRenderer.send(channel, args);
        },
        /**
         * 渲染进程(renderer)监听信息
         * @param channel
         * @param func
         */
        on<T>(channel: BasicIpcConst, func: (args: BasicRendererIpcPreload<T>) => void): () => void {
            const subscription = (_event: IpcRendererEvent, args: BasicRendererIpcPreload<T>) => func(args);
            ipcRenderer.on(channel, subscription);
            return () => {
                ipcRenderer.removeListener(channel, subscription);
            };
        },
        /**
         * 渲染进程(renderer)接收消息
         * @param channel
         * @param func
         */
        once<T>(channel: BasicIpcConst, func: (args: BasicRendererIpcPreload<T>) => void) {
            ipcRenderer.once(channel, (_event, args: BasicRendererIpcPreload<T>) => func(args));
        },
    },
    clipboard: {
        ...clipboard,
    },
    platformHelper: {
        validatePlatform: PlatformHelper.validatePlatform
    },
};

contextBridge.exposeInMainWorld('electron', electronHandler);

export type ElectronHandler = typeof electronHandler;
