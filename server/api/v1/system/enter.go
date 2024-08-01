package system

import (
	"github/stable-diffusion-go/server/service"
)

type ApiGroup struct {
	ProjectApi
	ProjectDetailApi
	ProjectDetailInfoApi
	SettingsApi
	StableDiffusionLorasApi
	StableDiffusionApi
	StableDiffusionNegativePromptApi
	AudioSrtApi
}

var (
	projectService                       = service.ServiceGroupApp.SystemServiceGroup.ProjectService
	projectDetailService                 = service.ServiceGroupApp.SystemServiceGroup.ProjectDetailService
	projectDetailParticipleListService   = service.ServiceGroupApp.SystemServiceGroup.ProjectDetailParticipleInfoService
	settingsService                      = service.ServiceGroupApp.SystemServiceGroup.SettingsService
	stableDiffusionLorasService          = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionLorasService
	stableDiffusionService               = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionService
	stableDiffusionNegativePromptService = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionNegativePromptService
	audioSrtService                      = service.ServiceGroupApp.SystemServiceGroup.AudioSrtService
)
