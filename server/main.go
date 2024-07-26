package main

import (
	"fmt"
	"github/stable-diffusion-go/server/core"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/initialize"
	"github/stable-diffusion-go/server/utils"
	"go.uber.org/zap"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
)

func batchGoRun(bookName string) {
	// 1. 读取测试.txt文件
	file, err := os.Open(global.BookPath)
	fmt.Print("开始读取:" + bookName + "\n")
	if err != nil {
		log.Fatal("文件不存在:", err)
	}
	defer file.Close()
	// 2. 创建participle目录
	err = utils.EnsureDirectory(global.OutParticiplePath)
	if err != nil {
		log.Fatal("创建目录失败:", err)
	}
	// 2. 创建images目录
	err = utils.EnsureDirectory(global.OutImagesPath)
	if err != nil {
		log.Fatal("创建图片目录失败:", err)
	}
	// 2. 创建video目录
	err = utils.EnsureDirectory(global.OutVideoPath)
	if err != nil {
		log.Fatal("创建视频目录失败:", err)
	}
	//3. 处理文本文件
	//err = core.SplitText()
	//if err != nil {
	//	panic(err)
	//}
	// 4. 翻译文本
	//err = core.Translate()
	//if err != nil {
	//	panic(err)
	//}
	//5.翻译成功后进行字幕提取
	//core.TextToSrt()
	// 6.调用
	//core.StableDiffusion()
	// 7.合成视频
	//core.VideoComposition()
}

func main() {
	global.Viper = core.InitViper()
	core.InitGlobalConfig("")
	// 1. 创建文件目录
	err := utils.EnsureDirectory(global.Config.Local.Path)
	if err != nil {
		log.Fatal("创建目录失败:", err)
	}
	global.Log = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.Log)
	global.DB = initialize.Gorm() // gorm连接数据库
	if global.DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.DB.DB()
		defer db.Close()
	}
	core.RunServer()
	bookName := global.Config.Book.Name
	if bookName == "123" {
		fmt.Println("测试阶段")
		return
	}
	if global.Config.Book.Batch {
		// 使用正则表达式提取名称和数字
		re := regexp.MustCompile(`(.*?)(\d+)-(\d+)`)
		matches := re.FindStringSubmatch(bookName)
		if len(matches) >= 4 {
			var wg sync.WaitGroup
			namePrefix := matches[1]
			start, _ := strconv.Atoi(matches[2])
			end, _ := strconv.Atoi(matches[3])
			for i := start; i <= end; i++ {
				name := namePrefix + strconv.Itoa(i)
				core.InitGlobalConfig(name)
				wg.Add(1)
				go func() {
					batchGoRun(name)
					defer wg.Done()
				}()
				wg.Wait()
			}
		}
	} else {
		core.InitGlobalConfig(bookName)
		batchGoRun(bookName)
	}
}
