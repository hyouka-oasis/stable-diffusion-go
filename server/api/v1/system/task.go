package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
)

type TaskApi struct{}

// GetTaskList 获取task列表
func (s *TaskApi) GetTaskList(c *gin.Context) {
	var pageInfo systemRequest.TaskPageInfoRequest
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := taskService.GetTaskList(pageInfo)
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

// GetTask 获取task
func (s *TaskApi) GetTask(c *gin.Context) {
	var params systemRequest.TaskPageInfoRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	if params.TaskId == 0 && params.ProjectDetailId == 0 {
		response.FailWithMessage("任务Id和详情Id必须传一个", c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	task, err := taskService.GetTask(params)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(&task, "获取成功", c)
}
