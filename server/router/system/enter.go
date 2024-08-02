package system

type RouterGroup struct {
	ProjectRouter
	ProjectDetailRouter
	InfoRouter
	SettingsRouter
	StableDiffusionLorasRouter
	StableDiffusionRouter
	StableDiffusionNegativePromptRouter
	AudioSrtRouter
}
