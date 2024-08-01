package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type StableDiffusionApi struct{}

// StableDiffusionTextToImage 批量文件转图片
func (s *StableDiffusionApi) StableDiffusionTextToImage(c *gin.Context) {
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
	images, err := stableDiffusionService.StableDiffusionTextToImage(params)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(images, "生成成功", c)
}
