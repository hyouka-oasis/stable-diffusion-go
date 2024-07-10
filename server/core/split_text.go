package core

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"os"
	"strings"
	"sync"
)

func ProcessText() {
	// 异步处理文本
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		processLines(global.OriginBookPath, global.OutBookPath)
	}()
	wg.Wait()
}
func combineStrings(strings []string, min_words, max_words int) []string {
	var combined []string
	currentStr := ""
	for _, s := range strings {
		fmt.Println(s, "文本")
		if len(currentStr)+len(s) <= max_words && len(currentStr)+len(s) >= min_words {
			combined = append(combined, currentStr+s+"\n")
			currentStr = ""
		} else if len(currentStr) > max_words {
			combined = append(combined, currentStr+"\n")
			currentStr = s
		} else {
			currentStr += s + " "
		}
	}
	if currentStr != "" {
		combined = append(combined, currentStr+"\n")
	}
	return combined
}

func processLines(inputFilePath string, outputFilePath string) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		return
	}
	defer file.Close()
	// 读取文件的前 1024 个字节

	//for i := 0; i < len(file); i++ {
	//	char := fmt.Sprintf("%c", file[i])
	//	fmt.Println(char, string(file[i]), "内容1")
	//}
	fmt.Println(file, "内容")
	//strings := []string{"This is the first sentence.", "This is the second sentence.", "This is the third sentence.", "This is the fourth sentence."}
	//min_words := 20
	//max_words := 50
	//fmt.Println(strings)
	//combined := combineStrings(strings, min_words, max_words)
	//text := "检测到宿主已经非常完美，本系统教无可教，就此解散，拜拜~\n你特么把老子带到这里，自己就这么跑了？系统！\n\n李念凡在心中疯狂的咆哮，然而空无回应。\n\n我去你妹的！狗系统！"

	//arrayStrings, err := ReadStringsFromFile(inputFilePath)
	//if err != nil {
	//	log.Fatal("打开文件失败:", err)
	//	return
	//}
	//content := ReplaceBlankBatch(arrayStrings)
	//processedContent := processLine(content)
	//writeToFile(combined, outputFilePath)
	//err = WriteToFile(outputFilePath, processedContent)
	//if err != nil {
	//	log.Fatal("写入文件失败:", err)
	//	return
	//}
}

// removeLeadingTrailingSpaces 函数接受一个字符串切片,并返回去除前后空格的字符串切片
func removeLeadingTrailingSpaces(lines []string) []string {
	maxWords := global.Config.Potential.MaxWords
	minWords := global.Config.Potential.MinWords
	var combined []string
	currentStr := ""
	for _, s := range lines {
		fmt.Println("文本", s)
		if len(currentStr)+len(s) <= maxWords && len(currentStr)+len(s) >= minWords {
			combined = append(combined, currentStr+s+"\n")
			currentStr = ""
		} else if len(currentStr) > maxWords {
			combined = append(combined, currentStr+"\n")
			currentStr = s
		} else {
			currentStr += s + " "
		}
	}
	if currentStr != "" {
		combined = append(combined, currentStr+"\n")
	}
	return combined
}

var PUNCTUATION = []string{",", ".", "!", "？", ";", ":", "”", ",", "!", "…"}

// clause 处理非正常符号
func clause(text string) []string {
	result := []string{}
	start := 0
	for i := 0; i < len(text); i++ {
		char := fmt.Sprintf("%c", text[i])
		fmt.Println(char, text[i], "内容")
		if contains(PUNCTUATION, string(text[i])) {
			for contains(PUNCTUATION, string(text[i])) {
				i++
			}
			result = append(result, strings.TrimSpace(text[start:i]))
			start = i
		}
	}
	// Add the last clause if it exists
	if start < len(text) {
		result = append(result, strings.TrimSpace(text[start:]))
	}
	return result
}
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
func processLine(content []string) {
	//paragraphs := clause(content)
	//fmt.Println("处理后的内容")
	//resultContext := removeLeadingTrailingSpaces(paragraphs)
	//return resultContext
}
