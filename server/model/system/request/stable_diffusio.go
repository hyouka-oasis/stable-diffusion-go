package request

import (
	"github/stable-diffusion-go/server/model/common/request"
)

type StableDiffusionParams struct {
	request.GetById
	ProjectDetailId uint `json:"projectDetailId"`
}
