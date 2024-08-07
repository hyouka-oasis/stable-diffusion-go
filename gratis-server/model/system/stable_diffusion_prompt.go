package system

import (
	"github/stable-diffusion-go/server/global"
)

type StableDiffusionNegativePrompt struct {
	global.Model
	Name string `json:"name"` //反向提示词名字
	Text string `json:"text"` //反向提示词
}

func (StableDiffusionNegativePrompt) TableName() string {
	return "stable_diffusion_negative_prompt"
}
