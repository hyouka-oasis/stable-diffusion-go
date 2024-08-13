import { getApi } from "renderer/request/request";
import { BasicArrayResponses } from "renderer/api/response/basicPageInfoResponse";


export const getOllamaModelList = (): Promise<BasicArrayResponses<any>> => {
    return getApi({
        url: "/ollama/get",
    });
};

