package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
	"strconv"
)

type ProjectDetailApi struct{}

// UploadProjectDetailFile 上传文件
func (s *ProjectDetailApi) UploadProjectDetailFile(c *gin.Context) {
	projectDetailId, err := strconv.Atoi(c.PostForm("id"))
	saveType := c.DefaultPostForm("saveType", "create")
	whetherParticiple := c.DefaultPostForm("whetherParticiple", "yes")
	file, err := c.FormFile("file")
	err = projectDetailService.UploadProjectDetailFile(uint(projectDetailId), file, saveType, whetherParticiple)
	if err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// UpdateProjectDetail 更新详情
func (s *ProjectDetailApi) UpdateProjectDetail(c *gin.Context) {
	var config request.UpdateProjectDetailRequestParams
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(config, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectDetailService.UpdateProjectDetail(config)
	if err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithDetailed(&config, "更新成功", c)
}

// GetProjectDetail 获取详情
func (s *ProjectDetailApi) GetProjectDetail(c *gin.Context) {
	var config system.ProjectDetail
	err := c.ShouldBindQuery(&config)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(config, utils.ProjectDetailVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	detail, err := projectDetailService.GetProjectDetail(config)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(detail, "获取成功", c)
}
