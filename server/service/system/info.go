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
	"sync"
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
	var projectDetailParticipleList []system.Info
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var currentSettings system.Settings
		err := tx.Model(&system.Settings{}).Preload("OllamaConfig").First(&currentSettings).Error
		if err != nil {
			return errors.New("请先初始化配置")
		}
		if projectDetailParticipleParams.ProjectDetailId != 0 {
			err = tx.Model(&system.Info{}).Find(&projectDetailParticipleList, "project_detail_id = ?", projectDetailParticipleParams.ProjectDetailId).Error
		}
		if projectDetailParticipleParams.Id != 0 {
			err = tx.Model(&system.Info{}).Find(&projectDetailParticipleList, "id = ?", projectDetailParticipleParams.Id).Error
		}
		if currentSettings.TranslateType == "sd-prompt-translator" {
			return errors.New("当前配置不需要进行prompt转换")
		}
		if currentSettings.TranslateType == "ollama" {
			var wg sync.WaitGroup
			wg.Wait()
			// 异步处理翻译
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, projectDetailParticiple := range projectDetailParticipleList {
					lorasText := ""
					var loras []system.StableDiffusionLoras
					err = tx.Model(&system.StableDiffusionLoras{}).Find(&loras).Error
					if err != nil {
						return
					}
					for _, lora := range loras {
						if strings.Contains(lora.Roles, projectDetailParticiple.Role) {
							lorasText += lora.Name + ","
						}
					}

					prompt, _ := source.ChatgptOllama(projectDetailParticiple.Text, currentSettings.OllamaConfig)
					if lorasText != "" {
						projectDetailParticiple.Prompt = prompt + "," + lorasText
					} else {
						projectDetailParticiple.Prompt = prompt
					}
					err = tx.Model(&projectDetailParticiple).Select("prompt").Updates(&projectDetailParticiple).Error
				}
			}()
			wg.Wait()
		}
		return err
	})
}
