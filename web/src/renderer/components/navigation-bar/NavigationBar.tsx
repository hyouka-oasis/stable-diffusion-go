import React, { useMemo } from "react";
import { MenuProps } from "antd";
import { BROWSER_WINDOW_KEY } from "main/shared/ipcConst";
import { NavigationBarContentWrap, NavigationBarWrap } from "renderer/components/navigation-bar/StyleComponents";
import NavigationBarActions from "renderer/components/navigation-bar/NavigationBarActions";
import { conversionReactChildren, InsertNavbarChildType } from "renderer/shared/basic/reactElementUtils";

interface NavigationBarDebugOptions {
    /**
     * 是否显示debug按钮 默认true
     */
    visibleDebug?: boolean;
    /**
     * 自定义下拉菜单
     */
    debugDropdownMenu?: MenuProps["items"];

    /**
     * 拦截打开devtools事件
     */
    onOpenDebugDevTools?(): void;

    /**
     * 三连击事件
     */
    onHandleTripleClick?(e: React.MouseEvent): void;
}

/**
 * 不要往里面塞额外的东西，参数已经能够满足大部分的业务和样式要求了
 */
export interface NavigationBarProps {
    /**
     * 图标
     */
    logo?: string;
    /**
     * 版本号
     */
    version?: string;
    /**
     * 标题
     */
    title?: string;
    /**
     * 操作项前的自定义组件
     * */
    customActionBeforeChild?: React.ReactNode | InsertNavbarChildType | InsertNavbarChildType[];
    /**
     * 版本号后的自定义组件
     */
    customVersionAfterChild?: React.ReactNode | InsertNavbarChildType | InsertNavbarChildType[];
    /**
     * 嵌入在版本号与操作栏之间的元素
     * 如果存在这个元素则customBeforeActionNode和customVersionAfterChild将不生效
     */
    customNavbarChild?: React.ReactNode | InsertNavbarChildType | InsertNavbarChildType[];
    /**
     * 左边标题样式
     */
    titleStyle?: React.CSSProperties;
    /**
     * browserWindowKey
     */
    browserKey?: BROWSER_WINDOW_KEY;
    /**
     * 背景颜色
     */
    backgroundColor?: React.CSSProperties["backgroundColor"];
    /**
     * 覆盖头部文本样式
     */
    navigationTextColor?: string;
    /**
     * 是否渲染子组件
     */
    renderReactNode?: boolean;
    /**
     * 开发者按钮配置
     */
    debugOptions?: NavigationBarDebugOptions;
    /**
     * 右上角操作按钮事件
     */
    navigationActionHandler?: {
        /**
         * 点击关闭按钮时的操作
         * 返回一个回调，关闭程序的回调
         */
        onNavigationCloseHandler?(closeHandler?: () => void): void;
        /**
         * 是否在终止之前查询任务
         * 为了后续扩展
         */
        queryTaskOptions?: {
            queryTask?: boolean;
        };
    };
    children?: React.ReactElement;

    /**
     * 整体点击事件
     * @param tabKey
     * @param child
     * @param event
     */
    customNavbarTabClick?(tabKey: string, event?: React.MouseEvent, child?: InsertNavbarChildType): void;

    /**
     * 版本号后面插入的元素点击
     * @param tabKey
     * @param child
     * @param event
     */
    customVersionAfterTabClick?(tabKey: string, event?: React.MouseEvent, child?: InsertNavbarChildType): void;

    /**
     * 操作栏前的元素点击
     * @param tabKey
     * @param child
     * @param event
     */
    customActionBeforeTabClick?(tabKey: string, event?: React.MouseEvent, child?: InsertNavbarChildType): void;

    /**
     * 标题点击事件
     */
    onTitleClick?(): void;
}

const NavigationBar: React.FC<NavigationBarProps> = (props) => {
    const {
        customActionBeforeChild,
        browserKey, customNavbarChild = [],
        children, customNavbarTabClick, customActionBeforeTabClick,
        backgroundColor, navigationTextColor,
        renderReactNode = true, navigationActionHandler, debugOptions
    } = props;
    if (!renderReactNode) {
        return null;
    }

    const onNavChildClickHandler = (key: string, e?: React.MouseEvent, child?: InsertNavbarChildType) => {
        customNavbarTabClick?.(key, e, child);
    };

    const memoNavbarChild = useMemo(() => conversionReactChildren(customNavbarChild, { onTabClickHandler: onNavChildClickHandler, selfChildren: children }), [ customNavbarChild ]);

    return (
        <NavigationBarWrap backgroundColor={backgroundColor} navigationTextColor={navigationTextColor}>
            <NavigationBarContentWrap>{memoNavbarChild}</NavigationBarContentWrap>
            <NavigationBarActions
                browserKey={browserKey}
                customActionBeforeChild={customActionBeforeChild}
                customActionBeforeTabClick={customActionBeforeTabClick}
                navigationActionHandler={navigationActionHandler}
                debugOptions={debugOptions}
            />
        </NavigationBarWrap>
    );
};

export default NavigationBar;

