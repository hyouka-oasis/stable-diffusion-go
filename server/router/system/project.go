package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type ProjectRouter struct{}

func (s *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	projectRouter := Router.Group("project")
	projectRouterApi := api.ApiGroupApp.SystemApiGroup.ProjectApi
	{
		projectRouter.GET("list", projectRouterApi.GetProjectList)
	}
	{
		projectRouter.POST("create", projectRouterApi.CreateProject)
		projectRouter.DELETE("delete", projectRouterApi.DeleteProject)
	}
	return projectRouter
}
