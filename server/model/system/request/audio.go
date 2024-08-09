package request

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
)

type AudioSrtRequestParams struct {
	global.Model
	system.AudioConfig
}
