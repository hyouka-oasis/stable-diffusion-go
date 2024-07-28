package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type SettingsRouter struct{}

func (s *SettingsRouter) InitSettingsRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	settingsRouter := Router.Group("settings")
	settingsApi := api.ApiGroupApp.SystemApiGroup.SettingsApi
	{
		settingsRouter.GET("get", settingsApi.GetSettings)
	}
	{
		settingsRouter.POST("create", settingsApi.CreateSettings)
		settingsRouter.POST("update", settingsApi.UpdateSettings)
		//settingsRouter.DELETE("delete", settingsRouterApi.DeleteProject)
	}
	return settingsRouter
}
