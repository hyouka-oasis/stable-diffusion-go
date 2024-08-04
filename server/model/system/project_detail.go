package system

import (
	"github/stable-diffusion-go/server/global"
)

type ProjectDetail struct {
	global.Model
	ProjectId             uint             `json:"projectId" form:"projectId"` // 项目Id
	FileName              string           `json:"fileName"`                   // 文件名称
	ParticipleConfig      ParticipleConfig `json:"participleConfig" form:"participleConfig"`
	InfoList              []Info           `json:"infoList"`
	StableDiffusionConfig string           `json:"stableDiffusionConfig" gorm:"type:json;default:\"{\n  \"width\":512,\n  \"height\": 512\n}\""` // api调用参数
	Language              string           `json:"language" gorm:"default:zh"`                                                                   //语言
	AudioConfig           AudioConfig      `json:"audioConfig"`
	BreakAudio            bool             `json:"breakAudio" gorm:"default:true"`  // 是否跳过存在的音频
	BatchAudio            bool             `json:"batchAudio" gorm:"default:false"` // 是否全量替换音频配置
}

func (ProjectDetail) TableName() string {
	return "project_detail"
}
