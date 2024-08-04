import { BasicResponse } from "renderer/api/response/basicPageInfoResponse";

export interface FileResponse extends BasicResponse {
    name: string;
    url: string;
    tag: string;
    key: string;
}
