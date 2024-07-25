package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/core"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/response"
	"go.uber.org/zap"
	"strings"
)

var PUNCTUATION = []string{"，", "。", "！", "？", "；", "：", "”", ",", "!", "…"}

type ProjectDetailApi struct{}

func clause(text string) []string {
	start := 0
	punctuation := "，。！?；：，!…”“"
	var textList []string
	for i, c := range text {
		if strings.ContainsRune(punctuation, c) {
			for j := i; j < len(text) && strings.ContainsRune(punctuation, rune(text[j])); j++ {
				i++
			}
			textList = append(textList, strings.TrimSpace(text[start:i]))
			start = i + 1
		}
	}
	if start < len(text) {
		textList = append(textList, strings.TrimSpace(text[start:]))
	}
	return textList
}

// UpdateProjectDetailFile 上传文件
func (s *ProjectDetailApi) UpdateProjectDetailFile(c *gin.Context) {
	//var potential system.ProjectDetailPotential
	//err, _ := c.GetPostForm("minLength")
	//if err != nil {
	//	global.Log.Error("参数获取失败!", zap.Error(err))
	//	response.FailWithMessage("参数获取失败", c)
	//	return
	//}
	//noSave := c.DefaultQuery("noSave", "0")
	file, err := c.FormFile("file")
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
	err = core.ProcessText()
	if err != nil {
		panic(err)
	}
	//err = projectService.CreateProject(projectConfig)
	//if err != nil {
	//	global.Log.Error("新增失败!", zap.Error(err))
	//	response.FailWithMessage("添加失败", c)
	//	return
	//}
	//response.OkWithMessage("添加成功", c)
}
