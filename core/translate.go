package core

import (
	"ComicTweetsGo/global"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

func Translate() {
	translateType := global.Config.TranslateConfig.Type
	if translateType == "ollama" {
		// 异步处理翻译
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			translateOllama(global.BookPath, global.BookJsonPath)
		}()
		wg.Wait()
	} else if translateType == "chatgpt" {

	} else if translateType == "aliyun" {

	} else {
		print("必须要传入翻译类型")
	}
}

func translateOllama(inputFilePath string, outputFilePath string) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal("打开文件失败:", err)
	}
	defer file.Close()
	jsonContent := []map[string]string{} // 数组对象
	scanner := bufio.NewScanner(file)
	for i := 1; scanner.Scan(); i++ {
		line := scanner.Text()
		translation, err := ChatgptOllama(line)
		fmt.Print("开始通过ollama进行转换prompt，当前正在转换第", i, "段\n")
		if err != nil {
			return
		}
		jsonContent = append(jsonContent, map[string]string{
			"prompt":          translation,
			"negative_prompt": "nsfw,(low quality,normal quality,worst quality,jpeg artifacts),cropped,monochrome,lowres,low saturation,((watermark)),(white letters)",
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("读取文件失败:", err)
	}
	jsonBytes, _ := json.MarshalIndent(jsonContent, "", "  ")
	err = os.WriteFile(outputFilePath, jsonBytes, 0644)
	if err != nil {
		log.Fatal("写入文件失败:", err)
	}
}
