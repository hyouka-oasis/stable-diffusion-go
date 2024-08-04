/**
 * @author hyouka
 * @description 文件选择通讯
 */
import { BrowserWindow } from "electron";
import browserWindowHelp from "../shared/browserWindowHelp";
import { BROWSER_WINDOW_KEY, IpcConst } from "../shared/ipcConst";
import FileController from "../controllers/fileController";


export const fileIpc = {
    [IpcConst.FOLDER_READ](config) {
        const mainBrowserWindow = browserWindowHelp.getBrowserWindow(BROWSER_WINDOW_KEY.MAIN_BROWSER);
        return FileController.folderSelect(mainBrowserWindow as BrowserWindow, config);
    },
    [IpcConst.FILE_OPEN](filePath) {
        return FileController.openFilePath(filePath);
    },
    [IpcConst.FILE_READ](data) {
        const { filePath, options } = data;
        return FileController.readFile(filePath, options);
    },
};
