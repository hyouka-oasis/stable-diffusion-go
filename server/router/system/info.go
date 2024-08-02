package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type InfoRouter struct{}

func (s *InfoRouter) InitInfoRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	infoRouter := Router.Group("info")
	infoApi := api.ApiGroupApp.SystemApiGroup.ProjectDetailInfoApi
	{
		infoRouter.GET("get", infoApi.GetInfo)
	}
	{
		infoRouter.POST("update", infoApi.UpdateInfo)
		infoRouter.POST("extractRole", infoApi.ExtractTheInfoRole)
		infoRouter.POST("translate", infoApi.TranslateInfoPrompt)
		infoRouter.DELETE("delete", infoApi.DeleteInfo)
	}
	return infoRouter
}
