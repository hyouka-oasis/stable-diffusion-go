import { BasicResponse } from "renderer/api/response/basicPageInfoResponse";

export interface LorasResponse extends BasicResponse {
    name: string;
    roles?: string;
    image?: string;
}
