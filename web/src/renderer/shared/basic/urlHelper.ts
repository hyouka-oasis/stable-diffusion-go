//获取浏览器地址指定参数
const getQueryString = function (name) {
    const reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    const r = window.location.search.substr(1).match(reg);
    if (r != null) return r[2];
    return null;
};

// 获取浏览器地址所有参数
const getQueryData = <T>(url?: string) => {
    const r = (url ?? window.location.href).split("?")[1],
        obj = {};
    if (r) {
        r.split("&").forEach(function (value) {
            const arr = value.split("=");
            obj[arr[0]] = arr[1];
        });
    }
    return r ? obj as T : null;
};
/**
 * 对象转换params
 * @param query
 */
const objectToQueryParams = <T>(query?: {
    [key in keyof T]: any;
}) => {
    let queryParams = "";
    let index = 0;
    if (query) {
        for (const key of Object.keys(query)) {
            queryParams += `${index ? "&" : "?"}${key}=${query[key]}`;
            index++;
        }
    }
    return queryParams;
};

export {
    getQueryString,
    getQueryData,
    objectToQueryParams
};
