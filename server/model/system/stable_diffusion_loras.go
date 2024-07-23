package system

import (
	"github/stable-diffusion-go/server/global"
)

type StableDiffusionLoras struct {
	global.Model
	Name              string `json:"name" gorm:"comment:lora标签"`            // lora
	Alias             string `json:"alias" gorm:"comment:lora名字"`           // lora
	Image             string `json:"image" gorm:"comment:图片"`               // lora
	StableDiffusionId uint   `json:"stableDiffusionId" gorm:"comment:父级id"` // lora
}

func (StableDiffusionLoras) TableName() string {
	return "stable_diffusion_loras"
}
