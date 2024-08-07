import { join } from 'path';
import { app, BrowserWindow, BrowserWindowConstructorOptions, dialog, shell } from 'electron';
import { isDebug, resolveHtmlPath } from './util';
import { closedBrowserWindowHandler, closeDevTools, killRelaunchServer, readyToShowBrowserWindowHandler } from "./shared/electronFormalHelper";
import { resourcePath } from "./shared/pathHelper";
import { ElectronBrowserWindowDefaultConfig } from "./shared/browserWindowConfigUtils";
import browserWindowHelp from "./shared/browserWindowHelp";
import { BROWSER_WINDOW_KEY } from "./shared/ipcConst";
import { BasicIpc } from "./ipc/basicIpc";
import { error } from "./shared/debugLog";
import { system } from "systeminformation";
import GoServiceController from "./controllers/goServiceController";
import { ProcessHelper } from "./shared/processHelper";

let mainWindow: BrowserWindow | null = null;
let ipcHandler: BasicIpc | null = null;

const goServiceController = new GoServiceController();

if (isDebug) {
    require('electron-debug')();
    const sourceMapSupport = require('source-map-support');
    sourceMapSupport.install();
}

async function startGoService(): Promise<string | undefined> {
    return await goServiceController.startGoProgram();
}

const createWindow = async () => {
    /**
     * 如果是正式版本关闭开发者工具
     */
    if (app.isPackaged) {
        closeDevTools(mainWindow);
    }
    const pids = ProcessHelper.getOasisServerPids();
    const getAssetPath = (...paths: string[]): string => join(resourcePath, ...paths);
    const port = await startGoService();
    if (!port || (pids && pids.length && app.isPackaged)) {
        error("go启动出现了问题, 不允许被启动", port, pids);
        killRelaunchServer({
            icon: getAssetPath("icon.png"),
        });
        return;
    }
    const mainBrowserWindowConfig: BrowserWindowConstructorOptions = ElectronBrowserWindowDefaultConfig.mainBrowserWindowConfig({ icon: getAssetPath("icon.png") });
    mainWindow = browserWindowHelp.initBrowserWindow(BROWSER_WINDOW_KEY.MAIN_BROWSER, mainBrowserWindowConfig);
    // 一定得确保在browserWindow后初始化IPC以确保拿到正确的browserWindow
    ipcHandler = new BasicIpc();
    const webRenderPath = resolveHtmlPath("index.html");
    /**
     * 关闭的时候处理
     */
    closedBrowserWindowHandler(mainWindow);
    /**
     * 显示的时候处理
     */
    readyToShowBrowserWindowHandler(mainWindow, (mainWindow) => {
        mainWindow?.webContents.send("SERVER_PORT", port);
    });
    /**
     * 在app准备完成并且加载完file的时候发送给渲染进程
     */
    mainWindow?.webContents.send("SERVER_PORT", port);

    await mainWindow?.loadURL(webRenderPath);

    mainWindow?.webContents.setWindowOpenHandler(edata => {
        shell.openExternal(edata.url);
        return { action: 'deny' };
    });
};

app.on("quit", () => {
    if (goServiceController) {
        goServiceController.endGoProgram();
    }
    ipcHandler?.replace();
    browserWindowHelp?.destroyBrowserWindow();
});

app.on('window-all-closed', () => {
    app.quit();
});

const onWhenReady = () => {
    app.whenReady()
        .then(async () => {
            await createWindow();
            app.on("activate", () => {
                if (mainWindow === null) {
                    createWindow();
                }
            });
        })
        .catch(error);
};

system().then(systemData => {
    if (systemData.virtual) {
        dialog.showErrorBox("启动失败", "Reality 3D禁止在虚拟机中使用");
        app.quit();
        return;
    }
    onWhenReady();
}).catch(e => {
    error(e, "校验虚拟机失败原因");
});
