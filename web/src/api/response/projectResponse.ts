import { BasicResponse } from "./basicPageInfoResponse.ts";

export interface ProjectResponse extends BasicResponse {
    name: string;
}

export interface ProjectDetailParticiple {
    minLength: string;
    maxLength: string;
    minWords: string;
    maxWords: string;
}

export interface ProjectDetailInfo extends BasicResponse {
    text: string;
    prompt: string;
    negativePrompt: string;
    role: string;
    projectDetailId?: number;
}

export interface ProjectDetailResponse extends BasicResponse {
    projectId: number;
    fileName: string;
    participleConfig: ProjectDetailParticiple;
    projectDetailInfoList: ProjectDetailInfo[];
    stableDiffusionConfig?: string;
}
