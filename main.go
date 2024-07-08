package main

import (
	"ComicTweetsGo/core"
	"ComicTweetsGo/global"
	"fmt"
	"log"
	"os"
)

// ... (OpenAI API key and endpoint)

func main() {
	core.InitPiver()
	bookName := global.Config.Book.Name
	global.OutPath = bookName + "/participle/"
	global.ImagesPath = bookName + "/images/"
	global.OriginBookPath = bookName + ".txt"
	global.BookPath = global.OutPath + global.OriginBookPath
	global.BookJsonPath = global.OutPath + bookName + ".json"
	global.BookMp3Path = global.OutPath + bookName + ".mp3"
	// 1. 读取测试.txt文件
	file, err := os.Open(global.OriginBookPath)
	fmt.Print("开始读取:" + bookName + "\n")
	if err != nil {
		log.Fatal("文件不存在:", err)
	}
	defer file.Close()
	// 2. 创建participle目录
	err = core.EnsureDirectory(global.OutPath)
	if err != nil {
		log.Fatal("创建目录失败:", err)
	}
	// 2. 创建participle目录
	err = core.EnsureDirectory(global.ImagesPath)
	if err != nil {
		log.Fatal("创建图片目录失败:", err)
	}
	// 3. 处理文本文件
	core.ProcessText()
	// 4. 翻译文本
	core.Translate()
	// 5.翻译成功后进行字幕提取
	core.TextToSrt()
	// 5.调用
	//core.StableDiffusion()
}
