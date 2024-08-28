import { BasicResponse } from "renderer/api/response/basicPageInfoResponse";

export interface TaskErrors {
    error: string;
}

export interface TaskResponse extends BasicResponse {
    errors: TaskErrors[];
    status: number;
    progress: number;
    message: string;
}

