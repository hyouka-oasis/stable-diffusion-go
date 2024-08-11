import { deleteApi, postApi } from "renderer/request/request";
import { Info, ProjectDetailResponse } from "renderer/api/response/projectResponse";
import { BasicResponse } from "renderer/api/response/basicPageInfoResponse";

/**
 * 进行角色提取
 * @param data
 */
export const extractTheCharacterProjectDetailParticipleList = (data: Pick<BasicResponse, "id">): Promise<ProjectDetailResponse> => {
    return postApi({
        url: "/info/extractRole",
        data
    });
};
/**
 * 转换
 * @param data
 */
export const translateProjectDetailParticipleList = (data: Pick<Partial<BasicResponse>, "id"> & Pick<Info, "projectDetailId">): Promise<ProjectDetailResponse> => {
    return postApi({
        url: "/info/translate",
        data
    });
};
/**
 * 关键字提取
 * @param data
 */
export const keywordsExtractInfoList = (data: Pick<Partial<BasicResponse>, "id"> & Pick<Info, "projectDetailId">): Promise<ProjectDetailResponse> => {
    return postApi({
        url: "/info/keywords",
        data
    });
};
/**
 * 更新数据
 * @param data
 */
export const updateProjectDetailInfo = (data: any): Promise<void> => {
    return postApi({
        url: "/info/update",
        data
    });
};


export const deleteInfo = (data: { id: number }) => {
    return deleteApi({
        url: "/info/delete",
        data
    });
};
export const updateAudio = (data: { projectDetailId?: number }) => {
    return postApi({
        url: "/info/updateAudio",
        data
    });
};

