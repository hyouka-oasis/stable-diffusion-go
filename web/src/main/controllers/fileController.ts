import { BrowserWindow, dialog, OpenDialogReturnValue, shell } from "electron";
import fs from "fs";
import { error, log } from "../shared/debugLog";
import { FileFolderOptions } from "../../renderer/ipc/types";

class FileController {

    folderSelect(
        manWindow: BrowserWindow,
        options: FileFolderOptions,
    ): Promise<OpenDialogReturnValue> {
        return new Promise((resolve, reject) => {
            dialog
                .showOpenDialog(manWindow, options)
                .then(value => {
                    log("当前选中的文件目录", value);
                    resolve(value);
                })
                .catch(e => {
                    error("当前选中的文件目录", e);
                    reject(e);
                });
        });
    }

    /**
     * 创建文件夹
     * @param filePath
     */
    createFolder(filePath: string) {
        return new Promise((resolve, reject) => {
            try {
                if (!fs.existsSync(filePath)) {
                    fs.mkdirSync(filePath, { recursive: true });
                    log("创建文件夹成功", filePath);
                    resolve(true);
                } else {
                    log("创建文件夹成功1", filePath);
                    resolve(true);
                }
            } catch (e) {
                reject(`创建目录失败${e?.toString()}`);
            }
        });
    }

    /**
     * 异步读取文件
     * @param filePath
     * @param options
     */
    readFile(filePath, options?: { encoding: BufferEncoding, flag?: string | undefined } | BufferEncoding): Promise<string> {
        return new Promise((resolve, reject) => {
            if (!fs.existsSync(filePath)) {
                reject(null);
            } else {
                fs.readFile(filePath, options ?? "utf8", (error, data) => {
                    if (error) {
                        reject(error);
                    } else {
                        resolve(data);
                    }
                });
            }
        });
    }

    /**
     * 同步读取文件
     * @param filePath
     * @param options
     */
    readFileSync(filePath, options?: { encoding: BufferEncoding, flag?: string | undefined } | BufferEncoding) {
        const newOptions = options ?? "utf8";
        return fs.readFileSync(filePath, newOptions);
    }

    /**
     * 打开文件所在目录
     * @param filePath
     */
    openFilePath(filePath): Promise<string> {
        return new Promise<string>((resolve, reject) => {
            this.checkWritePermissions(filePath).then(() => {
                shell.showItemInFolder(filePath);
                resolve("打开成功");
            }).catch(reject);
        });
    }

    /**
     * 校验文件是否存在写操作
     * @param folderPath
     * @param constants
     */
    async checkWritePermissions(folderPath: string, constants: number = fs.constants.F_OK): Promise<boolean> {
        return new Promise((resolve, reject) => {
            try {
                fs.access(folderPath, constants, errorMessage => {
                    if (errorMessage) {
                        error(`读取文件信息失败信息: ${constants} ${errorMessage}`);
                        reject(errorMessage);
                    } else {
                        log(`读取文件权限成功信息: ${constants}`, folderPath);
                        // 有写入权限
                        resolve(true);
                    }
                });
            } catch (errorMessage) {
                error(`读取失败信息:${constants} ${errorMessage}`);
                reject(errorMessage);
            }
        });
    }
}

export default new FileController();
