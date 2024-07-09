package core

import (
	"os"
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
