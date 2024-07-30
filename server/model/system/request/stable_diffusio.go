package request

import (
	"github/stable-diffusion-go/server/model/common/request"
)

type StableDiffusionParams struct {
	request.IdsReq
	ProjectDetailId       uint   `json:"projectDetailId"`
	StableDiffusionConfig string `json:"stableDiffusionConfig"`
}
