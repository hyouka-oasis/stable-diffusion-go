package system

import (
	"github/stable-diffusion-go/server/global"
)

type ProjectFormDetail struct {
	global.Model
	ProjectDetailId uint   `json:"projectDetailId"` // 项目详情Id
	Text            string `json:"text"`            // 文本
}

type ProjectDetailPotential struct {
	ProjectDetailId uint `json:"projectDetailId"`                             // 项目详情Id
	MinLength       int  `json:"minLength" form:"minLength" gorm:"default:2"` // 名称最小长度
	MaxLength       int  `json:"maxLength" form:"maxLength" gorm:"default:4"` // 名称最大长度
	MinWords        int  `json:"minWords" form:"minWords" gorm:"default:30"`  // 每张图片的最小文字数量
	MaxWords        int  `json:"maxWords" form:"maxWords" gorm:"default:40"`  // 每张图片的最大文字数量
}

type ProjectDetail struct {
	global.Model
	ProjectId uint                   `json:"projectId" form:"projectId"` // 项目Id
	FileName  string                 `json:"fileName"`                   // 文件名称
	Potential ProjectDetailPotential `json:"potential" form:"potential"`
	FormList  []ProjectFormDetail    `json:"formList"`
}

func (ProjectDetail) TableName() string {
	return "project_detail"
}
