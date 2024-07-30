package system

import (
	"github/stable-diffusion-go/server/global"
)

type StableDiffusionConfigOverrideSettings struct {
	SDModelCheckpoint       string `json:"sd_model_checkpoint"`
	SDVae                   string `json:"sd_vae"`
	StableDiffusionConfigId uint   `json:"stableDiffusionConfigId"`
}

type StableDiffusionConfig struct {
	global.Model
	Width             int  `json:"width" gorm:"comment:图片宽度;default:512"`    // 图片宽度
	Height            int  `json:"height" gorm:"comment:图片高度;default:512"`   // 图片高度
	Seed              int  `json:"seed" gorm:"comment:图片高度;default:-1"`      // 随机数种子
	Steps             int  `json:"steps" gorm:"comment:图片高度;default:50"`     // 迭代步数
	BatchSize         int  `json:"batch_size" gorm:"comment:批次数量;default:1"` // 批次数量
	Niter             int  `json:"n_iter" gorm:"comment:图片高度;default:1"`     // 每次批量的数量
	CFGScale          int  `json:"cfg_scale"`                                // 提示词引导系数
	ClipSkip          int  `json:"clip_skip"`
	DenoisingStrength int  `json:"denoising_strength"` // 去噪强度
	DoNotSaveGrid     bool `json:"do_not_save_grid"`
	DoNotSaveSamples  bool `json:"do_not_save_samples"`
	// 高清修复
	EnableHr         bool   `json:"enable_hr"`
	Eta              int    `json:"eta"`
	HrNegativePrompt string `json:"hr_negative_prompt"`
	HrPrompt         string `json:"hr_prompt"`
	HrResizeX        int    `json:"hr_resize_x"`
	HrResizeY        int    `json:"hr_resize_y"`
	// 高清修复算法
	HrSamplerName string `json:"hr_sampler_name"`
	// 高清修复放大倍数
	HrScale int `json:"hr_scale"`
	// 高清修复步数
	HrSecondPassSteps int `json:"hr_second_pass_steps"`
	// 高清修复名称
	HrUpscaler string `json:"hr_upscaler"`
	// 反向提示词
	NegativePrompt   string                                `json:"negative_prompt"`
	OverrideSettings StableDiffusionConfigOverrideSettings `json:"override_settings"`
	// 默认 true
	OverrideSettingsRestoreAfterwards bool `json:"override_settings_restore_afterwards"`
	// 正向提示词
	Prompt          string `json:"prompt"`
	RestoreFaces    bool   `json:"restore_faces"` // 面部修复
	SChurn          int    `json:"s_churn"`
	SMinUncond      int    `json:"s_min_uncond"`
	SNoise          string `json:"s_noise"`
	STMax           string `json:"s_tmax"`
	STMin           string `json:"s_tmin"`
	SamplerIndex    int    `json:"sampler_index"`
	SamplerName     string `json:"sampler_name"`
	SaveImages      bool   `json:"save_images"`
	ScriptName      string `json:"script_name"`
	SeedResizeFromH int    `json:"seed_resize_from_h"`
	SeedResizeFromW int    `json:"seed_resize_from_w"`
	SendImages      bool   `json:"send_images"` // 默认 true
	SubSeed         int    `json:"subseed"`
	SubSeedStrength int    `json:"subseed_strength"`
	Tiling          bool   `json:"tiling"`
	ProjectDetailId uint   `json:"projectDetailId"`
}

func (StableDiffusionConfig) TableName() string {
	return "stable_diffusion_config"
}
