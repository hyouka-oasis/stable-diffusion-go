package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type ProjectRouter struct{}

func (s *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	projectRouter := Router.Group("project")
	projectApi := api.ApiGroupApp.SystemApiGroup.ProjectApi
	{
		projectRouter.GET("list", projectApi.GetProjectList)
	}
	{
		projectRouter.POST("create", projectApi.CreateProject)
		projectRouter.DELETE("delete", projectApi.DeleteProject)
	}
	return projectRouter
}
