package system

import (
	"encoding/json"
	"errors"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"io"
	"net/http"
)

type StableDiffusionService struct{}

func getSettingsConfig() (settings system.Settings, err error) {
	err = global.DB.Model(&system.Settings{}).Preload("StableDiffusionConfig").First(&settings).Error
	if err != nil {
		return settings, errors.New("请先配置")
	}
	if settings.StableDiffusionConfig.Url == "" {
		return settings, errors.New("stable-diffusion-url不能为空")
	}
	return settings, nil
}

// GetStableDiffusionSdModels 获取stable-diffusion模型列表
func (s *StableDiffusionService) GetStableDiffusionSdModels() (list []interface{}, err error) {
	settings, err := getSettingsConfig()
	if err != nil {
		return
	}
	apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/sd-models"
	client := &http.Client{
		Timeout: 0, // 设置超时时间为60秒
	}
	// 发送POST请求
	resp, err := client.Get(apiUrl)
	if err != nil {
		return list, errors.New("请检查当前stable-diffusion是否正确开启")
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	var respData []interface{}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return list, err
	}
	return respData, err
}

// GetStableDiffusionSdVae 获取stable-diffusion-vae
func (s *StableDiffusionService) GetStableDiffusionSdVae() (list []interface{}, err error) {
	settings, err := getSettingsConfig()
	if err != nil {
		return
	}
	apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/sd-vae"
	client := &http.Client{
		Timeout: 0, // 设置超时时间为60秒
	}
	// 发送POST请求
	resp, err := client.Get(apiUrl)
	if err != nil {
		return list, errors.New("请检查当前stable-diffusion是否正确开启")
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	var respData []interface{}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return list, err
	}
	return respData, err
}

// GetStableDiffusionSamplers 获取stable-diffusion采样器
func (s *StableDiffusionService) GetStableDiffusionSamplers() (list []interface{}, err error) {
	settings, err := getSettingsConfig()
	if err != nil {
		return
	}
	apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/samplers"
	client := &http.Client{
		Timeout: 0, // 设置超时时间为60秒
	}
	// 发送POST请求
	resp, err := client.Get(apiUrl)
	if err != nil {
		return list, errors.New("请检查当前stable-diffusion是否正确开启")
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	var respData []interface{}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return list, err
	}
	return respData, err
}

// GetStableDiffusionSchedulers 获取stable-diffusion调度类型
func (s *StableDiffusionService) GetStableDiffusionSchedulers() (list []interface{}, err error) {
	settings, err := getSettingsConfig()
	if err != nil {
		return
	}
	apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/schedulers"
	client := &http.Client{
		Timeout: 0, // 设置超时时间为60秒
	}
	// 发送POST请求
	resp, err := client.Get(apiUrl)
	if err != nil {
		return list, errors.New("请检查当前stable-diffusion是否正确开启")
	}
	defer resp.Body.Close()
	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	var respData []interface{}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return list, err
	}
	return respData, err
}
