package system

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/system"
	systemResponse "github/stable-diffusion-go/server/model/system/response"
)

type StableDiffusionSettingsService struct{}

// CreateStableDiffusionSettings 创建stabled-diffusion配置
func (s *StableDiffusionSettingsService) CreateStableDiffusionSettings(params systemResponse.StableDiffusionSettingsResponse) (err error) {
	err = global.DB.Create(&params).Error
	if err != nil {
		return err
	}
	return err
}

// UpdateStableDiffusionSettings 更新stabled-diffusion配置
func (s *StableDiffusionSettingsService) UpdateStableDiffusionSettings(params systemResponse.StableDiffusionSettingsResponse) (err error) {
	err = global.DB.Model(&system.StableDiffusionSettings{}).Where("id = ?", params.Id).Updates(&params).Update("enable_hr", params.EnableHr).Update("hr_upscaler", params.HrUpscaler).Update("hr_second_pass_steps", params.HrSecondPassSteps).Update("denoising_strength", params.DenoisingStrength).Error
	if err != nil {
		return err
	}
	if params.OverrideSettings != (system.StableDiffusionOverrideSettings{}) {
		err = global.DB.Model(&system.StableDiffusionOverrideSettings{}).Where("stable_diffusion_settings_id = ?", params.Id).Updates(&params.OverrideSettings).Error
	}
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
		err = global.DB.Delete(&system.StableDiffusionOverrideSettings{}, "stable_diffusion_settings_id = ?", id).Error
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
	db := global.DB.Preload("OverrideSettings").Model(&system.StableDiffusionSettings{})
	err = db.Count(&total).Error
	var stableDiffusionList []*systemResponse.StableDiffusionSettingsResponse
	if err != nil {
		return stableDiffusionList, total, err
	}
	db = db.Limit(limit).Offset(offset)
	OrderStr := "id desc"
	err = db.Order(OrderStr).Where("project_detail_id = ?", 0).Find(&stableDiffusionList).Error
	if err != nil {
		return stableDiffusionList, total, err
	}
	return &stableDiffusionList, total, err
}

// GetStableDiffusionSettings 获取stable-diffusion-settings
func (s *StableDiffusionSettingsService) GetStableDiffusionSettings(id uint) (settings system.StableDiffusionSettings, err error) {
	err = global.DB.Model(&system.StableDiffusionSettings{}).Preload("OverrideSettings").Where("id = ?", id).First(&settings).Error
	if err != nil {
		return settings, err
	}
	return
}
