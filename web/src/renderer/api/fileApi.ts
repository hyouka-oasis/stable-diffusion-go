import { RcFile } from "antd/lib/upload";
import { BasicPageInfoRequest } from "renderer/api/request/basicPageInfoRequest";
import { getApi, postFormApi } from "renderer/request/request";
import { FileResponse } from "renderer/api/response/fileResponse";
import { BasicArrayResponses } from "renderer/api/response/basicPageInfoResponse";


export const getFileList = (data: BasicPageInfoRequest): Promise<BasicArrayResponses<FileResponse>> => {
    return getApi({
        url: "/file/getList",
        data
    });
};


export const uploadFile = (data: {
    file: File | RcFile;
    fileType?: string;
}): Promise<FileResponse> => {
    return postFormApi({
        url: "/file/upload",
        data
    });
};
