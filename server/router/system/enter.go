package system

type RouterGroup struct {
	ProjectRouter
	ProjectDetailRouter
	InfoRouter
	BasicRouter
	TaskRouter

	AudioSrtRouter
	VideoRouter
	OllamaRouter

	StableDiffusionRouter
	StableDiffusionLorasRouter
	StableDiffusionImagesRouter
	StableDiffusionSettingsRouter

	SettingsRouter
}
