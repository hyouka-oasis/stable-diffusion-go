package main

import (
	"fmt"
	"github/stable-diffusion-go/server/core"
	"github/stable-diffusion-go/server/global"
	"log"
	"os"
)

func main() {
	core.InitPiver()
	bookName := global.Config.Book.Name
	core.InitGlobalConfig()
	// 1. 读取测试.txt文件
	file, err := os.Open(global.OriginBookPath)
	fmt.Print("开始读取:" + bookName + "\n")
	if err != nil {
		log.Fatal("文件不存在:", err)
	}
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	// 2. 创建participle目录
	err = core.EnsureDirectory(global.OutPath)
	if err != nil {
		log.Fatal("创建目录失败:", err)
	}
	// 2. 创建participle目录
	err = core.EnsureDirectory(global.OutImagesPath)
	if err != nil {
		log.Fatal("创建图片目录失败:", err)
	}
	// 3. 处理文本文件
	core.ProcessText()
	// 4. 翻译文本
	core.Translate()
	// 5.翻译成功后进行字幕提取
	core.TextToSrt()
	// 6.调用
	//core.StableDiffusion()
	// 7.合成视频
	core.VideoComposition()
}
