package system

import (
	"github/stable-diffusion-go/server/service"
)

type ApiGroup struct {
	StableDiffusionApi
	ProjectApi
	ProjectDetailApi
	SettingsApi
}

var (
	stableDiffusionService = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionService
	projectService         = service.ServiceGroupApp.SystemServiceGroup.ProjectService
	projectDetailService   = service.ServiceGroupApp.SystemServiceGroup.ProjectDetailService
	settingsService        = service.ServiceGroupApp.SystemServiceGroup.SettingsService
)
