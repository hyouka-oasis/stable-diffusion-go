import fs from "fs";
import { log } from "../shared/debugLog";

/**
 * @author hyouka
 * @description 文件操作
 */

export default class FileHelper {

    /**
     * 创建文件
     */
    static createFile(filePath: string, data: string | NodeJS.ArrayBufferView): Promise<string> {
        return new Promise((resolve, reject) => {
            // FileHelper.createFolder(filePath, true);
            fs.writeFile(filePath, data, error => {
                if (error) {
                    reject(error);
                    log("文件创建失败", error);
                }
                log("文件创建成功", filePath);
                resolve(filePath);
            });
        });
    }
}
