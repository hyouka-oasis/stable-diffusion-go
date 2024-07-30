package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/source"
	"gorm.io/gorm"
)

type StableDiffusionService struct{}

// StableDiffusionTextToImage 批量文字转图片
func (s *StableDiffusionService) StableDiffusionTextToImage(params systemRequest.StableDiffusionParams) (images []string, err error) {
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
	return images, global.DB.Transaction(func(tx *gorm.DB) error {
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
		fmt.Println("开始")
		var projectDetailInfo system.ProjectDetailInfo
		// 查到单个的列表
		err = tx.Model(&system.ProjectDetailInfo{}).Where("id = ?", params.Id).Find(&projectDetailInfo).Error
		request["prompt"] = projectDetailInfo.Prompt
		request["negative_prompt"] = projectDetailInfo.NegativePrompt
		apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/txt2img"
		stableDiffusionImages, generateError := source.StableDiffusionGenerateImage(apiUrl, request)
		if generateError != nil {
			err = generateError
		}
		images = stableDiffusionImages
		return err
	})
}
