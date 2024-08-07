package system

import (
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/model/common/response"
	"os"
)

type BasicApi struct{}

// ExitGin 终止gin
func (s *BasicApi) ExitGin(c *gin.Context) {
	c.Abort()
	response.OkWithMessage("终止", c)
	os.Exit(1) // 使用非零退出码终止程序
}
