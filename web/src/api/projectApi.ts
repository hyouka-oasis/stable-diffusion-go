import { deleteApi, getApi, postApi } from "../utils/request";
import { BasicPageInfoRequest } from "./request/basicPageInfoRequest.ts";
import { BasicArrayResponses } from "./response/basicPageInfoResponse.ts";
import { ProjectResponse } from "./response/projectResponse.ts";

interface ProjectApiProps {
    name: string;
}

export const createProject = (data: ProjectApiProps) => {
    return postApi({
        url: "/project/create",
        data
    });
};

export const getProjectList = (data: BasicPageInfoRequest<Partial<ProjectApiProps>>): Promise<BasicArrayResponses<ProjectResponse>> => {
    return getApi({
        url: "/project/list",
        data
    });
};

export const deleteProject = (data: { id: number }) => {
    return deleteApi({
        url: "/project/delete",
        data
    });
};
