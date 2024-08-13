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
        url: "/stableDiffusion/settings/get",
        data
    });
};
/**
 * 创建配置
 * @param data
 */
export const createStableDiffusionSettings = (data: Partial<StableDiffusionNegativePromptResponse>): Promise<void> => {
    return postApi({
        url: "/stableDiffusion/settings/create",
        data
    });
};

/**
 * 更新
 * @param data
 */
export const updateStableDiffusionSettings = (data: Partial<StableDiffusionNegativePromptResponse>): Promise<void> => {
    return postApi({
        url: "/stableDiffusion/settings/update",
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
        url: "/stableDiffusion/settings/delete",
        data
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
    ids: number[];
    projectDetailId?: number;
}): Promise<string[]> => {
    return postApi({
        url: "/stableDiffusion/text2image",
        data
    });
};

export const stableDiffusionDeleteImage = (data: {
    ids: number[];
}): Promise<void> => {
    return deleteApi({
        url: "/stableDiffusion/deleteImage",
        data
    });
};

export const getNegativePromptList = (data: BasicPageInfoRequest): Promise<BasicArrayResponses<StableDiffusionNegativePromptResponse>> => {
    return getApi({
        url: "/stableDiffusion/negativePromptList",
        data
    });
};

export const createNegativePrompt = (data: Partial<StableDiffusionNegativePromptResponse>): Promise<void> => {
    return postApi({
        url: "/stableDiffusion/createNegativePrompt",
        data
    });
};

export const updateNegativePrompt = (data: Partial<StableDiffusionNegativePromptResponse>): Promise<void> => {
    return postApi({
        url: "/stableDiffusion/updateNegativePrompt",
        data
    });
};

export const deleteNegativePrompt = (data: Partial<StableDiffusionNegativePromptResponse>): Promise<void> => {
    return deleteApi({
        url: "/stableDiffusion/deleteNegativePrompt",
        data
    });
};

export const addImage = (data: Partial<any>): Promise<void> => {
    return postApi({
        url: "/stableDiffusion/addImage",
        data
    });
};
