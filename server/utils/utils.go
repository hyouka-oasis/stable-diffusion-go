package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/source"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"os/exec"
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
		err = os.MkdirAll(dirPath, os.ModePerm)
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

func Zip(a []string, b []string) [][]string {
	maxLen := maxLength(len(a), len(b))
	result := make([][]string, maxLen)
	for i := range result {
		result[i] = make([]string, 2)
		if i < len(a) {
			result[i][0] = a[i]
		} else {
			result[i][0] = ""
		}
		if i < len(b) {
			result[i][1] = b[i]
		} else {
			result[i][1] = "0.0"
		}
	}
	return result
}

func maxLength(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ExecCommand 异步执行cmd
func ExecCommand(name string, args []string) error {
	cmd := exec.Command(name, args...)
	err := cmd.Start()
	if err != nil {
		return errors.New("执行" + name + "start失败:" + err.Error())
	}
	err = cmd.Wait()
	if err != nil {
		return errors.New("执行" + name + "wait失败:" + err.Error())
	}
	body, err := cmd.CombinedOutput()
	if string(body) != "" {
		return errors.New("执行" + name + "combinedOutput失败:" + err.Error())
	}
	return nil
}

func extractStr(content string) (string, string) {
	parts := strings.Split(content, "**Negative Prompt:**")
	if len(parts) < 2 {
		return "", ""
	}

	prompt := strings.TrimSpace(strings.Replace(parts[0], "**Prompt:**", "", 1))
	prompt = strings.Replace(prompt, "Prompt:", "", 1)
	prompt = strings.Replace(prompt, "\n", "", -1)

	negativePrompt := strings.TrimSpace(strings.Replace(parts[1], "**Negative Prompt:**", "", 1))
	negativePrompt = strings.Replace(negativePrompt, "Negative", "", 1)
	negativePrompt = strings.Replace(negativePrompt, "Prompt:", "", 1)
	negativePrompt = strings.Replace(negativePrompt, "**Prompt:**", "", 1)
	negativePrompt = strings.Replace(negativePrompt, "\n", "", -1)

	return prompt, negativePrompt
}

func GetInterfaceToInt(t1 interface{}) int {
	var t2 int
	switch t1.(type) {
	case uint:
		t2 = int(t1.(uint))
		break
	case int8:
		t2 = int(t1.(int8))
		break
	case uint8:
		t2 = int(t1.(uint8))
		break
	case int16:
		t2 = int(t1.(int16))
		break
	case uint16:
		t2 = int(t1.(uint16))
		break
	case int32:
		t2 = int(t1.(int32))
		break
	case uint32:
		t2 = int(t1.(uint32))
		break
	case int64:
		t2 = int(t1.(int64))
		break
	case uint64:
		t2 = int(t1.(uint64))
		break
	case float32:
		t2 = int(t1.(float32))
		break
	case float64:
		t2 = int(t1.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(t1.(string))
		if t2 == 0 && len(t1.(string)) > 0 {
			f, _ := strconv.ParseFloat(t1.(string), 64)
			t2 = int(f)
		}
		break
	case nil:
		t2 = 0
		break
	case json.Number:
		t3, _ := t1.(json.Number).Int64()
		t2 = int(t3)
		break
	default:
		t2 = t1.(int)
		break
	}
	return t2
}

func SplitTextUploadFileToLocal(file *multipart.FileHeader, filePath string, config system.ProjectDetail) error {
	f, openError := file.Open() // 读取文件
	if openError != nil {
		return errors.New("打开文件时失败:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	out, createErr := os.Create(filePath)
	if createErr != nil {
		return errors.New("创建文件时失败:" + createErr.Error())
	}
	defer out.Close()             // 创建文件 defer 关闭
	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if createErr != nil {
		return errors.New("拷贝文件时失败:" + copyErr.Error())
	}
	splitTextError := source.SplitText(config)
	if splitTextError != nil {
		return errors.New("进行分词失败:" + splitTextError.Error())
	}
	return nil
}
