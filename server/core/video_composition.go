package core

import (
	"bufio"
	"fmt"
	"github/stable-diffusion-go/server/global"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func disposableSynthesisVideo() {
	// 指定 TXT 文件路径
	txtFilePath := global.OutAudioSrtMapPath

	// 打开 TXT 文件
	file, err := os.Open(txtFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// 创建 scanner 读取文件内容
	scanner := bufio.NewScanner(file)

	// 创建 map 存储结果
	var timeMap []string

	// 循环读取文件内容
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		// 去除方括号
		line = strings.TrimPrefix(line, "[")
		line = strings.TrimSuffix(line, "]")
		// 使用 strconv.Unquote 解析字符串
		value, err := strconv.Unquote(`"` + line + `"`)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		timeMap = append(timeMap, value)
		i++
	}

	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 输出 map
	fmt.Println(timeMap, "时间列表")
	//outAudioPath := global.OutAudioPath
	//outAudioSrtPath := global.OutAudioSrtPath
}

func VideoComposition() (err error) {
	imagesPath := global.OutImagesPath
	picturePathList, err := GetPicturePaths(imagesPath, ".png")
	if err != nil {
		log.Fatalln("获取文件目录失败")
		return err
	}
	// 对文件路径列表按照文件名中的数字大小进行排序
	sort.Slice(picturePathList, func(i, j int) bool {
		iNum, _ := ExtractNumber(picturePathList[i], ".png")
		jNum, _ := ExtractNumber(picturePathList[j], ".png")
		return iNum < jNum
	})
	disposableSynthesisVideo()
	log.Fatalln(picturePathList, "文件列表1")
	return
}
