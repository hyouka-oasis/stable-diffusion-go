package core

import (
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
