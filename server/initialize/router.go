package initialize

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/middleware"
	"github/stable-diffusion-go/server/router"
	"net/http"
	"os"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	Router.StaticFS(global.Config.Local.StorePath, justFilesFilesystem{http.Dir(global.Config.Local.StorePath)}) // Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	Router.Use(middleware.Cors())                                                                                // 直接放行全部跨域请求
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitProjectRouter(PublicGroup)                       // 项目基础接口
		systemRouter.InitProjectDetailRouter(PublicGroup)                 // 项目详情基础接口
		systemRouter.InitInfoRouter(PublicGroup)                          // 项目详情基础接口
		systemRouter.InitSettingsRouter(PublicGroup)                      // 基础设置接口
		systemRouter.InitStableDiffusionLorasRouter(PublicGroup)          // stableDiffusionLoras接口
		systemRouter.InitStableDiffusionRouter(PublicGroup)               // stableDiffusionLoras接口
		systemRouter.InitStableDiffusionNegativePromptRouter(PublicGroup) // 通用反向提示词接口
		systemRouter.InitAudioSrtRouter(PublicGroup)                      // 音频字幕生成接口
		systemRouter.InitVideoRouter(PublicGroup)                         // 视频生成接口
		systemRouter.InitStableDiffusionSettingsRouter(PublicGroup)       // stable-diffusion通用配置
		systemRouter.InitBasicRouter(PublicGroup)                         // 通用系统接口
		systemRouter.InitOllamaRouter(PublicGroup)                        // 通用系统接口
	}
	{
		exampleRouter.InitFileUploadAndDownloadRouter(PublicGroup)
	}
	// 注册业务路由
	global.Log.Info("路由注册成功")
	return Router
}
