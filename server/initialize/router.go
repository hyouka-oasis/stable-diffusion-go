package initialize

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/middleware"
	"github/stable-diffusion-go/server/router"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	if gin.Mode() == gin.DebugMode {
		Router.Use(gin.Logger())
	}
	systemRouter := router.RouterGroupApp.System
	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitStableDiffusionRouter(PublicGroup) // stableDiffusion配置接口
	}
	// 注册业务路由
	global.Log.Info("路由注册成功")
	return Router
}
