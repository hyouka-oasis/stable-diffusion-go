import React from 'react';
import { createRoot } from 'react-dom/client';
import App from './App';
import { StyleProvider } from '@ant-design/cssinjs';
import { HashRouter } from "react-router-dom";
import locale from "antd/locale/zh_CN";
import { ConfigProvider, theme } from "antd";
import "renderer/assets/styles/index.less";
import "renderer/theme/index.less";

const container = document.getElementById('root')!;
const root = createRoot(container);
root.render(
    <ConfigProvider
        locale={locale} theme={{
            algorithm: theme.darkAlgorithm
        }}>
        <StyleProvider>
            <HashRouter>
                <App/>
            </HashRouter>
        </StyleProvider>
    </ConfigProvider>
);
