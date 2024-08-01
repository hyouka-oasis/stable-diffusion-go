package system

import (
	"errors"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/source"
	"github/stable-diffusion-go/server/utils"
	"path"
	"path/filepath"
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
	filename := strings.TrimSuffix(projectDetail.FileName, path.Ext(projectDetail.FileName))
	savePath := path.Join(filepath.ToSlash(settings.SavePath), filename)
	err = utils.EnsureDirectory(savePath)
	if err != nil {
		return errors.New("创建目录失败:" + err.Error())
	}
	err = source.CreateAudioAndSrt(savePath, filename, projectDetail)
	if err != nil {
		return errors.New("生成音频和字幕失败:" + err.Error())
	}
	return nil
}
