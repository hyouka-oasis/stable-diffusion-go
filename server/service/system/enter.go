package system

type ServiceGroup struct {
	ProjectService
	ProjectDetailService
	InfoService
	SettingsService
	StableDiffusionLorasService
	StableDiffusionService
	StableDiffusionNegativePromptService
	AudioSrtService
	VideoService
}
