/**
 * @desc 基础导航栏-右侧（含最大最小化）
 * @author mxm
 * @date 2023/3/21 10:15
 */
import React, { useContext, useEffect, useMemo, useState } from "react";
import { NavigationBarActionWrap } from "./StyleComponents";
import NavBackUp from "renderer/assets/svg-com/nav-back-up.svg";
import NavMaximize from "renderer/assets/svg-com/nav-maximize.svg";
import NavMinimize from "renderer/assets/svg-com/nav-minimize.svg";
import NavShutDown from "renderer/assets/svg-com/nav-shut-down.svg";
import { ipcApi } from "renderer/ipc/BasicRendererIpcAdapter";
import { conversionReactChildren, InsertNavbarChildType } from "renderer/shared/basic/reactElementUtils";
import { NavigationBarProps } from "renderer/components/navigation-bar/NavigationBar";
import { AppGlobalContext } from "renderer/shared/context/appGlobalContext";

const { platformHelper } = window.electron ?? {};

export interface NavigationBarActionsProps extends Pick<NavigationBarProps, "customActionBeforeChild" |
    "browserKey" | "customActionBeforeTabClick" | "debugOptions" | "navigationActionHandler"> {
    children?: React.ReactElement;
}

const NavigationBarActions: React.FC<NavigationBarActionsProps> = (props) => {
    const {
        browserKey, customActionBeforeChild = [], children,
        customActionBeforeTabClick, debugOptions, navigationActionHandler
    } = props;
    const [ isMaximized, setIsMaximized ] = useState(false);
    const { openMessageBox } = useContext(AppGlobalContext);


    const getIsMaximizedState = async () => {
        try {
            const { isWindow } = platformHelper?.validatePlatform();
            const key = isWindow ? "isMaximized" : "isFullScreen";
            const result = await ipcApi.windowAdapter.getElectronWindowConfig(key, browserKey);
            if (!result.success) {
                console.error("获取窗口是否是最大化异常");
                openMessageBox({ type: "error", message: `获取窗口是否是最大化异常,${result}` });
                return;
            }
            setIsMaximized(result.data);
        } catch (e) {
            console.error(e);
        }
    };
    /**
     * 关闭事件
     */
    const onCloseBrowserWindowHandler = () => {
        // 如果browserKey存在则关闭当前窗口
        if (browserKey) {
            ipcApi.windowAdapter.closeBrowserWindow(browserKey);
            return;
        }
        // 反之关闭全部窗口
        ipcApi.windowAdapter.closeAllBrowserWindow();
    };

    const onCloseHandler = () => {
        if (navigationActionHandler?.onNavigationCloseHandler) {
            navigationActionHandler?.onNavigationCloseHandler?.(onCloseBrowserWindowHandler);
            return;
        }
        onCloseBrowserWindowHandler();
    };

    const onNavigationCloseHandler = async () => {
        onCloseHandler();
    };

    const onNavChildClickHandler = (key: string, e?: React.MouseEvent, child?: InsertNavbarChildType) => {
        customActionBeforeTabClick?.(key, e, child);
    };

    const memoNavbarChild = useMemo(() => conversionReactChildren(customActionBeforeChild, {
        onTabClickHandler: onNavChildClickHandler,
        selfChildren: children
    }), [ customActionBeforeChild ]);

    const setBrowserWindowMaxOrMin = (windowControllerType: "max" | "min", handlerDeps?: boolean) => {
        ipcApi.windowAdapter.setElectronWindowConfig({
            windowControllerType: windowControllerType,
            browserWindowKey: browserKey
        });
        if (handlerDeps) {
            getIsMaximizedState();
        }
    };

    useEffect(() => {
        window.addEventListener("resize", () => getIsMaximizedState());
        getIsMaximizedState();
        return () => {
            window.removeEventListener("resize", () => getIsMaximizedState());
        };
    }, []);

    return <NavigationBarActionWrap>
        <div className={"actions-container"}>
            {memoNavbarChild}
            <div
                className={"actions"}
            >
                <div className={"baseAction"} onClick={() => setBrowserWindowMaxOrMin("min")}>
                    <NavMinimize/>
                </div>
                <div className={"baseAction"} onClick={() => setBrowserWindowMaxOrMin("max", true)}>
                    {isMaximized ? <NavBackUp/> : <NavMaximize/>}
                </div>
                <div className={"baseAction"} onClick={onNavigationCloseHandler}>
                    <NavShutDown/>
                </div>
            </div>
        </div>
    </NavigationBarActionWrap>;
};
export default NavigationBarActions;
