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

export interface ProjectDetailParticipleList extends BasicResponse {
    text: string;
}

export interface ProjectDetailResponse extends BasicResponse {
    projectId: number;
    fileName: string;
    participle: ProjectDetailParticiple;
    participleList: ProjectDetailParticipleList[];
}
