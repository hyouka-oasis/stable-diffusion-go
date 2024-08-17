import { deleteApi, postApi } from "renderer/request/request";

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


export const addImage = (data: Partial<any>): Promise<void> => {
    return postApi({
        url: "/sdapi/images/addImage",
        data
    });
};
