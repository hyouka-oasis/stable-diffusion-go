import { getApi, postApi } from "renderer/request/request";
import { BasicPageInfoRequest } from "renderer/api/request/basicPageInfoRequest";
import { TaskResponse } from "renderer/api/response/taskResponse";


export const getTaskList = (data: BasicPageInfoRequest<{
    projectDetailId?: number;
    taskId?: number;
}>) => {
    return getApi({
        url: "/task/list",
        data
    });
};

export const getTaskDetail = (data: {
    projectDetailId?: number;
    taskId?: number;
}): Promise<TaskResponse> => {
    return postApi({
        url: "/task/get",
        data
    });
};

