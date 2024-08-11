import { BasicResponse } from "renderer/api/response/basicPageInfoResponse";

export interface FileResponse extends BasicResponse {
    fileId: number;
    name: string;
    url: string;
    tag: string;
    key: string;
}
