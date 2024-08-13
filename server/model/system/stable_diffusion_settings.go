package system

import (
	"github/stable-diffusion-go/server/global"
)

type StableDiffusionSettings struct {
	global.Model
	Name string `json:"name" gorm:"comment:文件名"`                                                     // 文件名
	Text string `json:"text" gorm:"type:json;default:\"{\n  \"width\":512,\n  \"height\": 512\n}\""` // api调用参数
}

func (StableDiffusionSettings) TableName() string {
	return "stable_diffusion_settings"
}
