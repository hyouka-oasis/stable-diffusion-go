package system

import (
	"github/stable-diffusion-go/server/global"
)

type Info struct {
	global.Model
	ProjectDetailId        uint                    `json:"projectDetailId"`              // 项目详情Id
	Text                   string                  `json:"text"`                         // 文本
	Prompt                 string                  `json:"prompt"`                       // 正向提示词
	NegativePrompt         string                  `json:"negativePrompt"`               // 反向提示词
	Role                   string                  `json:"role"`                         // 人物
	StableDiffusionImages  []StableDiffusionImages `json:"stableDiffusionImages"`        // 生成的图片
	StableDiffusionImageId uint                    `json:"stableDiffusionImageId"`       // 选定的图片 默认第一张
	AudioConfig            AudioConfig             `json:"audioConfig"`                  // 音频设定
	Loading                bool                    `json:"loading" gorm:"default:false"` // 是否加载状态
}

func (Info) TableName() string {
	return "info"
}
