package service

import (
	"github/stable-diffusion-go/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
