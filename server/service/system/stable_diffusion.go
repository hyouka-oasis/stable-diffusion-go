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

// 统一分装请求
func getStableResponse(apiUrl string) (list []interface{}, err error) {
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

// GetStableDiffusionSdModels 获取stable-diffusion模型列表
func (s *StableDiffusionService) GetStableDiffusionSdModels() (list []interface{}, err error) {
	settings, err := getSettingsConfig()
	if err != nil {
		return
	}
	apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/sd-models"
	list, err = getStableResponse(apiUrl)
	if err != nil {
		return list, err
	}
	return list, err
}

// GetStableDiffusionSdVae 获取stable-diffusion-vae
func (s *StableDiffusionService) GetStableDiffusionSdVae() (list []interface{}, err error) {
	settings, err := getSettingsConfig()
	if err != nil {
		return
	}
	apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/sd-vae"
	list, err = getStableResponse(apiUrl)
	if err != nil {
		return list, err
	}
	return list, err
}

// GetStableDiffusionSamplers 获取stable-diffusion采样器
func (s *StableDiffusionService) GetStableDiffusionSamplers() (list []interface{}, err error) {
	settings, err := getSettingsConfig()
	if err != nil {
		return
	}
	apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/samplers"
	list, err = getStableResponse(apiUrl)
	if err != nil {
		return list, err
	}
	return list, err
}

// GetStableDiffusionSchedulers 获取stable-diffusion调度类型
func (s *StableDiffusionService) GetStableDiffusionSchedulers() (list []interface{}, err error) {
	settings, err := getSettingsConfig()
	if err != nil {
		return
	}
	apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/schedulers"
	list, err = getStableResponse(apiUrl)
	if err != nil {
		return list, err
	}
	return list, err
}

// GetStableDiffusionUpscalers 获取stable-diffusion高清算法
func (s *StableDiffusionService) GetStableDiffusionUpscalers() (list []interface{}, err error) {
	settings, err := getSettingsConfig()
	if err != nil {
		return
	}
	apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/upscalers"
	list, err = getStableResponse(apiUrl)
	if err != nil {
		return list, err
	}
	return list, err
}
