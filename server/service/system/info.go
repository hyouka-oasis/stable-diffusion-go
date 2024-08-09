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

// UpdateInfoAudioConfig 更新项目的所有audio
func (s *InfoService) UpdateInfoAudioConfig(config system.AudioConfig) error {
	var projectDetail system.ProjectDetail
	err := global.DB.Preload("AudioConfig").Model(&system.ProjectDetail{}).Where("id = ?", config.ProjectDetailId).First(&projectDetail).Error
	if err != nil {
		return err
	}
	var infoList []system.Info
	err = global.DB.Model(&system.Info{}).Where("project_detail_id = ?", config.ProjectDetailId).Find(&infoList).Error
	if err != nil {
		return err
	}
	for _, info := range infoList {
		err = global.DB.Model(&system.AudioConfig{}).Where("info_id = ?", info.Id).Omit("info_id", "project_detail_id").Updates(&projectDetail.AudioConfig).Error
		if err != nil {
			continue
		}
	}
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
func (s *InfoService) TranslateInfoPrompt(infoParams system.Info) error {
	var infoList []system.Info
	var currentSettings system.Settings
	err := global.DB.Model(&system.Settings{}).Preload("OllamaConfig").Preload("AliyunConfig").First(&currentSettings).Error
	if err != nil {
		return errors.New("请先初始化配置")
	}
	if infoParams.ProjectDetailId != 0 {
		err = global.DB.Model(&system.Info{}).Find(&infoList, "project_detail_id = ?", infoParams.ProjectDetailId).Error
	}
	if infoParams.Id != 0 {
		err = global.DB.Model(&system.Info{}).Find(&infoList, "id = ?", infoParams.Id).Error
	}
	var loras []system.StableDiffusionLoras
	err = global.DB.Model(&system.StableDiffusionLoras{}).Find(&loras).Error
	if err != nil {
		return err
	}
	// 如果采用翻译模型
	if currentSettings.TranslateType == "sd-prompt-translator" {
		for _, info := range infoList {
			if err != nil {
				return err
			}
			lorasText := ""
			for _, lora := range loras {
				if info.Role == "" {
					continue
				}
				if strings.Contains(lora.Roles, info.Role) {
					lorasText += lora.Name + ","
				}
			}
			info.Prompt = lorasText
			err = global.DB.Model(&info).Update("prompt", info.Prompt).Error
			if err != nil {
				return err
			}
		}
		return err
	} else if currentSettings.TranslateType == "ollama" {
		//var message []openai.ChatCompletionMessage
		//_, openPromptError := os.ReadFile("/Users/hyouka/Desktop/代码/stable-diffusion-go/server/prompt.txt")
		//if openPromptError != nil {
		//	return openPromptError
		//}
		//message = append(message, openai.ChatCompletionMessage{
		//	Role:    openai.ChatMessageRoleSystem,
		//	Content: string(promptByte),
		//})
		for _, info := range infoList {
			lorasText := ""
			for _, lora := range loras {
				if info.Role == "" {
					continue
				}
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
			err = global.DB.Model(&info).Update("prompt", info.Prompt).Error
			if err != nil {
				return err
			}
		}
		return err
	} else if currentSettings.TranslateType == "aliyun" {
		for _, info := range infoList {
			lorasText := ""
			for _, lora := range loras {
				if info.Role == "" {
					continue
				}
				if strings.Contains(lora.Roles, info.Role) {
					lorasText += lora.Name + ","
				}
			}
			prompt, _ := source.TranslateAliyun(info.Text, currentSettings.AliyunConfig)
			if lorasText != "" {
				info.Prompt = prompt + "," + lorasText
			} else {
				info.Prompt = prompt
			}
			err = global.DB.Model(&info).Update("prompt", info.Prompt).Error
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("请选择正确的翻译配置")
	}
	return err
}
