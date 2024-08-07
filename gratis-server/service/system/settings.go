package system

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"path/filepath"
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
	config.SavePath = filepath.ToSlash(config.SavePath)
	err = global.DB.Create(&config).Error
	return err
}

// UpdateSettings 修改配置
func (s *SettingsService) UpdateSettings(config system.Settings) (err error) {
	config.SavePath = filepath.ToSlash(config.SavePath)
	err = global.DB.Model(&config).Updates(&config).Error
	if err != nil {
		return err
	}
	return err
}

// GetSettings 获取项目列表
func (s *SettingsService) GetSettings() (settings system.Settings, err error) {
	var config system.Settings
	err = global.DB.Preload("StableDiffusionConfig").Preload("OllamaConfig").Preload("AliyunConfig").First(&config).Error
	return config, err
}
