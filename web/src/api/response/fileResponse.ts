import { BasicResponse } from "./basicPageInfoResponse.ts";


export interface FileResponse extends BasicResponse {
    name: string;
    url: string;
    tag: string;
    key: string;
}
