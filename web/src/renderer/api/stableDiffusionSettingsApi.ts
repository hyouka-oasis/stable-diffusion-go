import { deleteApi, getApi, postApi } from "renderer/request/request";
import { BasicPageInfoRequest } from "renderer/api/request/basicPageInfoRequest";
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
 * 获取配置
 * @param data
 */
export const getStableDiffusionSettings = (data: {
    id: number
}): Promise<any> => {
    return postApi({
        url: "/sdapi/settings/detail",
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
