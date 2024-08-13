package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"go.uber.org/zap"
)

type OllamaApi struct{}

// GetOllamaModelList 获取ollama模型列表
func (s *OllamaApi) GetOllamaModelList(c *gin.Context) {
	list, err := ollamaService.GetOllamaModelList()
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list.Models,
		Total:    int64(len(list.Models)),
		Page:     -1,
		PageSize: 10,
	}, "获取成功", c)
}
