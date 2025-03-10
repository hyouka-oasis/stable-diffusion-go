package system

import (
	"github/stable-diffusion-go/server/global"
)

type StableDiffusionImages struct {
	global.Model
	InfoId          uint   `json:"infoId"`
	ProjectDetailId uint   `json:"projectDetailId"`            // 项目详情Id
	Name            string `json:"name" gorm:"comment:文件名"`    // 文件名
	Url             string `json:"url" gorm:"comment:文件地址"`    // 文件地址
	Tag             string `json:"tag" gorm:"comment:文件标签"`    // 文件标签
	Key             string `json:"key" gorm:"comment:编号"`      // 编号
	FileId          uint   `json:"fileId" gorm:"comment:文件Id"` // 编号
}

func (StableDiffusionImages) TableName() string {
	return "stable_diffusion_images"
}
