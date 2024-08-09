package main

import (
	_ "embed"
	"fmt"
	"github/stable-diffusion-go/server/core"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/initialize"
	"github/stable-diffusion-go/server/python_core"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
	"os"
)

func startGinServer() {
	err := utils.EnsureDirectory(global.Config.Local.StorePath)
	if err != nil {
		panic("创建目录失败:" + err.Error())
	}
	port := core.GetRandomPort()
	fmt.Println("Using port:", port, "...")
	// 加载默认词典
	global.Seg.LoadDict()
	global.Log = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.Log)
	global.DB = initialize.Gorm() // gorm连接数据库
	if global.DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunServer(port)
}

func main() {
	// 1. 创建python所需依赖
	_, err := os.Stat(python_core.PythonRequirementsName)
	if err != nil {
		file, createError := os.Create(python_core.PythonRequirementsName)
		if createError != nil {
			global.Log.Error("创建require失败")
			panic(createError.Error())
		}
		defer file.Close()
		_, err = file.WriteString(python_core.PythonRequirements) // 写入内容
		if err != nil {
			global.Log.Error("创建require失败文件失败")
			panic(createError.Error())
		}
	}
	global.Viper = core.InitViper()
	startGinServer()
}
