import { postApi } from "../utils/request";
import { BasicResponse } from "./response/basicPageInfoResponse.ts";


export const createAudioSrt = (data: Pick<BasicResponse, "id">): Promise<void> => {
    return postApi({
        url: "/audioSrt/create",
        data
    });
};

