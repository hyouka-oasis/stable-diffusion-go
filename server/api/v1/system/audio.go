package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type AudioApi struct{}

// CreateAudioAndSrt 创建音频文件以及字幕文件
func (s *AudioApi) CreateAudioAndSrt(c *gin.Context) {
	var params system.AudioConfig
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(params, utils.AudioRequestParamsVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = audioService.CreateAudioAndSrt(params)
	if err != nil {
		global.Log.Error("生成音频和字幕失败!", zap.Error(err))
		response.FailWithMessage("生成音频和字幕失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("生成音频和字幕成功", c)
}
