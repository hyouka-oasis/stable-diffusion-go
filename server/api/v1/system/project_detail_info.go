package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type ProjectDetailInfoApi struct{}

// DeleteProjectDetailInfo 删除单条记录
func (s *ProjectDetailInfoApi) DeleteProjectDetailInfo(c *gin.Context) {
	var formList system.ProjectDetailInfo
	err := c.ShouldBindJSON(&formList)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(formList, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectDetailParticipleListService.DeleteProjectDetailInfo(formList.Id)
	if err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateProjectDetailInfo 更新单条记录
func (s *ProjectDetailInfoApi) UpdateProjectDetailInfo(c *gin.Context) {
	var info system.ProjectDetailInfo
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
	err = projectDetailParticipleListService.UpdateProjectDetailInfo(info)
	if err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetProjectDetailInfo 获取单条记录
func (s *ProjectDetailInfoApi) GetProjectDetailInfo(c *gin.Context) {
	var info system.ProjectDetailInfo
	err := c.ShouldBindQuery(&info)
	fmt.Println(info)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(info, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := projectDetailParticipleListService.GetProjectDetailInfo(info.Id)
	if err != nil {
		global.Log.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithDetailed(&res, "获取成功", c)
}

// ExtractTheRoleProjectDetailInfoList 提取角色
func (s *ProjectDetailInfoApi) ExtractTheRoleProjectDetailInfoList(c *gin.Context) {
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
	err = projectDetailParticipleListService.ExtractTheRoleProjectDetailInfoList(projectDetail.Id)
	if err != nil {
		global.Log.Error("角色提取失败!", zap.Error(err))
		response.FailWithMessage("角色提取失败", c)
		return
	}
	response.OkWithMessage("角色提取成功", c)
}

// TranslateProjectDetailInfoList 进行翻译
func (s *ProjectDetailInfoApi) TranslateProjectDetailInfoList(c *gin.Context) {
	var projectDetailParticipleParams system.ProjectDetailInfo
	err := c.ShouldBindJSON(&projectDetailParticipleParams)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	if projectDetailParticipleParams.ProjectDetailId == 0 && projectDetailParticipleParams.Id == 0 {
		response.FailWithMessage("projectDetailId和Id必须传一个", c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectDetailParticipleListService.TranslateProjectDetailInfoList(system.ProjectDetailInfo{
		ProjectDetailId: projectDetailParticipleParams.ProjectDetailId,
		Model: global.Model{
			Id: projectDetailParticipleParams.Id,
		},
	})
	if err != nil {
		global.Log.Error("翻译失败!", zap.Error(err))
		response.FailWithMessage("翻译失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("翻译成功", c)
}
