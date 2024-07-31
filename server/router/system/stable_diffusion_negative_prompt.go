package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type StableDiffusionNegativePromptRouter struct{}

func (s *StableDiffusionNegativePromptRouter) InitStableDiffusionNegativePromptRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	stableDiffusionNegativePromptRouter := Router.Group("stableDiffusion")
	stableDiffusionNegativePromptApi := api.ApiGroupApp.SystemApiGroup.StableDiffusionNegativePromptApi
	{
		stableDiffusionNegativePromptRouter.POST("createNegativePrompt", stableDiffusionNegativePromptApi.CreateStableDiffusionNegativePrompt)
	}
	return stableDiffusionNegativePromptRouter
}
