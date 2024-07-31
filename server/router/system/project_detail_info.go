package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type ProjectDetailParticipleListRouter struct{}

func (s *ProjectDetailParticipleListRouter) InitProjectDetailParticipleListRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	projectDetailInfoRouter := Router.Group("projectDetailInfo")
	projectDetailInfoApi := api.ApiGroupApp.SystemApiGroup.ProjectDetailInfoApi
	{
		projectDetailInfoRouter.GET("get", projectDetailInfoApi.GetProjectDetailInfo)
	}
	{
		projectDetailInfoRouter.POST("update", projectDetailInfoApi.UpdateProjectDetailInfo)
		projectDetailInfoRouter.POST("extractRole", projectDetailInfoApi.ExtractTheRoleProjectDetailInfoList)
		projectDetailInfoRouter.POST("translate", projectDetailInfoApi.TranslateProjectDetailInfoList)
		projectDetailInfoRouter.DELETE("delete", projectDetailInfoApi.DeleteProjectDetailInfo)
	}
	return projectDetailInfoRouter
}
