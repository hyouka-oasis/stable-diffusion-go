import React, { cloneElement, lazy, ReactElement } from "react";
import { Route } from "react-router-dom";
import { FormOutlined, QuestionOutlined, SettingOutlined, SignatureOutlined, TagOutlined } from "@ant-design/icons";

const ProjectPage = lazy(() => import("renderer/pages/project/Project"));
const LorasPage = lazy(() => import("renderer/pages/loras/Loras"));
const ProjectDetailPage = lazy(() => import("renderer/pages/project-detail/ProjectDetail"));
const FilesPage = lazy(() => import("renderer/pages/file/File"));
const SettingsPage = lazy(() => import("renderer/pages/settings/Settings"));
const HelpPage = lazy(() => import("renderer/pages/help/Help"));
const StableDiffusionSettingsPage = lazy(() => import("renderer/pages/stable-diffusion-settings/StableDiffusionSettings"));


export const routeRender = (routeList: any[]) => {
    return routeList.map((item) => {
        const { element, path, ...props } = item;
        const Com = cloneElement(element as ReactElement, { ...props });
        return (
            <Route
                path={item.path}
                element={Com}
                key={item.path}
            >
                {item?.children && routeRender(item.children)}
            </Route>
        );
    });
};

export const routers = [
    {
        path: "/",
        element: <ProjectPage/>,
        icon: <FormOutlined/>,
        label: "项目管理"
    },
    {
        path: "/detail",
        element: <ProjectDetailPage/>,
        hidden: true,
    },
    {
        path: "/stableDiffusionSettings",
        element: <StableDiffusionSettingsPage/>,
        icon: <SignatureOutlined/>,
        label: '通用stable-diffusion配置',
    },
    {
        path: "/loras",
        element: <LorasPage/>,
        icon: <TagOutlined/>,
        label: 'loras管理',
    },
    {
        path: "/files",
        element: <FilesPage/>,
        hidden: true,
    },
    {
        path: "/settings",
        element: <SettingsPage/>,
        icon: <SettingOutlined/>,
        label: '系统设置',
    },
    {
        path: "/help",
        element: <HelpPage/>,
        icon: <QuestionOutlined/>,
        label: '使用帮助',
    },
];

export default routeRender(routers);
