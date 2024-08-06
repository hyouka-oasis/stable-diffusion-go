package system

import (
	"github/stable-diffusion-go/server/service"
)

type ApiGroup struct {
	ProjectApi
	ProjectDetailApi
	InfoApi
	SettingsApi
	StableDiffusionLorasApi
	StableDiffusionApi
	StableDiffusionNegativePromptApi
	AudioSrtApi
	VideoApi
}

var (
	projectService                       = service.ServiceGroupApp.SystemServiceGroup.ProjectService
	projectDetailService                 = service.ServiceGroupApp.SystemServiceGroup.ProjectDetailService
	infoService                          = service.ServiceGroupApp.SystemServiceGroup.InfoService
	settingsService                      = service.ServiceGroupApp.SystemServiceGroup.SettingsService
	stableDiffusionLorasService          = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionLorasService
	stableDiffusionService               = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionService
	stableDiffusionNegativePromptService = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionNegativePromptService
	audioSrtService                      = service.ServiceGroupApp.SystemServiceGroup.AudioSrtService
	videoService                         = service.ServiceGroupApp.SystemServiceGroup.VideoService
)
