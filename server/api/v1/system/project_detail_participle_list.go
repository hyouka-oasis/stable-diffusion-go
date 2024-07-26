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
