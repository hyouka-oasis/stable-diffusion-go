package system

import (
	"errors"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/source"
	"github/stable-diffusion-go/server/utils"
	"path"
	"strconv"
	"strings"
)

type AudioSrtService struct{}

// CreateAudioAndSrt 批量文字转图片
func (s *AudioSrtService) CreateAudioAndSrt(params systemRequest.AudioSrtRequestParams) error {
	var settings system.Settings
	err := global.DB.Model(&system.Settings{}).First(&settings).Error
	if err != nil {
		return errors.New("请先初始化配置")
	}
	var projectDetail system.ProjectDetail
	err = global.DB.Preload("AudioConfig").Model(&system.ProjectDetail{}).Where("id = ?", params.Id).First(&projectDetail).Error
	if err != nil {
		return errors.New("查找项目详情失败:" + err.Error())
	}
	var project system.Project
	err = global.DB.Model(&system.Project{}).Where("id = ?", projectDetail.ProjectId).First(&project).Error
	if err != nil {
		return errors.New("查找项目失败:" + err.Error())
	}
	var infoList []system.Info
	if params.AudioConfig.InfoId != 0 {
		err = global.DB.Preload("AudioConfig").Model(&system.Info{}).Where("id = ?", params.AudioConfig.InfoId).Find(&infoList).Error
		if err != nil {
			return errors.New("查找项目详情失败:" + err.Error())
		}
	} else {
		err = global.DB.Preload("AudioConfig").Model(&system.Info{}).Where("project_detail_id = ?", projectDetail.Id).Find(&infoList).Error
		if err != nil {
			return errors.New("查找项目详情失败:" + err.Error())
		}
	}
	filename := strings.TrimSuffix(projectDetail.FileName, path.Ext(projectDetail.FileName))
	projectPath := path.Join(settings.SavePath, project.Name, filename)
	err = utils.EnsureDirectory(projectPath)
	if err != nil {
		return errors.New("创建目录失败:" + err.Error())
	}
	for _, info := range infoList {
		err = global.DB.Model(&info).Update("loading", true).Error
		if err != nil {
			return err
		}
		savePath := path.Join(projectPath, strconv.Itoa(int(info.Id)))
		var config source.AudioAndSrtParams
		config.SavePath = savePath
		config.Language = projectDetail.Language
		config.Content = info.Text
		config.BreakAudio = projectDetail.BreakAudio
		if params.AudioConfig.InfoId != 0 {
			config.BreakAudio = false
		}
		err = utils.EnsureDirectory(savePath)
		if err != nil {
			continue
		}
		if info.AudioConfig == (system.AudioConfig{}) {
			config.AudioConfig = projectDetail.AudioConfig
		} else {
			config.AudioConfig = info.AudioConfig
		}
		config.Name = filename + "-" + strconv.Itoa(int(projectDetail.Id)) + "-" + strconv.Itoa(int(info.Id))
		err = source.CreateAudioAndSrt(config)
		err = global.DB.Model(&info).Update("loading", false).Error
		if err != nil {
			continue
		}
	}
	return err
}
