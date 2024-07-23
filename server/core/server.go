package core

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/initialize"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()
	//Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", global.Config.System.Addr)
	s := initServer(address, Router)
	global.Log.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	后端启动成功:http://127.0.0.1%s
	`, address)
	global.Log.Error(s.ListenAndServe().Error())
}
