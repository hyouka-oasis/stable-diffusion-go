package v1

import (
	"github/stable-diffusion-go/server/api/v1/example"
	"github/stable-diffusion-go/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
