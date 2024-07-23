import http from "../utils/request"

export const getStableDiffusionList = () => {
    return http({
        method: "get",
        url: "/stableDiffusion/getConfig"
    })
}
