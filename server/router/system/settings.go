package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type SettingsRouter struct{}

func (s *SettingsRouter) InitSettingsRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	settingsRouter := Router.Group("settings")
	settingsRouterApi := api.ApiGroupApp.SystemApiGroup.SettingsApi
	{
		settingsRouter.GET("get", settingsRouterApi.GetSettings)
	}
	{
		settingsRouter.POST("create", settingsRouterApi.CreateSettings)
		settingsRouter.POST("update", settingsRouterApi.UpdateSettings)
		//settingsRouter.DELETE("delete", settingsRouterApi.DeleteProject)
	}
	return settingsRouter
}
