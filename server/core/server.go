package core

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/initialize"
	"go.uber.org/zap"
	"net"
)

type server interface {
	ListenAndServe() error
}

func GetRandomPort() int {
	// 使用 net.Listen 获取一个随机可用端口
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port
}

func RunServer(port int) {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d", port)
	s := initServer(address, Router)
	global.Log.Info("后端运行与:", zap.String("address", address))
	fmt.Printf(`
	后端启动成功:http://127.0.0.1%s
	`, address)
	global.Log.Error(s.ListenAndServe().Error())
}
