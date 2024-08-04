import { BasicIpcConst } from "main/shared/ipcConst";

const { ipcRenderer } = window.electron ?? {};

class BasicRendererStaticIpcAdapter {
    private static primaryKey: number = 0;
    private static ipcRendererResolveMap: Map<number, (value) => void> = new Map();
    private static ipcRendererRejectMap: Map<number, (value) => void> = new Map();

    constructor() {
        ipcRenderer?.on(BasicIpcConst.BASIC_RENDERER_IPC_KEY, (response) => {
            if (response.id === undefined) {
                console.error("id为必传属性");
                return;
            }
            if (response.data !== null) {
                BasicRendererStaticIpcAdapter.getIpcRendererResolverMap(response.id)?.(response);
            } else {
                BasicRendererStaticIpcAdapter.getIpcRendererRejectMap(response.id)?.(response.message);
            }
            BasicRendererStaticIpcAdapter.deleteIpcRendererResolverMap(response.id);
            BasicRendererStaticIpcAdapter.deleteIpcRendererRejectMap(response.id);
        });
    }

    static getPrimaryKey(): number {
        return this.primaryKey++;
    }

    /**
     * 注意这边的T会自动继承new Promise<Interface>
     * @param key
     * @param resolve
     */
    static setIpcRendererResolverMap<T>(key: number, resolve: (value: T) => void) {
        this.ipcRendererResolveMap.set(key, resolve);
    }

    static getIpcRendererResolverMap(key: number) {
        return this.ipcRendererResolveMap.get(key);
    }

    static deleteIpcRendererResolverMap(key: number) {
        this.ipcRendererResolveMap.delete(key);
    }

    static setIpcRendererRejectMap<T>(key: number, resolve: (value: T) => void) {
        this.ipcRendererRejectMap.set(key, resolve);
    }

    static getIpcRendererRejectMap<T>(key: number) {
        return this.ipcRendererRejectMap.get(key);
    }

    static deleteIpcRendererRejectMap(key: number) {
        this.ipcRendererRejectMap.delete(key);
    }
}

export {
    BasicRendererStaticIpcAdapter,
};
