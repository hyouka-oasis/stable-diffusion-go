package system

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
)

type SettingsService struct{}

// CreateSettings 新增项目
func (s *SettingsService) CreateSettings(config system.Settings) (err error) {
	var settings []system.Settings
	err = global.DB.Find(&settings).Error
	if err != nil {
		return err
	}
	if len(settings) > 0 {
		return fmt.Errorf("已经存在配置文件")
	}
	err = global.DB.Create(&config).Error
	return err
}

// UpdateSettings 修改配置
func (s *SettingsService) UpdateSettings(config system.Settings) (err error) {
	err = global.DB.Model(&config).Updates(&config).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&config.StableDiffusionConfig).Updates(&config.StableDiffusionConfig).Error
	if err != nil {
		return err
	}
	err = global.DB.Model(&config.OllamaConfig).Updates(&config.OllamaConfig).Error
	if err != nil {
		return err
	}
	return nil
}

// GetSettings 获取项目列表
func (s *SettingsService) GetSettings() (settings system.Settings, err error) {
	var config system.Settings
	err = global.DB.Preload("StableDiffusionConfig").Preload("OllamaConfig").First(&config).Error
	fmt.Println(config)
	return config, err
}
