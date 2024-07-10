package core

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// EnsureDirectory 检查并创建文件夹
func EnsureDirectory(dirPath string) error {
	// 检查文件夹是否存在
	_, err := os.Stat(dirPath)
	if err == nil {
		// 文件夹已存在，跳过创建
		return nil
	}

	// 如果文件夹不存在
	if os.IsNotExist(err) {
		// 创建文件夹
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			return err
		}
		return nil
	}

	// 其他错误
	return err
}

// ExtractNumber 获取文件名称
func ExtractNumber(filename string, suffix string) (int, error) {
	// 从文件名中提取数字
	basename := filepath.Base(filename)
	numStr := strings.TrimSuffix(basename, suffix)
	return strconv.Atoi(numStr)
}

// GetPicturePaths 获取指定目录下的特定后缀列表
func GetPicturePaths(dirPath string, suffix string) ([]string, error) {
	var picturePathList []string

	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), suffix) {
			picturePathList = append(picturePathList, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return picturePathList, nil
}

// ReplaceBlank 清除空格
func ReplaceBlank(content string) string {
	novel := strings.Replace(content, "\n", "", -1)
	novel = strings.Replace(novel, "\r", "", -1)
	novel = strings.Replace(novel, "\r\n", "", -1)
	novel = strings.Replace(novel, "\u2003", "", -1)
	return novel
}

// ReplaceBlankBatch 清除空格数组
func ReplaceBlankBatch(content []string) []string {
	var arrayString []string
	for _, line := range content {
		novel := ReplaceBlank(line)
		fmt.Println("输出的内容", line)
		arrayString = append(arrayString, novel)
	}
	return arrayString
}

// ReadStringsFromFile 将string转换为数组形式
func ReadStringsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var arrayString []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arrayString = append(arrayString, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return arrayString, nil
}

// WriteToFile 函数将字符串切片写入到指定的输出文件中
func WriteToFile(filename string, lines []string) error {
	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for _, line := range lines {
		fmt.Println(line, "文本内容")
		_, err := outputFile.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
