package request

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
)

type AudioRequestParams struct {
	global.Model
	system.AudioConfig
}
