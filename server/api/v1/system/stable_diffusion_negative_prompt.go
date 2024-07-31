package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type StableDiffusionNegativePromptApi struct{}

func (s *StableDiffusionNegativePromptApi) CreateStableDiffusionNegativePrompt(c *gin.Context) {
	var stableDiffusionNegativePrompt system.StableDiffusionNegativePrompt
	err := c.ShouldBindJSON(&stableDiffusionNegativePrompt)
	if err != nil {
		global.Log.Error("请传入参数!")
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(stableDiffusionNegativePrompt, utils.StableDiffusionNegativePromptVerify)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stableDiffusionNegativePromptService.CreateStableDiffusionNegativePrompt(stableDiffusionNegativePrompt)
	if err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
