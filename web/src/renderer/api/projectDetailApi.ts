import { deleteApi, getApi, postApi, postFormApi } from "renderer/request/request";
import { ProjectDetailResponse, ProjectResponse } from "renderer/api/response/projectResponse";

export const createProjectDetail = (data: any): Promise<ProjectDetailResponse> => {
    return postApi({
        url: "/projectDetail/create",
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
 * 删除详情
 * @param data
 */
export const deleteProjectDetail = (data: {
    id: number
}): Promise<ProjectDetailResponse> => {
    return deleteApi({
        url: "/projectDetail/delete",
        data
    });
};
/**
 * 更新数据
 * @param data
 */
export const updateProjectDetail = (data: Partial<ProjectDetailResponse & {
    batch?: boolean
}>): Promise<ProjectDetailResponse> => {
    return postApi({
        url: "/projectDetail/update",
        data
    });
};

