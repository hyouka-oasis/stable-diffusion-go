package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type ProjectDetailRouter struct{}

func (s *ProjectDetailRouter) InitProjectDetailRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	projectDetailRouter := Router.Group("projectDetail")
	projectDetailRouterApi := api.ApiGroupApp.SystemApiGroup.ProjectDetailApi
	{
		//projectDetailRouter.GET("list", projectDetailRouterApi.GetProjectList)
	}
	{
		projectDetailRouter.POST("upload", projectDetailRouterApi.UpdateProjectDetailFile)
	}
	return projectDetailRouter
}
