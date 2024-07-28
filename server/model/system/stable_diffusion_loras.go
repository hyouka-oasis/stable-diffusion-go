package system

import (
	"github/stable-diffusion-go/server/global"
)

type StableDiffusionLoras struct {
	global.Model
	Name    string `json:"name" gorm:"comment:lora标签"` // lora
	Roles   string `json:"roles"`                      //对应的角色名称
	ImageId uint   `json:"imageId"`                    // 图片Id
	Url     string `json:"url"`
}

func (StableDiffusionLoras) TableName() string {
	return "stable_diffusion_loras"
}
