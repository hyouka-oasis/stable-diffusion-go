package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type StableDiffusionNegativePromptRouter struct{}

func (s *StableDiffusionNegativePromptRouter) InitStableDiffusionNegativePromptRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	stableDiffusionNegativePromptRouter := Router.Group("sdapi/prompt")
	stableDiffusionNegativePromptApi := api.ApiGroupApp.SystemApiGroup.StableDiffusionNegativePromptApi
	{
		stableDiffusionNegativePromptRouter.GET("negativePromptList", stableDiffusionNegativePromptApi.GetStableDiffusionNegativePromptList)
	}
	{
		stableDiffusionNegativePromptRouter.POST("createNegativePrompt", stableDiffusionNegativePromptApi.CreateStableDiffusionNegativePrompt)
		stableDiffusionNegativePromptRouter.POST("updateNegativePrompt", stableDiffusionNegativePromptApi.UpdateStableDiffusionNegativePrompt)
		stableDiffusionNegativePromptRouter.DELETE("deleteNegativePrompt", stableDiffusionNegativePromptApi.DeleteStableDiffusionNegativePrompt)
	}
	return stableDiffusionNegativePromptRouter
}
