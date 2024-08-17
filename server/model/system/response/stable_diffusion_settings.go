package response

import (
	"github/stable-diffusion-go/server/model/system"
)

type StableDiffusionSettingsResponse struct {
	Name string `json:"name" form:"name"`
	system.StableDiffusionSettings
}
