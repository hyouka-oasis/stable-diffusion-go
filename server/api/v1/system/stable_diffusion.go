package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"go.uber.org/zap"
)

type StableDiffusionApi struct{}

// GetStableDiffusionSdModels 获取stable-diffusion模型列表
func (s *StableDiffusionApi) GetStableDiffusionSdModels(c *gin.Context) {
	list, err := stableDiffusionService.GetStableDiffusionSdModels()
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    int64(len(list)),
		Page:     -1,
		PageSize: 10,
	}, "获取成功", c)
}

// GetStableDiffusionSdVae 获取stable-diffusionVAE列表
func (s *StableDiffusionApi) GetStableDiffusionSdVae(c *gin.Context) {
	list, err := stableDiffusionService.GetStableDiffusionSdVae()
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    int64(len(list)),
		Page:     -1,
		PageSize: 10,
	}, "获取成功", c)
}

// GetStableDiffusionSamplers 获取stable-diffusion采样器
func (s *StableDiffusionApi) GetStableDiffusionSamplers(c *gin.Context) {
	list, err := stableDiffusionService.GetStableDiffusionSamplers()
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    int64(len(list)),
		Page:     -1,
		PageSize: 10,
	}, "获取成功", c)
}

// GetStableDiffusionSchedulers 获取stable-diffusion调度类型
func (s *StableDiffusionApi) GetStableDiffusionSchedulers(c *gin.Context) {
	list, err := stableDiffusionService.GetStableDiffusionSchedulers()
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    int64(len(list)),
		Page:     -1,
		PageSize: 10,
	}, "获取成功", c)
}
