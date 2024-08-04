import { cloneElement, lazy, ReactElement } from "react";
import { Route } from "react-router-dom";

const ProjectPage = lazy(() => import("renderer/pages/project/Project"));
const LorasPage = lazy(() => import("renderer/pages/loras/Loras"));
const ProjectDetailPage = lazy(() => import("renderer/pages/project-detail/ProjectDetail"));
const FilesPage = lazy(() => import("renderer/pages/file/File"));
const SettingsPage = lazy(() => import("renderer/pages/settings/Settings"));
const NegativePromptsPage = lazy(() => import("renderer/pages/negative-prompts/NegativePrompts"));


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

const routers = [
    {
        path: "/",
        element: <ProjectPage/>,
    },
    {
        path: "/detail",
        element: <ProjectDetailPage/>,
    },
    {
        path: "/loras",
        element: <LorasPage/>,
    },
    {
        path: "/files",
        element: <FilesPage/>,
    },
    {
        path: "/negativePrompts",
        element: <NegativePromptsPage/>,
    },
    {
        path: "/settings",
        element: <SettingsPage/>,
    },
];

export default routeRender(routers);
