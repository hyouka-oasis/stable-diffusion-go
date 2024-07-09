package core

import (
	"bufio"
	"fmt"
	"github/stable-diffusion-go/server/global"
	"log"
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

func processLines(inputFilePath string, outputFilePath string) {
	file, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatal("打开文件失败:", err)
	}
	defer file.Close()

	content := readFileContent(file)
	processedContent := processLine(content)
	err = writeToFile(outputFilePath, processedContent)
	if err != nil {
		log.Fatal("写入文件失败:", err)
		return
	}
}

func readFileContent(file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	if err := scanner.Err(); err != nil {
		log.Fatal("读取文件失败:", err)
	}
	return scanner
}

// removeLeadingTrailingSpaces 函数接受一个字符串切片,并返回去除前后空格的字符串切片
func removeLeadingTrailingSpaces(lines []string) []string {
	var result []string
	for index, line := range lines {
		fmt.Print("开始处理", index+1, "段\n")
		if strings.TrimSpace(line) != "" {
			result = append(result, strings.TrimSpace(line))
		}
	}
	return result
}

func processLine(scanner *bufio.Scanner) []string {
	// 初始化变量
	var currentParagraph []string
	var paragraphs []string
	const maxLineLength = 30
	isSplit := global.Config.Potential.Split
	if !isSplit {
		for scanner.Scan() {
			paragraphs = append(paragraphs, scanner.Text())
		}
		return removeLeadingTrailingSpaces(paragraphs)
	}
	// 逐行扫描文件
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 如果当前行为空行,说明一个段落结束
		if line == "" {
			// 处理当前段落
			if len(currentParagraph) > 0 {
				paragraph := strings.Join(currentParagraph, " ")
				if len(paragraph) >= maxLineLength {
					// 如果段落长度超过 maxLineLength,则需要换行
					for len(paragraph) >= maxLineLength {
						paragraphs = append(paragraphs, paragraph[:maxLineLength])
						paragraph = paragraph[maxLineLength:]
					}
					if len(paragraph) > 0 {
						paragraphs = append(paragraphs, paragraph)
					}
				} else {
					paragraphs = append(paragraphs, paragraph)
				}
				currentParagraph = nil
			}
		} else {
			// 将当前行添加到当前段落
			currentParagraph = append(currentParagraph, line)
		}
	}

	// 处理最后一个段落
	if len(currentParagraph) > 0 {
		paragraph := strings.Join(currentParagraph, " ")
		if len(paragraph) >= maxLineLength {
			// 如果段落长度超过 maxLineLength,则需要换行
			for len(paragraph) >= maxLineLength {
				paragraphs = append(paragraphs, paragraph[:maxLineLength])
				paragraph = paragraph[maxLineLength:]
			}
			if len(paragraph) > 0 {
				paragraphs = append(paragraphs, paragraph)
			}
		} else {
			paragraphs = append(paragraphs, paragraph)
		}
	}
	return paragraphs
}

// writeToFile 函数将字符串切片写入到指定的输出文件中
func writeToFile(filename string, lines []string) error {
	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
