package system

import (
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github/stable-diffusion-go/server/config"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/python_core"
	"github/stable-diffusion-go/server/source"
	"github/stable-diffusion-go/server/utils"
	"gorm.io/gorm"
	"os"
	"path"
	"path/filepath"
	"strconv"
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
	var infoList []system.Info
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&system.Info{}).Find(&infoList, "project_detail_id = ?", id).Error
		for _, info := range infoList {
			info.Role = utils.CutPos(info.Text)
			err = tx.Model(&info).Select("role").Updates(&info).Error
			return err
		}
		return err
	})
}

// TranslateInfoPrompt 进行prompt转换
func (s *InfoService) TranslateInfoPrompt(infoParams request.InfoTranslateRequest) error {
	var infoList []system.Info
	var currentSettings system.Settings
	err := global.DB.Model(&system.Settings{}).Preload("OllamaConfig").Preload("AliyunConfig").First(&currentSettings).Error
	if err != nil {
		return errors.New("请先初始化配置")
	}
	var projectDetail system.ProjectDetail
	err = global.DB.Model(&system.ProjectDetail{}).Where("id = ?", infoParams.ProjectDetailId).First(&projectDetail).Error
	if err != nil {
		return err
	}
	if infoParams.Id != 0 {
		err = global.DB.Model(&system.Info{}).Find(&infoList, "id = ?", infoParams.Id).Error
	} else {
		err = global.DB.Model(&system.Info{}).Find(&infoList, "project_detail_id = ?", infoParams.ProjectDetailId).Error
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
		var messageList []openai.ChatCompletionMessage
		if projectDetail.PromptText != "" {
			messageList = append(messageList, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleSystem,
				Content: projectDetail.PromptText,
			})
			// 这边先走一遍上下文初始化
			content, err := source.OpenaiClient(currentSettings.OllamaConfig, &messageList)
			if err != nil {
				return errors.New("初始化上下文失败")
			}
			messageList = append(messageList, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: content,
			})
		}
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
			var infoText string
			if info.KeywordsText == "" {
				infoText = info.Text
			} else {
				infoText = info.KeywordsText
			}
			fmt.Println(messageList, "messageList")
			prompt, _ := source.ChatgptOllama(infoText, currentSettings.OllamaConfig, projectDetail.OpenContext, &messageList)
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
			var infoText string
			if info.KeywordsText == "" {
				infoText = info.Text
			} else {
				infoText = info.KeywordsText
			}
			prompt, _ := source.TranslateAliyun(infoText, currentSettings.AliyunConfig)
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

// KeywordExtractionInfo 关键词提取
func (s *InfoService) KeywordExtractionInfo(id uint) error {
	tmpFile, err := os.Create(filepath.Join(config.ExecutePath, "keyword_extraction.py"))
	if err != nil {
		fmt.Println("创建python文件失败:", err)
		return err
	}
	_, err = tmpFile.Write([]byte(python_core.PythonKeywordExtractionPath))
	if err != nil {
		fmt.Println("写入python内容失败", err)
		return err
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())
	var settings system.Settings
	err = global.DB.Model(&system.Settings{}).First(&settings).Error
	if err != nil {
		return errors.New("请先初始化配置")
	}
	var projectDetail system.ProjectDetail
	err = global.DB.Preload("AudioConfig").Model(&system.ProjectDetail{}).Where("id = ?", id).First(&projectDetail).Error
	if err != nil {
		return errors.New("查找项目详情失败:" + err.Error())
	}
	var project system.Project
	err = global.DB.Model(&system.Project{}).Where("id = ?", projectDetail.ProjectId).First(&project).Error
	if err != nil {
		return errors.New("查找项目失败:" + err.Error())
	}
	var infoList []system.Info
	err = global.DB.Model(&system.Info{}).Where("project_detail_id = ?", id).Find(&infoList).Error
	if err != nil {
		fmt.Println("查找列表失败", err)
		return errors.New("查找列表失败")
	}
	filename := strings.TrimSuffix(projectDetail.FileName, path.Ext(projectDetail.FileName))
	projectPath := path.Join(settings.SavePath, project.Name+"-"+strconv.Itoa(int(project.Id)), filename)
	err = utils.EnsureDirectory(projectPath)
	if err != nil {
		return err
	}
	for _, info := range infoList {
		name := filename + "-" + strconv.Itoa(int(info.Id))
		savePath := path.Join(projectPath, strconv.Itoa(int(info.Id)))
		keywordPath := path.Join(savePath, name+"-keyword.txt")
		err = utils.EnsureDirectory(savePath)
		if err != nil {
			continue
		}
		err = source.KeywordExtraction(info.Text, tmpFile.Name(), keywordPath)
		if err != nil {
			continue
		}
		// 打开文件
		participleBook, readError := os.ReadFile(keywordPath)
		if readError != nil {
			continue
		}
		err = global.DB.Model(&info).Update("keywords_text", string(participleBook)).Error
		if readError != nil {
			continue
		}
	}
	return err
}
