import { postApi } from "renderer/request/request";
import { BasicResponse } from "renderer/api/response/basicPageInfoResponse";


export const createAudioSrt = (data: Pick<Partial<BasicResponse>, "id"> & {
    infoId?: number
}): Promise<void> => {
    return postApi({
        url: "/audioSrt/create",
        data
    });
};

