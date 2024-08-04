package system

import (
	"encoding/json"
	"errors"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/source"
)

type StableDiffusionService struct{}

// StableDiffusionTextToImage 批量文字转图片
func (s *StableDiffusionService) StableDiffusionTextToImage(params systemRequest.StableDiffusionRequestParams) (images []string, err error) {
	var settings system.Settings
	err = global.DB.Model(&system.Settings{}).Preload("StableDiffusionConfig").First(&settings).Error
	if err != nil {
		return images, errors.New("请先配置")
	}
	if settings.StableDiffusionConfig.Url == "" {
		return images, errors.New("stable-diffusion-url不能为空")
	}
	var projectDetail system.ProjectDetail
	err = global.DB.Model(&system.ProjectDetail{}).Where("id = ?", params.ProjectDetailId).First(&projectDetail).Error
	if err != nil {
		return images, errors.New("获取项目详情失败")
	}
	stableDiffusionParams := map[string]interface{}{}
	request := map[string]interface{}{}
	err = json.Unmarshal([]byte(projectDetail.StableDiffusionConfig), &stableDiffusionParams)
	if err == nil {
		// 如果json解析成功则合并 Stable Diffusion 配置参数
		for key, value := range stableDiffusionParams {
			request[key] = value
		}
	}
	// 异步处理翻译
	var projectDetailInfo system.Info
	// 查到单个的列表
	err = global.DB.Model(&system.Info{}).Where("id = ?", params.Id).Find(&projectDetailInfo).Error
	if projectDetailInfo.Prompt == "" {
		request["prompt"] = projectDetailInfo.Text
	} else {
		request["prompt"] = projectDetailInfo.Prompt
	}
	if projectDetailInfo.NegativePrompt != "" {
		request["negative_prompt"] = projectDetailInfo.NegativePrompt
	}
	apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/txt2img"
	stableDiffusionImages, generateError := source.StableDiffusionGenerateImage(apiUrl, request)
	if generateError != nil {
		err = generateError
	}
	images = stableDiffusionImages
	return images, err
}
