package system

import (
	"github/stable-diffusion-go/server/global"
)

var (
	START    = 1
	PENDING  = 2
	RESOLVED = 3
	REJECTED = 4
)

type Task struct {
	global.Model
	ProjectDetailId uint         `json:"projectDetailId"` // 项目详情Id
	Errors          []TaskErrors `json:"errors"`
	Progress        float64      `json:"progress"` // 进度
	Status          int          `json:"status"`   // 状态
	Message         string       `json:"message"`  // 消息文本
}

func (Task) TableName() string {
	return "task"
}
