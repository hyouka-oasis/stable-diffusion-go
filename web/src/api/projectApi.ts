import { deleteApi, getApi, postApi, postFormApi } from "../utils/request";
import { BasicPageInfoRequest } from "./request/basicPageInfoRequest.ts";
import { BasicArrayResponses, BasicResponse } from "./response/basicPageInfoResponse.ts";
import { ProjectDetailResponse, ProjectResponse } from "./response/projectResponse.ts";

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

export const updateProjectDetail = (data: any): Promise<ProjectResponse> => {
    return postFormApi({
        url: "/projectDetail/upload",
        data
    });
};

export const getProjectDetail = (data: any): Promise<ProjectDetailResponse> => {
    return getApi({
        url: "/projectDetail/get",
        data
    });
};
/**
 * 进行角色提取
 * @param data
 */
export const extractTheCharacterProjectDetailParticipleList = (data: Pick<BasicResponse, "id">): Promise<ProjectDetailResponse> => {
    return postApi({
        url: "/projectDetailParticipleList/extractCharacter",
        data
    });
};
/**
 * 转换
 * @param data
 */
export const TranslateProjectDetailParticipleList = (data: Pick<BasicResponse, "id">): Promise<ProjectDetailResponse> => {
    return postApi({
        url: "/projectDetailParticipleList/translate",
        data
    });
};
