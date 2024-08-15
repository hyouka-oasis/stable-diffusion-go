package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type StableDiffusionImagesApi struct{}

// StableDiffusionTextToImage 批量文件转图片
func (s *StableDiffusionImagesApi) StableDiffusionTextToImage(c *gin.Context) {
	var params systemRequest.StableDiffusionRequestParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(params, utils.StableDiffusionParamsVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	images, err := stableDiffusionImagesService.StableDiffusionTextToImage(params)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(images, "生成成功", c)
}

// DeleteStableDiffusionImage 删除图片
func (s *StableDiffusionImagesApi) DeleteStableDiffusionImage(c *gin.Context) {
	var params request.IdsReq
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(params, utils.IdsVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stableDiffusionImagesService.DeleteStableDiffusionImage(params)
	if err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// AddStableDiffusionImage 添加图片
func (s *StableDiffusionImagesApi) AddStableDiffusionImage(c *gin.Context) {
	var params system.StableDiffusionImages
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(params, utils.AddStableDiffusionImageVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stableDiffusionImagesService.AddStableDiffusionImage(params)
	if err != nil {
		global.Log.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}
