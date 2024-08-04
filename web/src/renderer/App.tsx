import { Layout, Menu, message, notification, theme } from 'antd';
import React, { Suspense, useEffect, useState } from 'react';
import { Routes, useLocation, useNavigate } from 'react-router-dom';
import { SettingOutlined, UserOutlined, VideoCameraOutlined } from "@ant-design/icons";
import router from "renderer/router";
import { ReactSmoothScrollbar } from "renderer/components/smooth-scroll/SmoothScroll";
import { AppGlobalContext, MessageBoxConfigProps } from "renderer/shared/context/appGlobalContext";
import NavigationBar from "renderer/components/navigation-bar/NavigationBar";
import { navBarHeight } from "renderer/shared";
import styled from "styled-components";

const { Content, Sider } = Layout;

const AppWrapper = styled.div`
    height: 100%;
    width: 100%;

    .ant-menu {
        border-inline-end: none !important;
    }
`;

export default function App() {
    const [ collapsed, setCollapsed ] = useState(false);
    const navigate = useNavigate();
    const location = useLocation();
    const {
        token: { colorBgContainer },
    } = theme.useToken();
    const [ messageApi, messageContextHolder ] = message.useMessage();
    const [ menuSelectedKey, setMenuSelectedKey ] = useState<string[]>([]);
    const [ notificationApi, notificationContextHolder ] = notification.useNotification();


    /**
     * notification通知提醒框
     * @param config
     */
    const openRenderNotification = (config: Omit<MessageBoxConfigProps, "messageType">) => {
        const {
            type = "error",
            message,
            description,
            placement,
            messageBoxClassname,
            messageBoxStyle,
            notificationBtn,
            ...props
        } = config;
        notificationApi?.[type]({
            ...props,
            className: `${messageBoxClassname} global-notification-box`,
            message: <span title={(message as string) ?? ""}>{message as string}</span>,
            description,
            placement,
            style: messageBoxStyle,
            btn: notificationBtn,
        });
    };

    /**
     * message提示框
     * @param config
     */
    const openRenderMessage = (
        config: Omit<MessageBoxConfigProps,
            "messageType" | "description" | "placement" | "notificationBtn">,
    ): () => void => {
        const { type = "error", message, messageBoxClassname, messageBoxStyle, ...props } = config;
        return messageApi.open({
            type,
            className: `${messageBoxClassname} global-message-box`,
            content: <span title={(message as string) ?? ""}>{message as string}</span>,
            style: messageBoxStyle,
            ...(props as any),
        });
    };

    const openMessageBox = (config: MessageBoxConfigProps): (() => void) | void => {
        const {
            type = "error",
            message = "温馨提示",
            description = "",
            messageType = "message",
            notificationBtn,
            ...props
        } = config;
        if (!messageType) {
            return openRenderMessage({ type, message, ...props });
        }
        if (messageType === "notification") {
            return openRenderNotification({ type, message, description, notificationBtn, ...props });
        }
        if (messageType === "message") {
            return openRenderMessage({ type, message, ...props });
        }
    };


    const onMenuClick = async (info: any) => {
        const { key } = info;
        navigate(key);
        setMenuSelectedKey([ key ]);
    };


    useEffect(() => {
        if (location.pathname === "/" || location.pathname === "/detail") {
            setMenuSelectedKey([ "/" ]);
        } else {
            setMenuSelectedKey([ location.pathname ]);
        }
    }, [ location ]);
    return (
        <AppGlobalContext.Provider
            value={{
                openMessageBox
            }}>
            <AppWrapper>
                <Layout style={{ height: "100%" }}>
                    {messageContextHolder}
                    {notificationContextHolder}
                    <Sider theme={"light"} onCollapse={(value) => setCollapsed(value)} collapsible collapsed={collapsed}>
                        <Menu
                            selectedKeys={menuSelectedKey}
                            mode="inline"
                            defaultSelectedKeys={[ '/' ]}
                            onClick={onMenuClick}
                            items={[
                                {
                                    key: '/',
                                    icon: <UserOutlined/>,
                                    label: '项目管理',
                                },
                                {
                                    key: '/loras',
                                    icon: <VideoCameraOutlined/>,
                                    label: 'loras管理',
                                },
                                {
                                    key: '/negativePrompts',
                                    icon: <VideoCameraOutlined/>,
                                    label: '通用反向提示词管理',
                                },
                                {
                                    key: '/files',
                                    icon: <VideoCameraOutlined/>,
                                    label: '附件管理',
                                },
                                {
                                    key: '/settings',
                                    icon: <SettingOutlined/>,
                                    label: '系统设置',
                                },
                            ]}
                        />
                        <div className="demo-logo-vertical"/>
                    </Sider>
                    <Layout>
                        <NavigationBar/>
                        <ReactSmoothScrollbar style={{ maxHeight: `calc(100vh - ${navBarHeight()})` }}>
                            <span/>
                            <Content
                                style={{
                                    padding: 24,
                                    minHeight: 280,
                                }}
                            >
                                <Suspense>
                                    <Routes>
                                        {router}
                                    </Routes>
                                </Suspense>
                            </Content>
                        </ReactSmoothScrollbar>
                    </Layout>
                </Layout>
            </AppWrapper>
        </AppGlobalContext.Provider>
    );
}
