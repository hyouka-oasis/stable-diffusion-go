import { deleteApi, getApi, postApi } from "renderer/request/request";
import { BasicPageInfoRequest } from "renderer/api/request/basicPageInfoRequest";
import { ProjectResponse } from "renderer/api/response/projectResponse";
import { BasicArrayResponses } from "renderer/api/response/basicPageInfoResponse";

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

export const updateProject = (data: any) => {
    return postApi({
        url: "/project/update",
        data
    });
};


