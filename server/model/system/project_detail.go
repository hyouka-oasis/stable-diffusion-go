package system

import (
	"github/stable-diffusion-go/server/global"
)

type ProjectDetail struct {
	global.Model
	ProjectId      uint                          `json:"projectId" form:"projectId"` // 项目Id
	FileName       string                        `json:"fileName"`                   // 文件名称
	Participle     ProjectDetailParticiple       `json:"participle" form:"participle"`
	ParticipleList []ProjectDetailParticipleList `json:"participleList"`
}

func (ProjectDetail) TableName() string {
	return "project_detail"
}
