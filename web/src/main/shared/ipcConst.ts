/**
 * ipc 消息名称
 * TODO 注意这边由于多窗口事件IPC通讯与2023/8/23修改为多窗口通讯, 注册IPC方法中的第二个参数browserWindow变更为当前窗口
 * TODO 所以再调用IpcConst事件名称时请带上主体名称
 * TODO example 如果是多窗口供用的则保持 WINDOW_RESIZE(比如)形式
 * TODO 如在窗口使用的Ipc_Event则带上窗口Key 如(main-browser主窗口)MAIN_BROWSER_WINDOW_RESIZE
 */
export enum IpcConst {
    MAIN_BROWSER_WINDOW_SETTING = "MAIN_BROWSER_WINDOW_SETTING", // 窗口设置消息名称
    MAIN_BROWSER_GET_WINDOW_SETTING = "MAIN_BROWSER_GET_WINDOW_SETTING", // 获取窗口配置
    CLOSE_NEW_BROWSER_WINDOW = "CLOSE_NEW_BROWSER_WINDOW", // 关闭browser
    CLOSE_ALL_BROWSER_WINDOW = "CLOSE_ALL_BROWSER_WINDOW", // 关闭browser
    FOLDER_READ = "FOLDER_READ", //文件夹/文件选取
    FILE_OPEN = "FILE_OPEN", // 打开文件所在目录
    FILE_READ = "FILE_READ", // 文件读取
}


export const BasicIpcKeys = {
    ...IpcConst,
};


/**
 * 所有窗体的key
 */
export enum BROWSER_WINDOW_KEY {
    MAIN_BROWSER = "MAIN_BROWSER", // 主窗口
}


// 全局除了ElectronStoreIpcConst只有这一个通讯
export enum BasicIpcConst {
    BASIC_RENDERER_IPC_KEY = "BASIC_RENDERER_IPC_KEY", // renderer发送通讯给main主体
    BASIC_MAIN_IPC_KEY = "BASIC_MAIN_IPC_KEY", // main监听的key
}
