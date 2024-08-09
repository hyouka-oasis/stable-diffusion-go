package example

import (
	"github.com/gin-gonic/gin"
)

type FileUploadAndDownloadRouter struct{}

func (e *FileUploadAndDownloadRouter) InitFileUploadAndDownloadRouter(Router *gin.RouterGroup) {
	fileUploadAndDownloadRouter := Router.Group("file")
	{
		fileUploadAndDownloadRouter.GET("getList", exaFileUploadAndDownloadApi.GetFileList) // 获取上传文件列表
	}
	{
		fileUploadAndDownloadRouter.POST("upload", exaFileUploadAndDownloadApi.UploadFile)         // 上传文件
		fileUploadAndDownloadRouter.POST("delete", exaFileUploadAndDownloadApi.DeleteFile)         // 删除指定文件
		fileUploadAndDownloadRouter.POST("editFileName", exaFileUploadAndDownloadApi.EditFileName) // 编辑文件名或者备注
		//fileUploadAndDownloadRouter.GET("findFile", exaFileUploadAndDownloadApi.FindFile)          // 查询当前文件成功的切片
	}
}
