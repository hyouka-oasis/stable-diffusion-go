package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type ProjectDetailRouter struct{}

func (s *ProjectDetailRouter) InitProjectDetailRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	projectDetailRouter := Router.Group("projectDetail")
	projectDetailApi := api.ApiGroupApp.SystemApiGroup.ProjectDetailApi
	{
		projectDetailRouter.GET("get", projectDetailApi.GetProjectDetail)
	}
	{
		projectDetailRouter.POST("create", projectDetailApi.CreateProjectDetail)
		projectDetailRouter.POST("upload", projectDetailApi.UploadProjectDetailFile)
		projectDetailRouter.POST("update", projectDetailApi.UpdateProjectDetail)
		projectDetailRouter.DELETE("delete", projectDetailApi.DeleteProjectDetail)
	}
	return projectDetailRouter
}
