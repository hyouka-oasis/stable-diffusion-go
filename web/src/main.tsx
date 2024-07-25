import ReactDOM from 'react-dom/client';
import App from './App.tsx';
import './index.css';
import { ConfigProvider, theme } from "antd";
import locale from "antd/locale/zh_CN";
import { BrowserRouter } from "react-router-dom";


ReactDOM.createRoot(document.getElementById('root')!).render(
    <ConfigProvider locale={locale} theme={{
        algorithm: theme.defaultAlgorithm
    }}>
        <BrowserRouter>
            <App/>
        </BrowserRouter>
    </ConfigProvider>,
);
