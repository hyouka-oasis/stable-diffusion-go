package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type StableDiffusionSettingsRouter struct{}

func (s *StableDiffusionSettingsRouter) InitStableDiffusionSettingsRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	stableDiffusionSettingsRouter := Router.Group("sdapi/settings")
	stableDiffusionSettingsApi := api.ApiGroupApp.SystemApiGroup.StableDiffusionSettingsApi
	{
		stableDiffusionSettingsRouter.GET("get", stableDiffusionSettingsApi.GetStableDiffusionSettingsList)
	}
	{
		stableDiffusionSettingsRouter.POST("detail", stableDiffusionSettingsApi.GetStableDiffusionSettings)
		stableDiffusionSettingsRouter.POST("create", stableDiffusionSettingsApi.CreateStableDiffusionSettings)
		stableDiffusionSettingsRouter.POST("update", stableDiffusionSettingsApi.UpdateStableDiffusionSettings)
		stableDiffusionSettingsRouter.DELETE("delete", stableDiffusionSettingsApi.DeleteStableDiffusionSettings)
	}
	return stableDiffusionSettingsRouter
}
