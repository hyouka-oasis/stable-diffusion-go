package request

import (
	"github/stable-diffusion-go/server/model/common/request"
)

type StableDiffusionRequestParams struct {
	request.GetById
	ProjectDetailId uint `json:"projectDetailId"`
}
