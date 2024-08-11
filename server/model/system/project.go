package system

import (
	"github/stable-diffusion-go/server/global"
)

type Project struct {
	global.Model
	Name string          `json:"name" gorm:"comment:项目名称"` // 项目名称
	List []ProjectDetail `json:"list"`                     // 项目列表详情
}

func (Project) TableName() string {
	return "project"
}
