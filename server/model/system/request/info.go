package request

import (
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/system"
)

type InfoCreateVideoRequest struct {
	request.IdsReq
	ProjectDetailId uint `json:"projectDetailId"`
}

type InfoTranslateRequest struct {
	system.Info
	TranslateType uint `json:"translateType"`
}
