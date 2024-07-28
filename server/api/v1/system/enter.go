package system

import (
	"github/stable-diffusion-go/server/service"
)

type ApiGroup struct {
	ProjectApi
	ProjectDetailApi
	ProjectDetailParticipleListApi
	SettingsApi
	StableDiffusionApi
	StableDiffusionLorasApi
}

var (
	projectService                     = service.ServiceGroupApp.SystemServiceGroup.ProjectService
	projectDetailService               = service.ServiceGroupApp.SystemServiceGroup.ProjectDetailService
	projectDetailParticipleListService = service.ServiceGroupApp.SystemServiceGroup.ProjectDetailParticipleListService
	settingsService                    = service.ServiceGroupApp.SystemServiceGroup.SettingsService
	stableDiffusionService             = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionService
	stableDiffusionLorasService        = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionLorasService
)
