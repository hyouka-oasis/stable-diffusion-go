import { getApi, postApi } from "renderer/request/request";
import { BasicPageInfoRequest } from "renderer/api/request/basicPageInfoRequest";
import { LorasResponse } from "renderer/api/response/lorasResponse";
import { BasicArrayResponses } from "renderer/api/response/basicPageInfoResponse";

export const getStableDiffusionLorasList = (data: BasicPageInfoRequest<Partial<LorasResponse>>): Promise<BasicArrayResponses<LorasResponse>> => {
    return getApi({
        url: "/sdapi/loras/get",
        data
    });
};

export const createStableDiffusionLoras = (data: any) => {
    return postApi({
        url: "/sdapi/loras/create",
        data
    });
};
