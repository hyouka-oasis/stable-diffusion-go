import { deleteApi, getApi, postApi, postFormApi } from "../utils/request";
import { BasicPageInfoRequest } from "./request/basicPageInfoRequest.ts";
import { BasicArrayResponses, BasicResponse } from "./response/basicPageInfoResponse.ts";
import { ProjectDetailInfo, ProjectDetailResponse, ProjectResponse } from "./response/projectResponse.ts";

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

export const uploadProjectDetail = (data: any): Promise<ProjectResponse> => {
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
 * 更新数据
 * @param data
 */
export const updateProjectDetail = (data: Partial<ProjectDetailResponse>): Promise<ProjectDetailResponse> => {
    return postApi({
        url: "/projectDetail/update",
        data
    });
};
/**
 * 进行角色提取
 * @param data
 */
export const extractTheCharacterProjectDetailParticipleList = (data: Pick<BasicResponse, "id">): Promise<ProjectDetailResponse> => {
    return postApi({
        url: "/projectDetailInfo/extractRole",
        data
    });
};
/**
 * 转换
 * @param data
 */
export const translateProjectDetailParticipleList = (data: Pick<Partial<BasicResponse>, "id"> & Pick<ProjectDetailInfo, "projectDetailId">): Promise<ProjectDetailResponse> => {
    return postApi({
        url: "/projectDetailInfo/translate",
        data
    });
};
/**
 * 更新数据
 * @param data
 */
export const updateProjectDetailInfo = (data: any): Promise<void> => {
    return postApi({
        url: "/projectDetailInfo/update",
        data
    });
};

/**
 * 获取数据
 * @param data
 */
export const getProjectDetailInfo = (data: {
    id: number;
}): Promise<ProjectDetailInfo> => {
    return getApi({
        url: "/projectDetailInfo/get",
        data
    });
};
