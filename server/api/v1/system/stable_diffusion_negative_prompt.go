package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type StableDiffusionNegativePromptApi struct{}

// GetStableDiffusionNegativePromptList 获取列表
func (s *StableDiffusionNegativePromptApi) GetStableDiffusionNegativePromptList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := stableDiffusionNegativePromptService.GetStableDiffusionNegativePrompt(pageInfo)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// CreateStableDiffusionNegativePrompt 创建反向关键词
func (s *StableDiffusionNegativePromptApi) CreateStableDiffusionNegativePrompt(c *gin.Context) {
	var stableDiffusionNegativePrompt system.StableDiffusionNegativePrompt
	err := c.ShouldBindJSON(&stableDiffusionNegativePrompt)
	if err != nil {
		global.Log.Error("请传入参数!")
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(stableDiffusionNegativePrompt, utils.StableDiffusionNegativePromptParamsVerify)
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

// UpdateStableDiffusionNegativePrompt 更新反向关键词
func (s *StableDiffusionNegativePromptApi) UpdateStableDiffusionNegativePrompt(c *gin.Context) {
	var stableDiffusionNegativePrompt system.StableDiffusionNegativePrompt
	err := c.ShouldBindJSON(&stableDiffusionNegativePrompt)
	if err != nil {
		global.Log.Error("请传入参数!")
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(stableDiffusionNegativePrompt, utils.IdVerify)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stableDiffusionNegativePromptService.UpdateStableDiffusionNegativePrompt(stableDiffusionNegativePrompt)
	if err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteStableDiffusionNegativePrompt 删除反向关键词
func (s *StableDiffusionNegativePromptApi) DeleteStableDiffusionNegativePrompt(c *gin.Context) {
	var params request.GetById
	err := c.ShouldBindJSON(&params)
	if err != nil {
		global.Log.Error("请传入参数!")
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(params, utils.IdVerify)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stableDiffusionNegativePromptService.DeleteStableDiffusionNegativePrompt(params.Id)
	if err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
