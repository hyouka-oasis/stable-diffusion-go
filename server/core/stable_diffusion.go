package core

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github/stable-diffusion-go/server/global"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
)

type Response struct {
	Images []string `json:"images"`
}

func StableDiffusion() (err error) {
	// 读取 JSON 文件内容
	jsonContent, err := os.ReadFile(global.OutBookJsonPath)
	if err != nil {
		log.Fatal("读取 JSON 文件失败:", err)
		return err
	}

	// 解析 JSON 数据
	var jsonData []map[string]string
	err = json.Unmarshal(jsonContent, &jsonData)
	if err != nil {
		log.Fatal("解析 JSON 数据失败:", err)
		return err
	}

	// 读取 Stable Diffusion 配置文件
	stableDiffusionConfig, err := os.ReadFile("stable_diffusion.json")
	if err != nil {
		log.Fatal("读取 Stable Diffusion 配置文件失败:", err)
		return err
	}
	// 解析 Stable Diffusion 配置
	var stableDiffusionParams map[string]interface{}
	err = json.Unmarshal(stableDiffusionConfig, &stableDiffusionParams)
	if err != nil {
		log.Fatal("解析 Stable Diffusion 配置失败:", err)
		return err
	}
	height := global.Config.StableDiffusionConfig.Height
	width := global.Config.StableDiffusionConfig.Width
	//取出map中的所有key存入切片keys
	var keys = make([]int, 0, len(jsonData))
	for key := range jsonData {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Ints(keys)
	for _, index := range keys {
		data := jsonData[index]
		// 构造 Stable Diffusion 请求参数
		request := map[string]interface{}{
			"prompt":          data["prompt"],
			"negative_prompt": data["negative_prompt"],
			"height":          height,
			"width":           width,
		}
		// 合并 Stable Diffusion 配置参数
		for key, value := range stableDiffusionParams {
			request[key] = value
		}
		// 发送请求并生成图片
		err = generateImage(request, index+1)
		if err != nil {
			log.Fatal("生成图片失败:", err)
			return
		}
	}
	return nil
}

// 生成图片函数
func generateImage(request map[string]interface{}, index int) error {
	serverUrl := global.Config.StableDiffusionConfig.Url
	apiUrl := serverUrl + "/sdapi/v1/txt2img"
	fmt.Print("开始生成第", index, "张图片\n")
	// 发送请求并获取图片数据
	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("转换请求参数失败")
	}

	// 发送POST请求
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("请检查当前stable-diffusion是否正确开启")
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("响应失败")
	}
	var respData Response
	err = json.Unmarshal(body, &respData)
	if err != nil {
		log.Fatalf("解析响应数据失败: %v", err)
	}
	if len(respData.Images) > 0 {
		image := respData.Images[0]
		// 将 base64 编码的字符串解码为 []byte
		imageData, err := base64.StdEncoding.DecodeString(image)
		// 保存图片
		imagePath := filepath.Join(global.OutImagesPath, fmt.Sprintf("%d.png", index))
		err = os.WriteFile(imagePath, imageData, 0644)
		if err != nil {
			log.Fatalf("保存图片失败: %v", err)
		}
		fmt.Print("第", index, "张图片生成完成\n")
	} else {
		fmt.Print("第", index, "张图片生成失败\n", respData, err)
	}

	return nil
}
