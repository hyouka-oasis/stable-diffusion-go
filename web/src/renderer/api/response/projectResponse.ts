import { BasicResponse } from "renderer/api/response/basicPageInfoResponse";
import { FileResponse } from "renderer/api/response/fileResponse";

export interface ProjectResponse extends BasicResponse {
    name: string;
}

export interface ProjectDetailParticiple {
    minLength: string;
    maxLength: string;
    minWords: string;
    maxWords: string;
}

export interface Info extends BasicResponse {
    text: string;
    prompt: string;
    negativePrompt: string;
    role: string;
    projectDetailId?: number;
    stableDiffusionImages?: Array<Partial<FileResponse> & {
        InfoId?: number;
        fileId?: number;
        projectDetailId?: number;
    }>;
    stableDiffusionImageId: number;
    audioConfig?: {
        srtLimit: string;
        voice: string;
        rate: string;
        volume: string;
        pitch: string;
    };
    audioCreateStatus?: boolean;
    loading?: boolean;
}

export interface ProjectDetailResponse extends BasicResponse {
    projectId: number;
    fileName: string;
    participleConfig: ProjectDetailParticiple;
    infoList: Info[];
    stableDiffusionConfig?: string;
    audioConfig?: {
        srtLimit: string;
        voice: string;
        rate: string;
        volume: string;
        pitch: string;
    };
}
