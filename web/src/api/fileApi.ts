import { getApi, postFormApi } from "../utils/request";
import { BasicPageInfoRequest } from "./request/basicPageInfoRequest.ts";
import { BasicArrayResponses } from "./response/basicPageInfoResponse.ts";
import { FileResponse } from "./response/fileResponse.ts";
import { RcFile } from "antd/lib/upload";


export const getFileList = (data: BasicPageInfoRequest): Promise<BasicArrayResponses<FileResponse>> => {
    return getApi({
        url: "/file/getList",
        data
    });
};


export const uploadFile = (data: {
    file: File | RcFile;
}): Promise<FileResponse> => {
    return postFormApi({
        url: "/file/upload",
        data
    });
};
