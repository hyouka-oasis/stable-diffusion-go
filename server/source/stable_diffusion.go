package source

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github/stable-diffusion-go/server/model/system"
	systemResponse "github/stable-diffusion-go/server/model/system/response"
	"io"
	"net/http"
)

// StableDiffusionGenerateImage 生成图片函数
func StableDiffusionGenerateImage(url string, request system.StableDiffusionSettings) (images []string, err error) {
	// 发送请求并获取图片数据
	jsonData, err := json.Marshal(request)
	fmt.Println(string(jsonData), "给的内容")
	if err != nil {
		return images, errors.New("转换请求参数失败")
	}
	client := &http.Client{
		Timeout: 0, // 设置超时时间为60秒
	}
	// 发送POST请求
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return images, errors.New("请检查当前stable-diffusion是否正确开启")
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return images, errors.New("响应失败")
	}
	var respData systemResponse.StableDiffusionResponse
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return images, errors.New("解析响应数据失败")
	}
	return respData.Images, nil
}
