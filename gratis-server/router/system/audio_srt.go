package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type AudioSrtRouter struct{}

func (s *AudioSrtRouter) InitAudioSrtRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	audioRouter := Router.Group("audioSrt")
	audioApi := api.ApiGroupApp.SystemApiGroup.AudioSrtApi
	{
		audioRouter.POST("create", audioApi.CreateAudioAndSrt)
	}
	return audioRouter
}
