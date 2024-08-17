import { getApi } from "renderer/request/request";
import { BasicArrayResponses } from "renderer/api/response/basicPageInfoResponse";

/**
 * 获取sd模型
 */
export const getSdModels = (): Promise<BasicArrayResponses<any>> => {
    return getApi({
        url: "/sdapi/v1/sd-models",
    });
};
/**
 * 获取sd-vae
 */
export const getSdVae = (): Promise<BasicArrayResponses<any>> => {
    return getApi({
        url: "/sdapi/v1/sd-vae",
    });
};
/**
 * 获取取样器
 */
export const getSamplers = (): Promise<BasicArrayResponses<any>> => {
    return getApi({
        url: "/sdapi/v1/samplers",
    });
};
/**
 * 获取调度类型
 */
export const getSchedulers = (): Promise<BasicArrayResponses<any>> => {
    return getApi({
        url: "/sdapi/v1/schedulers",
    });
};
/**
 * 获取高清算法模型
 */
export const getUpscalers = (): Promise<BasicArrayResponses<any>> => {
    return getApi({
        url: "/sdapi/v1/upscalers",
    });
};
