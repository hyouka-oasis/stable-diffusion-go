import { BasicResponse } from "./basicPageInfoResponse.ts";

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

