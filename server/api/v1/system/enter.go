package system

import (
	"github/stable-diffusion-go/server/service"
)

type ApiGroup struct {
	ProjectApi
	ProjectDetailApi
	InfoApi

	AudioSrtApi
	VideoApi
	OllamaApi
	BasicApi

	StableDiffusionApi
	StableDiffusionSettingsApi
	StableDiffusionLorasApi
	StableDiffusionImagesApi

	SettingsApi
}

var (
	projectService       = service.ServiceGroupApp.SystemServiceGroup.ProjectService
	projectDetailService = service.ServiceGroupApp.SystemServiceGroup.ProjectDetailService
	infoService          = service.ServiceGroupApp.SystemServiceGroup.InfoService

	audioSrtService = service.ServiceGroupApp.SystemServiceGroup.AudioSrtService
	videoService    = service.ServiceGroupApp.SystemServiceGroup.VideoService
	ollamaService   = service.ServiceGroupApp.SystemServiceGroup.OllamaService

	stableDiffusionService         = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionService
	stableDiffusionSettingsService = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionSettingsService
	stableDiffusionLorasService    = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionLorasService
	stableDiffusionImagesService   = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionImagesService

	settingsService = service.ServiceGroupApp.SystemServiceGroup.SettingsService
)
