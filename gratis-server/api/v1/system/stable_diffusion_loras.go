package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type StableDiffusionLorasApi struct{}

// GetStableDiffusionLorasList 获取stable-diffusion-loras配置列表
func (s *StableDiffusionLorasApi) GetStableDiffusionLorasList(c *gin.Context) {
	var pageInfo systemRequest.StableDiffusionLorasRequestParams
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
	list, total, err := stableDiffusionLorasService.GetStableDiffusionLorasList(pageInfo.PageInfo)
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

// CreateStableDiffusionLora 创建loras
func (s *StableDiffusionLorasApi) CreateStableDiffusionLora(c *gin.Context) {
	var stableDiffusionLoras system.StableDiffusionLoras
	err := c.ShouldBindJSON(&stableDiffusionLoras)
	if err != nil {
		global.Log.Error("请传入参数!")
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(stableDiffusionLoras, utils.StableDiffusionLorasVerify)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stableDiffusionLorasService.CreateStableDiffusionLoras(stableDiffusionLoras)
	if err != nil {
		global.Log.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}
