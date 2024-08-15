import { deleteApi, getApi, postApi } from "renderer/request/request";
import { BasicPageInfoRequest } from "renderer/api/request/basicPageInfoRequest";
import { LorasResponse } from "renderer/api/response/lorasResponse";
import { BasicArrayResponses } from "renderer/api/response/basicPageInfoResponse";
import { StableDiffusionNegativePromptResponse } from "renderer/api/response/stableDiffusionResponse";

/**
 * 获取配置列表
 * @param data
 */
export const getStableDiffusionSettingsList = (data: BasicPageInfoRequest<Partial<StableDiffusionNegativePromptResponse>>): Promise<BasicArrayResponses<StableDiffusionNegativePromptResponse>> => {
    return getApi({
        url: "/sdapi/settings/get",
        data
    });
};
/**
 * 创建配置
 * @param data
 */
export const createStableDiffusionSettings = (data: Partial<StableDiffusionNegativePromptResponse>): Promise<void> => {
    return postApi({
        url: "/sdapi/settings/create",
        data
    });
};

/**
 * 更新
 * @param data
 */
export const updateStableDiffusionSettings = (data: Partial<StableDiffusionNegativePromptResponse>): Promise<void> => {
    return postApi({
        url: "/sdapi/settings/update",
        data
    });
};

/**
 * 删除
 * @param data
 */
export const deleteStableDiffusionSettings = (data: {
    ids: number[];
}): Promise<void> => {
    return deleteApi({
        url: "/sdapi/settings/delete",
        data
    });
};

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

export const stableDiffusionText2Image = (data: {
    ids: number[];
    projectDetailId?: number;
}): Promise<string[]> => {
    return postApi({
        url: "/sdapi/images/text2image",
        data
    });
};

export const stableDiffusionDeleteImage = (data: {
    ids: number[];
}): Promise<void> => {
    return deleteApi({
        url: "/sdapi/images/deleteImage",
        data
    });
};

export const getNegativePromptList = (data: BasicPageInfoRequest): Promise<BasicArrayResponses<StableDiffusionNegativePromptResponse>> => {
    return getApi({
        url: "/sdapi/negativePromptList",
        data
    });
};

export const createNegativePrompt = (data: Partial<StableDiffusionNegativePromptResponse>): Promise<void> => {
    return postApi({
        url: "/sdapi/createNegativePrompt",
        data
    });
};

export const updateNegativePrompt = (data: Partial<StableDiffusionNegativePromptResponse>): Promise<void> => {
    return postApi({
        url: "/sdapi/updateNegativePrompt",
        data
    });
};

export const deleteNegativePrompt = (data: Partial<StableDiffusionNegativePromptResponse>): Promise<void> => {
    return deleteApi({
        url: "/sdapi/deleteNegativePrompt",
        data
    });
};

export const addImage = (data: Partial<any>): Promise<void> => {
    return postApi({
        url: "/sdapi/images/addImage",
        data
    });
};


export const getSdModels = (): Promise<BasicArrayResponses<any>> => {
    return getApi({
        url: "/sdapi/v1/sd-models",
    });
};


export const getSdVae = (): Promise<BasicArrayResponses<any>> => {
    return getApi({
        url: "/sdapi/v1/sd-vae",
    });
};


export const getSamplers = (): Promise<BasicArrayResponses<any>> => {
    return getApi({
        url: "/sdapi/v1/samplers",
    });
};

export const getSchedulers = (): Promise<BasicArrayResponses<any>> => {
    return getApi({
        url: "/sdapi/v1/schedulers",
    });
};
