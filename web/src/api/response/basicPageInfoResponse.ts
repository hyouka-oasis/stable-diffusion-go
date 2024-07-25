export interface BasicArrayResponses<T> {
    list: T[];
    total: number;
}


export interface BasicResponse {
    id: number;
    createdAt: Date;
    updatedAt: Date;
}
