package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type StableDiffusionRouter struct{}

func (s *StableDiffusionRouter) InitStableDiffusionRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	stableDiffusionRouter := Router.Group("stableDiffusion")
	stableDiffusionApi := api.ApiGroupApp.SystemApiGroup.StableDiffusionApi
	{
		stableDiffusionRouter.GET("getConfig", stableDiffusionApi.GetStableDiffusionConfigList)
	}
	{
		stableDiffusionRouter.POST("create", stableDiffusionApi.CreateStableDiffusionConfig)
	}
	return stableDiffusionRouter
}
