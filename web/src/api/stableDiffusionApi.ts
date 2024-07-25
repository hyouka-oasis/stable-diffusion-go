import { getApi } from "../utils/request";

export const getStableDiffusionList = () => {
    return getApi({
        url: "/stableDiffusion/getConfig"
    });
};
