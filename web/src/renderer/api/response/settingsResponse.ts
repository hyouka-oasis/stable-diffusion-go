import { BasicResponse } from "renderer/api/response/basicPageInfoResponse";

export type TranslateType = "ollama" | "chatgpt" | "aliyun"

export interface SettingsResponse extends BasicResponse {
    stableDiffusionConfig?: {
        url: string
    };
    translateType: TranslateType;
    ollamaConfig: {
        modelName: string;
        url: string
    };
}

