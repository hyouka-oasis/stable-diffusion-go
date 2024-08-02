package request

import (
	"github/stable-diffusion-go/server/model/common/request"
)

type ProjectDetailRequestParams struct {
	request.GetById
	StableDiffusionImageIds []uint `json:"stableDiffusionImageIds"` // 文件Ids
}
