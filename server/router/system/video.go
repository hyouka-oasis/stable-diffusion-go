package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type VideoRouter struct{}

func (s *VideoRouter) InitVideoRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	videoRouter := Router.Group("video")
	videoApi := api.ApiGroupApp.SystemApiGroup.VideoApi
	{
		videoRouter.POST("create", videoApi.CreateVideo)
	}
	return videoRouter
}
