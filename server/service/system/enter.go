package system

var taskService TaskService

type ServiceGroup struct {
	ProjectService
	ProjectDetailService
	InfoService
	TaskService

	AudioSrtService
	VideoService
	OllamaService

	StableDiffusionService
	StableDiffusionSettingsService
	StableDiffusionLorasService
	StableDiffusionImagesService

	SettingsService
}
