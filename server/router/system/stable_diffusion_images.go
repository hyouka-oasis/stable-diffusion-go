package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type StableDiffusionImagesRouter struct{}

func (s *StableDiffusionImagesRouter) InitStableDiffusionImagesRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	stableDiffusionImagesRouter := Router.Group("sdapi/images")
	stableDiffusionImagesApi := api.ApiGroupApp.SystemApiGroup.StableDiffusionImagesApi
	{
		stableDiffusionImagesRouter.POST("text2image", stableDiffusionImagesApi.StableDiffusionTextToImage)
		stableDiffusionImagesRouter.POST("addImage", stableDiffusionImagesApi.AddStableDiffusionImage)
		stableDiffusionImagesRouter.DELETE("deleteImage", stableDiffusionImagesApi.DeleteStableDiffusionImage)
	}
	return stableDiffusionImagesRouter
}
