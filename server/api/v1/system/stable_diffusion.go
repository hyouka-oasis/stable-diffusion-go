package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strings"
)

type StableDiffusionApi struct{}

// StableDiffusionTextToImageBatch 批量文件转图片
func (s *StableDiffusionApi) StableDiffusionTextToImageBatch(c *gin.Context) {
	var params systemRequest.StableDiffusionParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	err = utils.Verify(params, utils.StableDiffusionParamsVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(params.Ids) == 0 {
		response.FailWithMessage("ids不能为空", c)
	}
	err = stableDiffusionService.StableDiffusionTextToImageBatchTest(params)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("生成成功", c)
}

func (s *StableDiffusionApi) StableDiffusionTextToImage(c *gin.Context) {
	var params systemRequest.StableDiffusionParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	resp, err := http.Post("http://127.0.0.1:7860/sdapi/v1/txt2img", "application/json", strings.NewReader(params.StableDiffusionConfig))
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	c.JSON(http.StatusOK, gin.H{"response": string(body)})
}
