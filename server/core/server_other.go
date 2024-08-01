//go:build !windows
// +build !windows

package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 0
	s.WriteTimeout = 0
	s.MaxHeaderBytes = 1 << 20
	return s
}
