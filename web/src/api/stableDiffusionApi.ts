import { getApi, postApi } from "../utils/request";
import { BasicPageInfoRequest } from "./request/basicPageInfoRequest.ts";
import { BasicArrayResponses } from "./response/basicPageInfoResponse.ts";
import { LorasResponse } from "./response/lorasResponse.ts";

export const getStableDiffusionList = () => {
    return getApi({
        url: "/stableDiffusion/getConfig"
    });
};


export const getStableDiffusionLorasList = (data: BasicPageInfoRequest<Partial<LorasResponse>>): Promise<BasicArrayResponses<LorasResponse>> => {
    return getApi({
        url: "/stableDiffusion/getLoras",
        data
    });
};


export const createStableDiffusionLoras = (data: any) => {
    return postApi({
        url: "/stableDiffusion/createLoras",
        data
    });
};

export const stableDiffusionText2Image = (data: {
    id: number;
    projectDetailId?: number;
}): Promise<string[]> => {
    return postApi({
        url: "/stableDiffusion/text2image",
        data
    });
};
