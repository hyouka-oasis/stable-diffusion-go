package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/common/response"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	systemResponse "github/stable-diffusion-go/server/model/system/response"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type StableDiffusionSettingsApi struct{}

// CreateStableDiffusionSettings 创建stabled-diffusion通用配置
func (s *StableDiffusionSettingsApi) CreateStableDiffusionSettings(c *gin.Context) {
	var params systemResponse.StableDiffusionSettingsResponse
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	if params.Name == "" {
		response.FailWithMessage("名称不能为空", c)
		return
	}
	err = stableDiffusionSettingsService.CreateStableDiffusionSettings(params)
	if err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateStableDiffusionSettings 更新stabled-diffusion通用配置
func (s *StableDiffusionSettingsApi) UpdateStableDiffusionSettings(c *gin.Context) {
	var params systemResponse.StableDiffusionSettingsResponse
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(params, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stableDiffusionSettingsService.UpdateStableDiffusionSettings(params)
	if err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteStableDiffusionSettings 删除stabled-diffusion通用配置
func (s *StableDiffusionSettingsApi) DeleteStableDiffusionSettings(c *gin.Context) {
	var params request.IdsReq
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(params, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stableDiffusionSettingsService.DeleteStableDiffusionSettings(params)
	if err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetStableDiffusionSettingsList 获取stable-diffusion-settings列表
func (s *StableDiffusionSettingsApi) GetStableDiffusionSettingsList(c *gin.Context) {
	var pageInfo systemRequest.StableDiffusionSettingsRequestParams
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := stableDiffusionSettingsService.GetStableDiffusionSettingsList(pageInfo.PageInfo)
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

// GetStableDiffusionSettings 获取stable-diffusion-settings
func (s *StableDiffusionSettingsApi) GetStableDiffusionSettings(c *gin.Context) {
	var params request.GetById
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(params, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	settings, err := stableDiffusionSettingsService.GetStableDiffusionSettings(params.Id)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(settings, "获取成功", c)
}
