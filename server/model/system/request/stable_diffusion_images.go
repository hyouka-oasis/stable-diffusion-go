package request

import (
	"github/stable-diffusion-go/server/model/common/request"
)

type StableDiffusionRequestParams struct {
	request.IdsReq
	ProjectDetailId uint `json:"projectDetailId"`
}
