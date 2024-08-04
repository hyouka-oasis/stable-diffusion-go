/**
 * @desc 基础窗口操作
 * @author hyouka
 */
import { BROWSER_WINDOW_KEY, IpcConst } from "../shared/ipcConst";
import browserWindowHelp from "../shared/browserWindowHelp";

export const browserWindowIpc = {
    /**
     * 关闭指定窗口
     * @param id
     */
    [IpcConst.CLOSE_NEW_BROWSER_WINDOW](id: BROWSER_WINDOW_KEY | BROWSER_WINDOW_KEY []) {
        const ids = Array.isArray(id) ? id : [ id ];
        for (const browserId of ids) {
            const browserWindow = browserWindowHelp.getBrowserWindow(browserId);
            if (!browserWindow) {
                continue;
            }
            browserWindowHelp.deleteBrowserWindow(browserId);
        }
    },
    /**
     * 关闭所有窗口
     */
    [IpcConst.CLOSE_ALL_BROWSER_WINDOW]() {
        browserWindowHelp.destroyBrowserWindow();
    },
};
