package system

import (
	"github/stable-diffusion-go/server/global"
)

type ProjectDetailParticipleList struct {
	global.Model
	ProjectDetailId uint   `json:"projectDetailId"` // 项目详情Id
	Text            string `json:"text"`            // 文本
	Prompt          string `json:"prompt"`          // 文本
	Character       string `json:"character"`       // 人物
}

func (ProjectDetailParticipleList) TableName() string {
	return "project_detail_participle_list"
}
