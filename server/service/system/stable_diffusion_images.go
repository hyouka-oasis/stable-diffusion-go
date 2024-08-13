package system

import (
	"encoding/json"
	"errors"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/example"
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
	stableDiffusionRequest := map[string]interface{}{}
	err = json.Unmarshal([]byte(projectDetail.StableDiffusionConfig), &stableDiffusionParams)
	if err == nil {
		// 如果json解析成功则合并 Stable Diffusion 配置参数
		for key, value := range stableDiffusionParams {
			stableDiffusionRequest[key] = value
		}
	}
	for _, infoId := range params.Ids {
		if err != nil {
			continue
		}
		// 异步处理翻译
		var info system.Info
		// 查到单个的列表
		err = global.DB.Model(&system.Info{}).Where("id = ?", infoId).Find(&info).Error
		if err != nil {
			continue
		}
		if info.Prompt == "" {
			stableDiffusionRequest["prompt"] = info.Text
		} else {
			stableDiffusionRequest["prompt"] = info.Prompt
		}
		if info.NegativePrompt != "" {
			stableDiffusionRequest["negative_prompt"] = info.NegativePrompt
		}
		apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/txt2img"
		stableDiffusionImages, generateError := source.StableDiffusionGenerateImage(apiUrl, stableDiffusionRequest)
		if generateError != nil {
			err = generateError
			continue
		}
		images = stableDiffusionImages
	}
	return images, err
}

// DeleteStableDiffusionImage 批量删除图片
func (s *StableDiffusionService) DeleteStableDiffusionImage(params request.IdsReq) error {
	err := global.DB.Model(&system.StableDiffusionImages{}).Error
	if err != nil {
		return err
	}
	for _, id := range params.Ids {
		var stableDiffusionImages system.StableDiffusionImages
		err = global.DB.Where("id = ?", id).First(&stableDiffusionImages).Error
		if err != nil {
			return err
		}
		var info system.Info
		err = global.DB.Where("id = ?", stableDiffusionImages.InfoId).First(&info).Error
		if err != nil {
			return err
		}
		if info.StableDiffusionImageId == id {
			err = global.DB.Model(&info).Update("stable_diffusion_image_id", 0).Error
		}
		err = global.DB.Where("id = ?", id).Delete(&system.StableDiffusionImages{}).Error
		if err != nil {
			return err
		}
		err = global.DB.Where("id = ?", stableDiffusionImages.FileId).Delete(&example.ExaFileUploadAndDownload{}).Error
		if err != nil {
			return err
		}
	}
	return err
}

// AddStableDiffusionImage 添加图片
func (s *StableDiffusionService) AddStableDiffusionImage(params system.StableDiffusionImages) error {
	err := global.DB.Model(&system.StableDiffusionImages{}).Create(&params).Error
	return err
}
