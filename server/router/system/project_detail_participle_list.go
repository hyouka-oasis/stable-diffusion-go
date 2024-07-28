package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type ProjectDetailParticipleListRouter struct{}

func (s *ProjectDetailParticipleListRouter) InitProjectDetailParticipleListRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	projectDetailParticipleListRouter := Router.Group("projectDetailParticipleList")
	projectDetailParticipleListApi := api.ApiGroupApp.SystemApiGroup.ProjectDetailParticipleListApi
	{
	}
	{
		projectDetailParticipleListRouter.POST("extractCharacter", projectDetailParticipleListApi.ExtractTheCharacterProjectDetailParticipleList)
		projectDetailParticipleListRouter.POST("translate", projectDetailParticipleListApi.TranslateProjectDetailParticipleList)
		projectDetailParticipleListRouter.DELETE("delete", projectDetailParticipleListApi.DeleteProjectDetailParticipleListItem)
	}
	return projectDetailParticipleListRouter
}
