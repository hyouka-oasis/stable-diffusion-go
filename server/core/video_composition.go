package core

import (
	"bufio"
	"fmt"
	"github/stable-diffusion-go/server/global"
	"log"
	"os"
	"sort"
	"strings"
)

// 获取字幕切片
func getAudioSrtMap() (audioSrtMap []string, err error) {
	// 指定的srt_map路径
	txtFilePath := global.OutAudioSrtMapPath

	file, err := os.Open(txtFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// 创建 scanner 读取文件内容
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// 去除方括号
		line = strings.TrimPrefix(line, "[")
		line = strings.TrimSuffix(line, "]")
		// 使用 strconv.Unquote 解析字符串
		parts := strings.Split(line, ",")
		for _, part := range parts {
			audioSrtMap = append(audioSrtMap, strings.TrimSpace(part))
		}
		if scannerError := scanner.Err(); scannerError != nil {
			fmt.Println("解析字符串失败:", scannerError)
			return
		}
	}

	// 检查是否有错误发生
	if scannerErr := scanner.Err(); err != nil {
		fmt.Println("读取文件内容失败:", scannerErr)
		return
	}

	return audioSrtMap, nil
}

// 获取图片列表
func getImagesMap() (picturePathList []string, err error) {
	imagesPath := global.OutImagesPath
	picturePathList, err = GetPicturePaths(imagesPath, ".png")
	if err != nil {
		log.Fatalln("获取文件目录失败")
		return nil, err
	}
	// 对文件路径列表按照文件名中的数字大小进行排序
	sort.Slice(picturePathList, func(i, j int) bool {
		iNum, _ := ExtractNumber(picturePathList[i], ".png")
		jNum, _ := ExtractNumber(picturePathList[j], ".png")
		return iNum < jNum
	})
	return picturePathList, nil
}

// 整合视频
func disposableSynthesisVideo(picturePathList []string, timeSrtMap []string) {
	for _, tuple := range zip(picturePathList, timeSrtMap) {
		fmt.Println(tuple)
		//imagePath, duration := tuple[0], tuple[1]
		//fmt.Printf("下标: %d, 图片: %s, 时间1: %.2f\n", index, imagePath, duration)
	}
}

func VideoComposition() (err error) {
	picturePathList, imageError := getImagesMap()
	audioSrtMap, audioSrtMapError := getAudioSrtMap()
	if audioSrtMapError != nil || imageError != nil {
		return fmt.Errorf("读取图片列表或者字幕切片时失败")
	}
	for _, a := range audioSrtMap {
		fmt.Println(a, "便利")
	}
	disposableSynthesisVideo(picturePathList, audioSrtMap)
	return nil
}
