import { getApi, postApi } from "../utils/request";
import { SettingsResponse } from "./response/settingsResponse.ts";

interface SettingsApiProps {
    name: string;
}

export const createProject = (data: SettingsApiProps) => {
    return postApi({
        url: "/project/create",
        data
    });
};

export const updateSettings = (data: Partial<SettingsResponse>) => {
    return postApi({
        url: "/settings/update",
        data
    });
};

export const getSettings = (): Promise<SettingsResponse> => {
    return getApi({
        url: "/settings/get",
    });
};
