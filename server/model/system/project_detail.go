package system

import (
	"github/stable-diffusion-go/server/global"
)

type ProjectDetail struct {
	global.Model
	ProjectId             uint                    `json:"projectId" form:"projectId"` // 项目Id
	FileName              string                  `json:"fileName"`                   // 文件名称
	ParticipleConfig      ProjectDetailParticiple `json:"participleConfig" form:"participleConfig"`
	ProjectDetailInfoList []ProjectDetailInfo     `json:"projectDetailInfoList"`
	StableDiffusionConfig string                  `json:"stableDiffusionConfig" gorm:"type:json"` // api调用参数
}

func (ProjectDetail) TableName() string {
	return "project_detail"
}
