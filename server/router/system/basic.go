package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type BasicRouter struct{}

func (s *BasicRouter) InitBasicRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	basicRouter := Router.Group("basic")
	basicApi := api.ApiGroupApp.SystemApiGroup.BasicApi
	{
		basicRouter.GET("exit", basicApi.ExitGin)
	}
	return basicRouter
}
