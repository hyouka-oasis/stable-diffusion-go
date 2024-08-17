package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type StableDiffusionRouter struct{}

func (s *StableDiffusionRouter) InitStableDiffusionRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	stableDiffusionRouter := Router.Group("sdapi/v1")
	stableDiffusionApi := api.ApiGroupApp.SystemApiGroup.StableDiffusionApi
	{
		stableDiffusionRouter.GET("sd-models", stableDiffusionApi.GetStableDiffusionSdModels)
		stableDiffusionRouter.GET("sd-vae", stableDiffusionApi.GetStableDiffusionSdVae)
		stableDiffusionRouter.GET("samplers", stableDiffusionApi.GetStableDiffusionSamplers)
		stableDiffusionRouter.GET("schedulers", stableDiffusionApi.GetStableDiffusionSchedulers)
		stableDiffusionRouter.GET("upscalers", stableDiffusionApi.GetStableDiffusionUpscalers)
	}
	return stableDiffusionRouter
}
