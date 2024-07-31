package system

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
)

type StableDiffusionNegativePromptService struct{}

// CreateStableDiffusionNegativePrompt 创建同用反向提示词
func (s *StableDiffusionNegativePromptService) CreateStableDiffusionNegativePrompt(stableDiffusionNegativePrompt system.StableDiffusionNegativePrompt) error {
	return global.DB.Create(&stableDiffusionNegativePrompt).Error
}
