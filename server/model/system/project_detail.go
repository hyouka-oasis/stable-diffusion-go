package system

import (
	"github/stable-diffusion-go/server/global"
)

type ProjectDetail struct {
	global.Model
	ProjectId             uint             `json:"projectId" form:"projectId"` // 项目Id
	FileName              string           `json:"fileName"`                   // 文件名称
	ParticipleConfig      ParticipleConfig `json:"participleConfig" form:"participleConfig"`
	InfoList              []Info           `json:"infoList"`
	StableDiffusionConfig string           `json:"stableDiffusionConfig" gorm:"type:json"` // api调用参数
	Language              string           `json:"language"`                               //语言
	AudioConfig           AudioConfig      `json:"audioConfig"`
}

func (ProjectDetail) TableName() string {
	return "project_detail"
}
