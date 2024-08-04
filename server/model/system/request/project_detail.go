package request

import (
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/system"
)

type ProjectDetailRequestParams struct {
	request.GetById
	StableDiffusionImageIds []uint `json:"stableDiffusionImageIds"` // 文件Ids
}

type UpdateProjectDetailRequestParams struct {
	system.ProjectDetail
}
