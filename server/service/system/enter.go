package system

type ServiceGroup struct {
	ProjectService
	ProjectDetailService
	InfoService

	AudioSrtService
	VideoService
	OllamaService

	StableDiffusionService
	StableDiffusionSettingsService
	StableDiffusionLorasService
	StableDiffusionImagesService

	SettingsService
}
