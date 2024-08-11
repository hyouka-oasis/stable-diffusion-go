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

type ProjectApi struct{}

// CreateProject 创建项目
func (s *ProjectApi) CreateProject(c *gin.Context) {
	var projectConfig system.Project
	err := c.ShouldBindJSON(&projectConfig)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(projectConfig, utils.ProjectVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectService.CreateProject(projectConfig)
	if err != nil {
		global.Log.Error("新增失败!", zap.Error(err))
		response.FailWithMessage("添加失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// DeleteProject 删除项目
func (s *ProjectApi) DeleteProject(c *gin.Context) {
	var project system.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(project, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectService.DeleteProject(project.Id)
	if err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetProjectList 获取项目列表
func (s *ProjectApi) GetProjectList(c *gin.Context) {
	var pageInfo systemRequest.ProjectRequestParams
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
	list, total, err := projectService.GetProjectList(pageInfo.Project, pageInfo.PageInfo)
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

// UpdateProject 更新
func (s *ProjectApi) UpdateProject(c *gin.Context) {
	var project system.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(project, utils.ProjectVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectService.UpdateProject(project)
	if err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
