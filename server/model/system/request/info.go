package request

import (
	"github/stable-diffusion-go/server/model/common/request"
)

type InfoCreateVideoRequest struct {
	request.IdsReq
	ProjectDetailId uint `json:"projectDetailId"`
}
