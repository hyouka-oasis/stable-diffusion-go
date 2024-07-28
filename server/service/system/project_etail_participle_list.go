package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/source"
	"github/stable-diffusion-go/server/utils"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type ProjectDetailParticipleListService struct{}

// DeleteProjectDetailParticipleListItem 删除单条记录
func (s *ProjectDetailParticipleListService) DeleteProjectDetailParticipleListItem(id uint) error {
	err := global.DB.Delete(&system.ProjectDetailParticipleList{}, "id = ?", id).Error
	return err
}

// ExtractTheCharacterProjectDetailParticipleList 进行人物提取
func (s *ProjectDetailParticipleListService) ExtractTheCharacterProjectDetailParticipleList(id uint) error {
	var currentProjectDetailParticipleList []system.ProjectDetailParticipleList
	//var currentSettings system.Settings
	//err := global.DB.Model(&system.Settings{}).First(&currentSettings).Error
	//if err != nil {
	//	return errors.New("请先初始化配置")
	//}
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&system.ProjectDetailParticipleList{}).Find(&currentProjectDetailParticipleList, "project_detail_id = ?", id).Error
		for _, projectDetailParticiple := range currentProjectDetailParticipleList {
			projectDetailParticiple.Character = utils.CutPos(projectDetailParticiple.Text)
			err = tx.Model(&system.ProjectDetailParticipleList{}).Select("character").Updates(&projectDetailParticiple).Error
			return err
		}
		return err
	})
}

// TranslateProjectDetailParticipleList 进行prompt转换
func (s *ProjectDetailParticipleListService) TranslateProjectDetailParticipleList(id uint, c *gin.Context) error {
	var projectDetailParticipleList []system.ProjectDetailParticipleList

	return global.DB.Transaction(func(tx *gorm.DB) error {
		var currentSettings system.Settings
		err := tx.Model(&system.Settings{}).Preload("OllamaConfig").First(&currentSettings).Error
		if err != nil {
			return errors.New("请先初始化配置")
		}
		err = tx.Model(&system.ProjectDetailParticipleList{}).Find(&projectDetailParticipleList, "project_detail_id = ?", id).Error
		if currentSettings.TranslateType == "sd-prompt-translator" {
			return errors.New("当前配置不需要进行prompt转换")
		}
		if currentSettings.TranslateType == "ollama" {
			var wg sync.WaitGroup
			wg.Wait()
			total := len(projectDetailParticipleList)
			processed := 0
			// 异步处理翻译
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, projectDetailParticiple := range projectDetailParticipleList {
					//lorasText := ""
					//var loras []system.StableDiffusionLoras
					//err = tx.Model(&system.StableDiffusionLoras{}).Find(&loras).Error
					//if err != nil {
					//	return errors.New("获取loras失败")
					//}
					prompt, _ := source.ChatgptOllama(projectDetailParticiple.Text, currentSettings.OllamaConfig)
					projectDetailParticiple.Prompt = prompt
					err = tx.Model(&projectDetailParticiple).Select("prompt").Updates(&projectDetailParticiple).Error
					processed++
					progress := int(float64(processed) / float64(total) * 100)
					c.Writer.Header().Set("X-Progress", strconv.Itoa(progress))
				}
			}()
			wg.Wait()
		}
		return nil
	})
}
