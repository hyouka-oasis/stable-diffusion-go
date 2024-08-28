package system

import (
	"github.com/gin-gonic/gin"
	api "github/stable-diffusion-go/server/api/v1"
)

type TaskRouter struct{}

func (s *TaskRouter) InitTaskRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	taskRouter := Router.Group("task")
	taskApi := api.ApiGroupApp.SystemApiGroup.TaskApi
	{
		taskRouter.GET("list", taskApi.GetTaskList)
	}
	{
		taskRouter.POST("get", taskApi.GetTask)
	}
	return taskRouter
}
