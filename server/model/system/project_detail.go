package system

import (
	"github/stable-diffusion-go/server/global"
)

type ProjectFormDetail struct {
	ProjectDetailId uint   `json:"projectDetailId"` // 项目详情Id
	Text            string `json:"text"`            // 文本
}

type ProjectDetailPotential struct {
	ProjectDetailId uint `json:"projectDetailId"`            // 项目详情Id
	MinLength       int  `json:"minLength" gorm:"default:2"` // 名称最小长度
	MaxLength       int  `json:"maxLength" gorm:"default:4"` // 名称最大长度
	MinWords        int  `json:"minWords" gorm:"default:30"` // 每张图片的最小文字数量
	MaxWords        int  `json:"maxWords" gorm:"default:40"` // 每张图片的最大文字数量
}

type ProjectDetail struct {
	global.Model
	ProjectId uint                   `json:"projectId"` // 项目Id
	FileName  string                 `json:"fileName"`  // 文件名称
	Potential ProjectDetailPotential `json:"potential"`
	FormList  []ProjectFormDetail    `json:"formList"`
}

func (ProjectDetail) TableName() string {
	return "project_detail"
}
