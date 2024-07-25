package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
	"strconv"
)

type ProjectDetailApi struct{}

// UpdateProjectDetailFile 上传文件
func (s *ProjectDetailApi) UpdateProjectDetailFile(c *gin.Context) {
	minWords, err := strconv.Atoi(c.PostForm("minWords"))
	maxWords, err := strconv.Atoi(c.PostForm("maxWords"))
	projectDetailId, err := strconv.Atoi(c.PostForm("id"))
	file, err := c.FormFile("file")
	projectDetail := system.ProjectDetail{
		FileName: file.Filename,
		Potential: system.ProjectDetailPotential{
			MinWords: minWords,
			MaxWords: maxWords,
		},
	}
	err = projectDetailService.UpdateProjectDetailFile(projectDetailId, projectDetail)
	return
	if err != nil {
		global.Log.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	filePath := global.Config.Local.Path + "/" + file.Filename
	// 保存文件到本地
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		global.Log.Error("保存文件失败!", zap.Error(err))
		response.FailWithMessage("保存文件失败", c)
		return
	}
	//3. 处理文本文件
	//err = core.ProcessText()
	//if err != nil {
	//	panic(err)
	//}
	//err = projectService.CreateProject(projectConfig)
	//if err != nil {
	//	global.Log.Error("新增失败!", zap.Error(err))
	//	response.FailWithMessage("添加失败", c)
	//	return
	//}
	//response.OkWithMessage("添加成功", c)
}

// GetProjectDetail 获取详情
func (s *ProjectDetailApi) GetProjectDetail(c *gin.Context) {
	var config system.ProjectDetail
	err := c.ShouldBindQuery(&config)
	if err != nil {
		response.FailWithMessage("请传入参数", c)
		return
	}
	fmt.Println(&config)
	err = utils.Verify(config, utils.ProjectDetailVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	config, err = projectDetailService.GetProjectDetail(config)
	if err != nil {
		global.Log.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(&config, "获取成功", c)

}
