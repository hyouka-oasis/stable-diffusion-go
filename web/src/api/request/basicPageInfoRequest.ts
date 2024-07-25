export type BasicPageInfoRequest<T> = {
    page: number;
    pageSize: number;
} & {
    [P in keyof T]: T[P];
};
