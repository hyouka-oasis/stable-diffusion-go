import { postApi } from "renderer/request/request";


export const createInfoVideo = (data: { ids?: number[], projectDetailId: number }) => {
    return postApi({
        url: "/video/create",
        data
    });
};

