/* eslint import/prefer-default-export: off */
import { URL } from 'url';
import path from 'path';

export const isDebug = process.env.NODE_ENV === "development" || process.env.DEBUG_PROD === "true";


export function resolveHtmlPath(htmlFileName: string) {
    if (process.env.NODE_ENV === 'development') {
        const port = process.env.PORT || 2580;
        const url = new URL(`http://localhost:${port}`);
        url.pathname = htmlFileName;
        return url.href;
    }
    return `file://${path.resolve(__dirname, '../renderer/', htmlFileName)}`;
}
