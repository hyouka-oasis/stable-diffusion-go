package system

import (
	"errors"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/source"
	"github/stable-diffusion-go/server/utils"
	"gorm.io/gorm"
	"strings"
	"sync"
)

type ProjectDetailParticipleInfoService struct{}

// DeleteProjectDetailInfo 删除单条记录
func (s *ProjectDetailParticipleInfoService) DeleteProjectDetailInfo(id uint) error {
	err := global.DB.Delete(&system.ProjectDetailInfo{}, "id = ?", id).Error
	return err
}

// UpdateProjectDetailInfo 更新单条记录
func (s *ProjectDetailParticipleInfoService) UpdateProjectDetailInfo(updateData system.ProjectDetailInfo) error {
	err := global.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updateData).Error
	return err
}

// GetProjectDetailInfo 获取单条记录
func (s *ProjectDetailParticipleInfoService) GetProjectDetailInfo(id uint) (info system.ProjectDetailInfo, err error) {
	err = global.DB.Model(&system.ProjectDetailInfo{}).Where("id = ?", id).First(&info).Error
	return
}

// ExtractTheRoleProjectDetailInfoList 进行人物提取
func (s *ProjectDetailParticipleInfoService) ExtractTheRoleProjectDetailInfoList(id uint) error {
	var currentProjectDetailParticipleList []system.ProjectDetailInfo
	//var currentSettings system.Settings
	//err := global.DB.Model(&system.Settings{}).First(&currentSettings).Error
	//if err != nil {
	//	return errors.New("请先初始化配置")
	//}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&system.ProjectDetailInfo{}).Find(&currentProjectDetailParticipleList, "project_detail_id = ?", id).Error
		for _, projectDetailParticiple := range currentProjectDetailParticipleList {
			projectDetailParticiple.Role = utils.CutPos(projectDetailParticiple.Text)
			err = tx.Model(&projectDetailParticiple).Select("role").Updates(&projectDetailParticiple).Error
			return err
		}
		return err
	})
}

// TranslateProjectDetailInfoList 进行prompt转换
func (s *ProjectDetailParticipleInfoService) TranslateProjectDetailInfoList(projectDetailParticipleParams system.ProjectDetailInfo) error {
	var projectDetailParticipleList []system.ProjectDetailInfo
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var currentSettings system.Settings
		err := tx.Model(&system.Settings{}).Preload("OllamaConfig").First(&currentSettings).Error
		if err != nil {
			return errors.New("请先初始化配置")
		}
		if projectDetailParticipleParams.ProjectDetailId != 0 {
			err = tx.Model(&system.ProjectDetailInfo{}).Find(&projectDetailParticipleList, "project_detail_id = ?", projectDetailParticipleParams.ProjectDetailId).Error
		}
		if projectDetailParticipleParams.Id != 0 {
			err = tx.Model(&system.ProjectDetailInfo{}).Find(&projectDetailParticipleList, "id = ?", projectDetailParticipleParams.Id).Error
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
