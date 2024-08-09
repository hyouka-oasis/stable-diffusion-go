package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type InfoApi struct{}

// DeleteInfo 删除单条记录
func (s *InfoApi) DeleteInfo(c *gin.Context) {
	var params request.ProjectDetailRequestParams
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
	err = infoService.DeleteInfo(params)
	if err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateInfo 更新单条记录
func (s *InfoApi) UpdateInfo(c *gin.Context) {
	var info system.Info
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(info, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = infoService.UpdateInfo(info)
	if err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// UpdateInfoAudioConfig 更新详情的音频
func (s *InfoApi) UpdateInfoAudioConfig(c *gin.Context) {
	var info system.AudioConfig
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(info, utils.InfoCreateVideoParamsVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = infoService.UpdateInfoAudioConfig(info)
	if err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetInfo 获取单条记录
func (s *InfoApi) GetInfo(c *gin.Context) {
	var info system.Info
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(info, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := infoService.GetInfo(info.Id)
	if err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithDetailed(&res, "获取成功", c)
}

// ExtractTheInfoRole 提取角色
func (s *InfoApi) ExtractTheInfoRole(c *gin.Context) {
	var projectDetail system.ProjectDetail
	err := c.ShouldBindJSON(&projectDetail)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(projectDetail, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = infoService.ExtractTheInfoRole(projectDetail.Id)
	if err != nil {
		global.Log.Error("角色提取失败!", zap.Error(err))
		response.FailWithMessage("角色提取失败", c)
		return
	}
	response.OkWithMessage("角色提取成功", c)
}

// TranslateInfoPrompt 进行翻译
func (s *InfoApi) TranslateInfoPrompt(c *gin.Context) {
	var infoParams system.Info
	err := c.ShouldBindJSON(&infoParams)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	if infoParams.ProjectDetailId == 0 && infoParams.Id == 0 {
		response.FailWithMessage("projectDetailId和Id必须传一个", c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = infoService.TranslateInfoPrompt(system.Info{
		ProjectDetailId: infoParams.ProjectDetailId,
		Model: global.Model{
			Id: infoParams.Id,
		},
	})
	if err != nil {
		global.Log.Error("翻译失败!", zap.Error(err))
		response.FailWithMessage("翻译失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("翻译成功", c)
}
