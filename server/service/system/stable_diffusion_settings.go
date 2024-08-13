package system

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/system"
)

type StableDiffusionSettingsService struct{}

// CreateStableDiffusionSettings 创建stabled-diffusion配置
func (s *StableDiffusionSettingsService) CreateStableDiffusionSettings(params system.StableDiffusionSettings) (err error) {
	err = global.DB.Create(&params).Error
	if err != nil {
		return err
	}
	return err
}

// UpdateStableDiffusionSettings 更新stabled-diffusion配置
func (s *StableDiffusionSettingsService) UpdateStableDiffusionSettings(params system.StableDiffusionSettings) (err error) {
	err = global.DB.Model(&system.StableDiffusionSettings{}).Where("id = ?", params.Id).Updates(&params).Error
	if err != nil {
		return err
	}
	return err
}

// DeleteStableDiffusionSettings 删除stabled-diffusion配置
func (s *StableDiffusionSettingsService) DeleteStableDiffusionSettings(ids request.IdsReq) (err error) {
	for _, id := range ids.Ids {
		err = global.DB.Delete(&system.StableDiffusionSettings{}, "id = ?", id).Error
		if err != nil {
			continue
		}
	}
	return err
}

// GetStableDiffusionSettingsList 获取stable-diffusion-settings列表
func (s *StableDiffusionSettingsService) GetStableDiffusionSettingsList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.StableDiffusionSettings{})
	err = db.Count(&total).Error
	var stableDiffusionList []*system.StableDiffusionSettings
	if err != nil {
		return stableDiffusionList, total, err
	}
	db = db.Limit(limit).Offset(offset)
	OrderStr := "id desc"
	err = db.Order(OrderStr).Find(&stableDiffusionList).Error
	if err != nil {
		return stableDiffusionList, total, err
	}
	return &stableDiffusionList, total, err
}
