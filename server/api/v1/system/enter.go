package system

import (
	"github/stable-diffusion-go/server/service"
)

type ApiGroup struct {
	StableDiffusionApi
	ProjectApi
	ProjectDetailApi
	ProjectDetailParticipleListApi
	SettingsApi
}

var (
	stableDiffusionService             = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionService
	projectService                     = service.ServiceGroupApp.SystemServiceGroup.ProjectService
	projectDetailService               = service.ServiceGroupApp.SystemServiceGroup.ProjectDetailService
	projectDetailParticipleListService = service.ServiceGroupApp.SystemServiceGroup.ProjectDetailParticipleListService
	settingsService                    = service.ServiceGroupApp.SystemServiceGroup.SettingsService
)
