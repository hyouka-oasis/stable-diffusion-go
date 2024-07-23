package router

import (
	"github/stable-diffusion-go/server/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
