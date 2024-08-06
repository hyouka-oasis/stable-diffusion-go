package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type VideoApi struct{}

// CreateVideo 生成视频
func (s *VideoApi) CreateVideo(c *gin.Context) {
	var infoParams request.InfoCreateVideoRequest
	err := c.ShouldBindJSON(&infoParams)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(infoParams, utils.InfoCreateVideoParamsVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = videoService.CreateVideo(infoParams)
	if err != nil {
		global.Log.Error("生成失败!", zap.Error(err))
		response.FailWithMessage("生成失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("生成成功", c)
}
