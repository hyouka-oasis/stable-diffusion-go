package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type AudioSrtApi struct{}

// CreateAudioAndSrt 创建音频文件以及字幕文件
func (s *AudioSrtApi) CreateAudioAndSrt(c *gin.Context) {
	var params systemRequest.AudioSrtRequestParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(params, utils.AudioSrtRequestParamsVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = audioSrtService.CreateAudioAndSrt(params)
	if err != nil {
		global.Log.Error("生成音频和字幕失败!", zap.Error(err))
		response.FailWithMessage("生成音频和字幕失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("生成音频和字幕成功", c)
}
