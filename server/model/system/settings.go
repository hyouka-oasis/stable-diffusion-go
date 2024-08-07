package system

import (
	"github/stable-diffusion-go/server/global"
)

type SettingsStableDiffusionConfig struct {
	SettingsId uint   `json:"settingsId" gorm:"index;primary_key"` // 外键
	Url        string `json:"url" gorm:"default:'http://127.0.0.1:7860'"`
}

type SettingsOllamaConfig struct {
	SettingsId uint   `json:"settingsId" gorm:"primary_key"` // 外键
	ModelName  string `json:"modelName" gorm:"comment:模型名称"` // 模型名称
	Url        string `json:"url" gorm:"comment:ollama地址;default:'http://127.0.0.1:11434/v1'"`
}

type SettingsAliyunConfig struct {
	SettingsId uint   `json:"settingsId" gorm:"primary_key"` // 外键
	KeyId      string `json:"keyId" gorm:"comment:阿里云key"`   // 模型名称
	KeySecret  string `json:"keySecret" gorm:"comment:阿里云密匙"`
}

type Settings struct {
	global.Model
	TranslateType         string                        `json:"translateType" gorm:"comment:翻译设置;default:ollama"` // 项目名称
	StableDiffusionConfig SettingsStableDiffusionConfig `json:"stableDiffusionConfig" gorm:"comment:stable-diffusion配置"`
	OllamaConfig          SettingsOllamaConfig          `json:"ollamaConfig" gorm:"comment:ollama配置"`
	AliyunConfig          SettingsAliyunConfig          `json:"aliyunConfig" gorm:"comment:阿里云配置"`
	SavePath              string                        `json:"savePath"` //保存路径
}

func (Settings) TableName() string {
	return "settings"
}
