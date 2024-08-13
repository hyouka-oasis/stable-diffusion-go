package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type OllamaRouter struct{}

func (s *OllamaRouter) InitOllamaRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	ollamaRouter := Router.Group("ollama")
	ollamaApi := api.ApiGroupApp.SystemApiGroup.OllamaApi
	{
		ollamaRouter.GET("get", ollamaApi.GetOllamaModelList)
	}
	return ollamaRouter
}
