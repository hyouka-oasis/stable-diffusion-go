export type BasicPageInfoRequest<T = any> = {
    page: number;
    pageSize: number;
} & {
    [P in keyof T]: T[P];
};
