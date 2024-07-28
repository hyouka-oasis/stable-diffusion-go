package service

import (
	"github/stable-diffusion-go/server/service/example"
	"github/stable-diffusion-go/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
