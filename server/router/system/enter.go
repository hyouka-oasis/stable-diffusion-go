package system

type RouterGroup struct {
	ProjectRouter
	ProjectDetailRouter
	InfoRouter
	BasicRouter

	AudioSrtRouter
	VideoRouter
	OllamaRouter

	StableDiffusionLorasRouter
	StableDiffusionRouter
	StableDiffusionNegativePromptRouter
	StableDiffusionSettingsRouter

	SettingsRouter
}
