package system

type RouterGroup struct {
	ProjectRouter
	ProjectDetailRouter
	InfoRouter
	BasicRouter

	AudioSrtRouter
	VideoRouter
	OllamaRouter

	StableDiffusionRouter
	StableDiffusionLorasRouter
	StableDiffusionImagesRouter
	StableDiffusionSettingsRouter

	SettingsRouter
}
