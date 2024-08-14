package system

type StableDiffusionOverrideSettings struct {
	StableDiffusionSettingsId uint64 `json:"stableDiffusionSettingsId"`
	SdVae                     string `json:"sdVae" gorm:"comment:指定vae;default:'Automatic'"` // 指定vae
	SdModelCheckpoint         string `json:"sd_model_checkpoint" gorm:"comment:指定模型"`        // 指定模型
}

func (StableDiffusionOverrideSettings) TableName() string {
	return "stable_diffusion_override_settings"
}
