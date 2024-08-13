package system

type ServiceGroup struct {
	ProjectService
	ProjectDetailService
	InfoService

	AudioSrtService
	VideoService
	OllamaService

	StableDiffusionSettingsService
	StableDiffusionLorasService
	StableDiffusionService
	StableDiffusionNegativePromptService

	SettingsService
}
