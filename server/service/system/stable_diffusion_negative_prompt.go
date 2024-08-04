package system

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/system"
)

type StableDiffusionNegativePromptService struct{}

// CreateStableDiffusionNegativePrompt 创建同用反向提示词
func (s *StableDiffusionNegativePromptService) CreateStableDiffusionNegativePrompt(stableDiffusionNegativePrompt system.StableDiffusionNegativePrompt) error {
	return global.DB.Create(&stableDiffusionNegativePrompt).Error
}

// GetStableDiffusionNegativePrompt 获取列表
func (s *StableDiffusionNegativePromptService) GetStableDiffusionNegativePrompt(params request.PageInfo) (list interface{}, total int64, err error) {
	limit := params.PageSize
	offset := params.PageSize * (params.Page - 1)
	db := global.DB.Model(&system.StableDiffusionNegativePrompt{})
	var stableDiffusionNegativePromptList []system.StableDiffusionNegativePrompt
	err = db.Count(&total).Error
	if err != nil {
		return stableDiffusionNegativePromptList, total, err
	}
	db = db.Limit(limit).Offset(offset)
	OrderStr := "id desc"
	err = db.Order(OrderStr).Find(&stableDiffusionNegativePromptList).Error
	return stableDiffusionNegativePromptList, total, err
}

// UpdateStableDiffusionNegativePrompt 更新同用反向提示词
func (s *StableDiffusionNegativePromptService) UpdateStableDiffusionNegativePrompt(stableDiffusionNegativePrompt system.StableDiffusionNegativePrompt) error {
	return global.DB.Model(&system.StableDiffusionNegativePrompt{}).Where("id = ?", stableDiffusionNegativePrompt.Id).Updates(&stableDiffusionNegativePrompt).Error
}

// DeleteStableDiffusionNegativePrompt 删除同用反向提示词
func (s *StableDiffusionNegativePromptService) DeleteStableDiffusionNegativePrompt(id uint) error {
	return global.DB.Where("id = ?", id).Delete(&system.StableDiffusionNegativePrompt{}).Error
}
