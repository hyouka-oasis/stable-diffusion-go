package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type StableDiffusionLorasRouter struct{}

func (s *StableDiffusionLorasRouter) InitStableDiffusionLorasRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	stableDiffusionLorasRouter := Router.Group("sdapi/loras")
	stableDiffusionLorasApi := api.ApiGroupApp.SystemApiGroup.StableDiffusionLorasApi
	{
		stableDiffusionLorasRouter.GET("get", stableDiffusionLorasApi.GetStableDiffusionLorasList)
	}
	{
		stableDiffusionLorasRouter.POST("create", stableDiffusionLorasApi.CreateStableDiffusionLora)
	}
	return stableDiffusionLorasRouter
}
