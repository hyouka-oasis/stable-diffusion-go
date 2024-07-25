import { cloneElement, lazy, ReactElement } from "react";
import { Route } from "react-router-dom";

const Project = lazy(() => import("../pages/project/Project.tsx"));
const ProjectDetail = lazy(() => import("../pages/project-detail/ProjectDetail.tsx"));
const Settings = lazy(() => import("../pages/settings/Settings.tsx"));


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
        element: <Project/>,
    },
    {
        path: "/detail",
        element: <ProjectDetail/>,
    },
    {
        path: "/settings",
        element: <Settings/>,
    }
];

export default routeRender(routers);
