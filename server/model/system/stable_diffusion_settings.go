package system

import (
	"github/stable-diffusion-go/server/global"
)

type StableDiffusionSettings struct {
	global.Model
	Name            string `json:"-"`
	ProjectDetailId uint   `json:"projectDetailId" gorm:"项目详情Id"` // 项目详情Id
	// 模型配置
	OverrideSettings StableDiffusionOverrideSettings `json:"override_settings"`                           // 模型配置
	ClipSkip         int                             `json:"clip_skip" gorm:"comment:Clip 跳过层;default:1"` // Clip 跳过层
	// 提示词配置
	Prompt         string `json:"prompt" gorm:"comment:正向提示词"`          // 正向提示词
	NegativePrompt string `json:"negative_prompt" gorm:"comment:反向提示词"` // 反向提示词
	// 生成配置
	SamplerName string `json:"sampler_name" gorm:"comment:采样器"`                 // 采样器
	Scheduler   string `json:"scheduler" gorm:"comment:调度类型;default:Automatic"` // 调度类型
	Steps       int    `json:"steps" gorm:"comment:迭代步数;default:20"`            // 迭代步数
	Width       int    `json:"width" gorm:"comment:图片宽度;default:512"`           // 图片宽度
	Height      int    `json:"height" gorm:"comment:图片高度;default:512"`          // 图片高度
	Niter       int    `json:"n_iter" gorm:"comment:生成批次;default:1"`            // 生成批次
	BatchSize   int    `json:"batch_size" gorm:"comment:生成数量;default:1"`        // 生成数量
	CfgScale    int    `json:"cfg_scale" gorm:"comment:提示词引导系数;default:7"`      // 提示词引导系数
	Seed        int    `json:"seed" gorm:"comment:随机数种子;default:-1"`            // 随机数种子
	// 高清修复配置
	EnableHr          bool    `json:"enable_hr" gorm:"comment:是否开启高清修复;default:false"`           // 是否开启高清修复
	HrUpscaler        string  `json:"hr_upscaler" gorm:"comment:高清算法,Upscaler"`                  // 高清算法
	HrSecondPassSteps float64 `json:"hr_second_pass_steps" gorm:"comment:迭代步数,Hires steps"`      // 迭代步数                              //
	DenoisingStrength float64 `json:"denoising_strength" gorm:"comment:去噪强度,Denoising strength"` // 去噪强度                              //
	HrScale           float64 `json:"hr_scale" gorm:"comment:放大倍数,Upscale by;default:1"`         // 放大倍数                              //
	HrResizeX         float64 `json:"hr_resize_x" gorm:"comment:指定宽高,Resize width to"`           // 指定宽高                              //
	HrResizeY         float64 `json:"hr_resize_y" gorm:"comment:指定宽高,Resize width to"`           // 指定宽高                              //
	HrPrompt          float64 `json:"hr_prompt" gorm:"comment:高清下的正向提示词"`                        // 高清下的正向提示词                              //
	HrNegativePrompt  float64 `json:"hr_negative_prompt" gorm:"comment:高清下的反向提示词"`               // 高清下的反向提示词                              //
}

func (StableDiffusionSettings) TableName() string {
	return "stable_diffusion_settings"
}
