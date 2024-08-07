import axios, { AxiosRequestConfig } from 'axios';
import { message } from "antd";
import { getQueryData } from "renderer/shared/basic/urlHelper";

export const getServerPort = (isDebug?: boolean) => {
    if (isDebug) {
        return "8881";
    }
    const urlPort = getQueryData()?.["server_port"];
    if (urlPort) {
        return urlPort.replace(/[^0-9]/g, "");
    }
    if (window["server_port"]) {
        return window["server_port"];
    }
};

export const host = (customPort?: string) => {
    const port = customPort ?? getServerPort();
    if (process.env.NODE_ENV === "development") {
        return `http://localhost:${port}`;
    }
    return `http://127.0.0.1:${port}`;
};


// 1. 创建axios实例
const instance = axios.create({
    baseURL: host(),
    // timeout: 5000,// 请求超时时间
    headers: { //设置请求头
        "Content-Type": "application/x-www-form-urlencoded;charset=utf-8",
    },
});

const errorMessageFun = (status: number, errorMessage: string) => {
    switch (status) {
        // 401: 未登录
        // 未登录则跳转登录页面，并携带当前页面的路径
        // 在登录成功后返回当前页面，这一步需要在登录页操作。
        case 401:
            // router.replace({
            //     path: '/login',
            //     query: { redirect: router.currentRoute.fullPath }
            // });
            break;
        // 403 token过期
        // 登录过期对用户进行提示
        // 清除本地token和清空vuex中token对象
        // 跳转登录页面
        case 403:
            message.error({
                content: '登录过期，请重新登录',
                duration: 10,
            });
            break;
        // 404请求不存在
        case 404:
            message.error({
                content: '网络请求不存在',
                duration: 10,
            });
            break;
        // 其他错误，直接抛出错误提示
        default:
            message.error({
                content: errorMessage,
                duration: 10,
            });
    }
};

// 请求拦截器
instance.interceptors.request.use(
    config => {
        // 每次发送请求之前判断是否存在token，如果存在，则统一在http请求的header都加上token，不用每次请求都手动添加了
        // 即使本地存在token，也有可能token是过期的，所以在响应拦截器中要对返回状态进行判断
        // const token = store.state.token;
        // token && (config.headers.Authorization = token);
        return config;
    },
    error => {
        return Promise.reject(error);
    });

// 响应拦截器
instance.interceptors.response.use(
    response => {
        if (response.status === 200) {
            if (response.data.code !== 200) {
                errorMessageFun(response.data.code, response.data.message);
                return Promise.reject(response.data.message);
            }
            return Promise.resolve(response.data.data);
        } else {
            return Promise.reject(response);
        }
    },
    // 服务器状态码不是200的情况
    error => {
        if (error.response.status) {

            return Promise.reject(error.response);
        }
    }
);

const postFormApi = <T extends any>(postConfig: {
    url: string,
    data?: any,
    config?: AxiosRequestConfig
}): Promise<T> => {
    const formData = new FormData();
    for (const key in postConfig.data) {
        formData.append(key, postConfig.data[key]);
    }
    return instance.post(postConfig.url, formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        ...postConfig.config,
    });
};

const postApi = <T extends any>(postConfig: {
    url: string,
    data?: any,
    config?: AxiosRequestConfig
}): Promise<T> => {
    return instance.post(postConfig.url, postConfig.data, {
        headers: {
            "Content-Type": "application/json;charset=utf-8",
        },
        ...postConfig.config,
    });
};

const deleteApi = <T extends any>(postConfig: {
    url: string,
    data?: any,
    config?: AxiosRequestConfig
}): Promise<T> => {
    return instance.delete(postConfig.url, {
        data: postConfig.data,
        headers: {
            "Content-Type": "application/json;charset=utf-8",
        },
        ...postConfig.config,
    });
};

const getApi = <T extends any>(getConfig: {
    url: string,
    data?: any,
    config?: AxiosRequestConfig
}): Promise<T> => {
    return instance.get(getConfig.url, {
        params: getConfig.data,
        headers: {
            "Content-Type": "application/x-www-form-urlencoded;charset=utf-8",
        },
        ...getConfig.config,
    });
};

export {
    postApi,
    getApi,
    deleteApi,
    postFormApi,
};
