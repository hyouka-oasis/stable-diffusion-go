package router

import (
	"github/stable-diffusion-go/server/router/example"
	"github/stable-diffusion-go/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
