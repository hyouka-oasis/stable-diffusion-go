package request

import (
	"github/stable-diffusion-go/server/model/common/request"
)

type TaskPageInfoRequest struct {
	request.PageInfo
	ProjectDetailId uint `json:"projectDetailId"`
	TaskId          uint `json:"taskId"`
}
