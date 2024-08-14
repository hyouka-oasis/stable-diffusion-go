package system

import (
	"github/stable-diffusion-go/server/global"
)

type StableDiffusionSettings struct {
	global.Model
	ProjectDetailId   uint                            `json:"infoId" gorm:"项目详情Id"`                   // 项目详情Id
	Name              string                          `json:"name" gorm:"comment:文件名"`                // 文件名
	Speed             int                             `json:"speed" gorm:"comment:随机数种子;default:-1"`  // 随机数种子
	Width             int                             `json:"width" gorm:"comment:图片宽度;default:512"`  // 图片宽度
	Height            int                             `json:"height" gorm:"comment:图片高度;default:512"` // 图片高度
	Prompt            string                          `json:"prompt" gorm:"comment:正向提示词"`            // 正向提示词
	NegativePrompt    string                          `json:"negative_prompt" gorm:"comment:反向提示词"`   // 反向提示词
	BatchSize         string                          `json:"batch_size" gorm:"comment:生成数量"`         // 生成数量
	Steps             string                          `json:"steps" gorm:"comment:迭代步数"`              // 迭代步数
	CfgScale          string                          `json:"cfg_scale" gorm:"comment:提示词引导系数"`       // 提示词引导系数
	SamplerName       string                          `json:"sampler_name" gorm:"comment:取样器"`        // 取样器
	OverrideSettings  StableDiffusionOverrideSettings `json:"overrideSettings"`
	EnableHr          bool                            `json:"enable_hr" gorm:"comment:是否开启高清修复;default:false"`                             // 是否开启高清修复
	HrUpscaler        string                          `json:"hr_upscaler" gorm:"comment:高清算法,Upscaler"`                                    // 高清算法
	HrSecondPassSteps float64                         `json:"hr_second_pass_steps" gorm:"comment:迭代步数,Hires steps"`                        // 迭代步数                              //
	DenoisingStrength float64                         `json:"denoising_strength" gorm:"comment:去噪强度,Denoising strength"`                   // 去噪强度                              //
	HrScale           float64                         `json:"hr_scale" gorm:"comment:放大倍数,Upscale by"`                                     // 放大倍数                              //
	HrResizeX         float64                         `json:"hr_resize_x" gorm:"comment:指定宽高,Resize width to"`                             // 指定宽高                              //
	HrResizeY         float64                         `json:"hr_resize_y" gorm:"comment:指定宽高,Resize width to"`                             // 指定宽高                              //
	HrPrompt          float64                         `json:"hr_prompt" gorm:"comment:高清下的正向提示词"`                                          // 高清下的正向提示词                              //
	HrNegativePrompt  float64                         `json:"hr_negative_prompt" gorm:"comment:高清下的反向提示词"`                                 // 高清下的反向提示词                              //
	Text              string                          `json:"text" gorm:"type:json;default:\"{\n  \"width\":512,\n  \"height\": 512\n}\""` // api调用参数
}

func (StableDiffusionSettings) TableName() string {
	return "stable_diffusion_settings"
}
