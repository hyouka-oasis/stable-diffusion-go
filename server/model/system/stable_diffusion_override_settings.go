package system

type StableDiffusionOverrideSettings struct {
	ProjectDetailId           uint   `json:"projectDetailId"`
	StableDiffusionSettingsId uint   `json:"stableDiffusionSettingsId"`
	SdVae                     string `json:"sd_vae" gorm:"comment:指定vae"`             // 指定vae
	SdModelCheckpoint         string `json:"sd_model_checkpoint" gorm:"comment:指定模型"` // 指定模型
}

func (StableDiffusionOverrideSettings) TableName() string {
	return "stable_diffusion_override_settings"
}
