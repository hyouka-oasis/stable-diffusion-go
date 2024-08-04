import { BasicResponse } from "./basicPageInfoResponse.ts";


export interface LorasResponse extends BasicResponse {
    name: string;
    roles?: string;
    image?: string;
}
