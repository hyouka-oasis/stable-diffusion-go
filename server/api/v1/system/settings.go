package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type SettingsApi struct{}

// CreateSettings 创建配置
func (s *SettingsApi) CreateSettings(c *gin.Context) {
	var settings system.Settings
	err := c.ShouldBindJSON(&settings)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = settingsService.CreateSettings(settings)
	if err != nil {
		global.Log.Error("新增失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// UpdateSettings 更新配置
func (s *SettingsApi) UpdateSettings(c *gin.Context) {
	var settings system.Settings
	err := c.ShouldBindJSON(&settings)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(settings, utils.SettingsVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(settings.StableDiffusionConfig, utils.StableDiffusionConfigVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if settings.TranslateType == "ollama" {
		err = utils.Verify(settings, utils.OllamaConfigVerify)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
	}
	err = settingsService.UpdateSettings(settings)
	if err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetSettings 获取配置
func (s *SettingsApi) GetSettings(c *gin.Context) {
	settings, err := settingsService.GetSettings()
	if err != nil {
		global.Log.Error("获取配置失败!", zap.Error(err))
		response.FailWithMessage("获取配置失败", c)
		return
	}
	response.OkWithDetailed(&settings, "获取配置成功", c)
}
