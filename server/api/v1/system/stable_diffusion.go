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

type StableDiffusionApi struct{}

// GetStableDiffusionConfig 获取stable-diffusion配置列表
func (s *StableDiffusionApi) GetStableDiffusionConfig(c *gin.Context) {
	var pageInfo systemRequest.StableDiffusionQueryParams
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
	list, total, err := stableDiffusionService.GetStableDiffusionList(pageInfo.PageInfo)
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

// CreateStableDiffusionConfig 创建stable-diffusion配置参数
func (s *StableDiffusionApi) CreateStableDiffusionConfig(c *gin.Context) {
	var stableDiffusionConfig system.StableDiffusion
	err := c.ShouldBindJSON(&stableDiffusionConfig)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(stableDiffusionConfig, utils.StableDiffusionVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = stableDiffusionService.CreateStableDiffusionConfig(stableDiffusionConfig)
	if err != nil {
		global.Log.Error("新增失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}
