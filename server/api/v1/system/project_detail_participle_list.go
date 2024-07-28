package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type ProjectDetailParticipleListApi struct{}

// DeleteProjectDetailParticipleListItem 删除单条记录
func (s *ProjectDetailParticipleListApi) DeleteProjectDetailParticipleListItem(c *gin.Context) {
	var formList system.ProjectDetailParticipleList
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
	err = projectDetailParticipleListService.DeleteProjectDetailParticipleListItem(formList.Id)
	if err != nil {
		global.Log.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// ExtractTheCharacterProjectDetailParticipleList 提取角色
func (s *ProjectDetailParticipleListApi) ExtractTheCharacterProjectDetailParticipleList(c *gin.Context) {
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
	err = projectDetailParticipleListService.ExtractTheCharacterProjectDetailParticipleList(projectDetail.Id)
	if err != nil {
		global.Log.Error("角色提取失败!", zap.Error(err))
		response.FailWithMessage("角色提取失败", c)
		return
	}
	response.OkWithMessage("角色提取成功", c)
}

// TranslateProjectDetailParticipleList 进行
func (s *ProjectDetailParticipleListApi) TranslateProjectDetailParticipleList(c *gin.Context) {
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
	err = projectDetailParticipleListService.TranslateProjectDetailParticipleList(projectDetail.Id, c)
	if err != nil {
		global.Log.Error("进行prompt转换失败!", zap.Error(err))
		response.FailWithMessage("进行prompt转换失败", c)
		return
	}
	response.OkWithMessage("进行prompt转换成功", c)
}
