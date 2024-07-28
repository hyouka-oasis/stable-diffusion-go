import { Button, Layout, Menu, theme } from "antd";
import { Suspense, useEffect, useState } from "react";
import { MenuFoldOutlined, MenuUnfoldOutlined, SettingOutlined, UserOutlined, VideoCameraOutlined } from "@ant-design/icons";
import { Routes, useNavigate } from "react-router-dom";
import routers from "./router";
import { useLocation } from "react-router";

const { Content, Header, Sider } = Layout;


function App() {
    const [ collapsed, setCollapsed ] = useState(false);
    const navigate = useNavigate();
    const location = useLocation();
    const {
        token: { colorBgContainer },
    } = theme.useToken();

    const [ menuSelectedKey, setMenuSelectedKey ] = useState<string[]>([]);


    const onMenuClick = async (info: any) => {
        const { key } = info;
        navigate(key);
        setMenuSelectedKey([ key ]);
    };


    useEffect(() => {
        setMenuSelectedKey([ location.pathname ]);
        console.log(location);
    }, [ location ]);

    return (
        <Layout style={{ height: "100%" }}>
            <Sider theme={"light"} trigger={null} collapsible collapsed={collapsed}>
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
            </Sider>
            <Layout>
                <Header style={{ padding: 0, backgroundColor: colorBgContainer }}>
                    <Button
                        type="text"
                        icon={collapsed ? <MenuUnfoldOutlined/> : <MenuFoldOutlined/>}
                        onClick={() => setCollapsed(!collapsed)}
                        style={{
                            fontSize: '16px',
                            width: 64,
                            height: 64,
                        }}
                    />
                </Header>
                <Content
                    style={{
                        padding: 24,
                        minHeight: 280,
                    }}
                >
                    <Suspense>
                        <Routes>
                            {routers}
                        </Routes>
                    </Suspense>
                </Content>
            </Layout>
        </Layout>
    );
}

export default App;
