package source

import (
	"fmt"
	"github/stable-diffusion-go/server/utils"
)

func KeywordExtraction(context string, pythonName string, savePath string) error {
	args := []string{
		pythonName,
		"--text", context,
		"--save_path", savePath,
	}
	err := utils.ExecCommand("python", args)
	if err != nil {
		return err
	}
	fmt.Println("关键字提取成功")
	return nil
}
