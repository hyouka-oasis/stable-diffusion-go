package system

import (
	"errors"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/source"
	"github/stable-diffusion-go/server/utils"
	"gorm.io/gorm"
	"strings"
)

type InfoService struct{}

// DeleteInfo 删除单条记录
func (s *InfoService) DeleteInfo(params request.ProjectDetailRequestParams) (err error) {
	if params.Id != 0 {
		err = global.DB.Transaction(func(tx *gorm.DB) error {
			err = tx.Delete(&system.Info{}, "id = ?", params.Id).Error
			if err != nil {
				return err
			}
			err = tx.Delete(&system.StableDiffusionImages{}, "info_id = ?", params.Id).Error
			if err != nil {
				return err
			}
			err = tx.Delete(&system.AudioConfig{}, "info_id = ?", params.Id).Error
			if err != nil {
				return err
			}
			return err
		})
	}
	return err
}

// UpdateInfo 更新单条记录
func (s *InfoService) UpdateInfo(updateData system.Info) error {
	err := global.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updateData).Error
	return err
}

// GetInfo 获取单条记录
func (s *InfoService) GetInfo(id uint) (info system.Info, err error) {
	err = global.DB.Model(&system.Info{}).Where("id = ?", id).First(&info).Error
	return
}

// ExtractTheInfoRole 进行人物提取
func (s *InfoService) ExtractTheInfoRole(id uint) error {
	var currentProjectDetailParticipleList []system.Info
	//var currentSettings system.Settings
	//err := global.DB.Model(&system.Settings{}).First(&currentSettings).Error
	//if err != nil {
	//	return errors.New("请先初始化配置")
	//}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&system.Info{}).Find(&currentProjectDetailParticipleList, "project_detail_id = ?", id).Error
		for _, projectDetailParticiple := range currentProjectDetailParticipleList {
			projectDetailParticiple.Role = utils.CutPos(projectDetailParticiple.Text)
			err = tx.Model(&projectDetailParticiple).Select("role").Updates(&projectDetailParticiple).Error
			return err
		}
		return err
	})
}

// TranslateInfoPrompt 进行prompt转换
func (s *InfoService) TranslateInfoPrompt(projectDetailParticipleParams system.Info) error {
	var infoList []system.Info
	var currentSettings system.Settings
	err := global.DB.Model(&system.Settings{}).Preload("OllamaConfig").First(&currentSettings).Error
	if err != nil {
		return errors.New("请先初始化配置")
	}
	if projectDetailParticipleParams.ProjectDetailId != 0 {
		err = global.DB.Model(&system.Info{}).Find(&infoList, "project_detail_id = ?", projectDetailParticipleParams.ProjectDetailId).Error
	}
	if projectDetailParticipleParams.Id != 0 {
		err = global.DB.Model(&system.Info{}).Find(&infoList, "id = ?", projectDetailParticipleParams.Id).Error
	}
	var loras []system.StableDiffusionLoras
	err = global.DB.Model(&system.StableDiffusionLoras{}).Find(&loras).Error
	if err != nil {
		return err
	}
	// 如果采用翻译模型
	if currentSettings.TranslateType == "sd-prompt-translator" {
		for _, info := range infoList {
			err = global.DB.Model(&info).Update("loading", true).Error
			if err != nil {
				return err
			}
			lorasText := ""
			for _, lora := range loras {
				if strings.Contains(lora.Roles, info.Role) {
					lorasText += lora.Name + ","
				}
			}
			info.Prompt = lorasText
			err = global.DB.Model(&info).Update("prompt", info.Prompt).Update("loading", false).Error
			if err != nil {
				return err
			}
		}
		return nil
	}
	if currentSettings.TranslateType == "ollama" {
		for _, info := range infoList {
			err = global.DB.Model(&info).Update("loading", true).Error
			lorasText := ""
			for _, lora := range loras {
				if strings.Contains(lora.Roles, info.Role) {
					lorasText += lora.Name + ","
				}
			}
			prompt, _ := source.ChatgptOllama(info.Text, currentSettings.OllamaConfig)
			if lorasText != "" {
				info.Prompt = prompt + "," + lorasText
			} else {
				info.Prompt = prompt
			}
			err = global.DB.Model(&info).Update("prompt", info.Prompt).Update("loading", false).Error
			if err != nil {
				return err
			}
		}
		return err
	}
	return err
}
