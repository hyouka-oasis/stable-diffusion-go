package system

import (
	"github/stable-diffusion-go/server/service"
)

type ApiGroup struct {
	StableDiffusionApi
}

var (
	stableDiffusionService = service.ServiceGroupApp.SystemServiceGroup.StableDiffusionService
)
