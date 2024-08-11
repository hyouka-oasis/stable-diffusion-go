package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/example"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/source"
	"github/stable-diffusion-go/server/utils"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type VideoService struct{}

// CreateVideo 批量文字转图片
func (s *VideoService) CreateVideo(infoParams request.InfoCreateVideoRequest) (err error) {
	var settings system.Settings
	err = global.DB.Model(&system.Settings{}).First(&settings).Error
	if err != nil {
		return errors.New("请先初始化配置")
	}
	var projectDetail system.ProjectDetail
	err = global.DB.Model(&system.ProjectDetail{}).Preload("VideoConfig").First(&projectDetail, "id = ?", infoParams.ProjectDetailId).Error
	if err != nil {
		return errors.New("查找项目详情失败:" + err.Error())
	}
	var project system.Project
	err = global.DB.Model(&system.Project{}).First(&project, "id = ?", projectDetail.ProjectId).Error
	if err != nil {
		return errors.New("查找项目失败:" + err.Error())
	}
	filename := strings.TrimSuffix(projectDetail.FileName, path.Ext(projectDetail.FileName))
	projectPath := path.Join(settings.SavePath, project.Name+"-"+strconv.Itoa(int(project.Id)), filename)
	err = utils.EnsureDirectory(projectPath)
	if err != nil {
		return err
	}
	var infoList []system.Info
	if len(infoParams.Ids) != 0 {
		var info system.Info
		for _, id := range infoParams.Ids {
			err = global.DB.Model(&system.Info{}).Find(&info, "id = ?", id).Error
			if err != nil {
				continue
			}
			infoList = append(infoList, info)
		}
	} else {
		err = global.DB.Model(&system.Info{}).Find(&infoList, "project_detail_id = ?", infoParams.ProjectDetailId).Error
		if err != nil {
			return nil
		}
	}
	var videoList []string
	var videoSubtitleList []string
	for _, info := range infoList {
		if info.StableDiffusionImageId == 0 {
			continue
		}
		savePath := path.Join(projectPath, strconv.Itoa(int(info.Id)))
		err = utils.EnsureDirectory(savePath)
		if err != nil {
			continue
		}
		var image example.ExaFileUploadAndDownload
		err = global.DB.Model(&example.ExaFileUploadAndDownload{}).Where("id = ?", info.StableDiffusionImageId).First(&image).Error
		if err != nil {
			continue
		}
		stableDiffusionParams := map[string]interface{}{}
		stableDiffusionRequest := map[string]interface{}{}
		err = json.Unmarshal([]byte(projectDetail.StableDiffusionConfig), &stableDiffusionParams)
		if err != nil {
			continue
		}
		// 如果json解析成功则合并 Stable Diffusion 配置参数
		for key, value := range stableDiffusionParams {
			stableDiffusionRequest[key] = value
		}
		width := utils.GetInterfaceToInt(stableDiffusionRequest["width"])
		height := utils.GetInterfaceToInt(stableDiffusionRequest["height"])
		params := source.DisposableSynthesisVideoParams{
			SavePath:                 savePath,
			Info:                     info,
			ExaFileUploadAndDownload: image,
			Width:                    width,
			Height:                   height,
			BreakVideo:               projectDetail.BreakVideo,
			OpenSubtitles:            projectDetail.OpenSubtitles,
			VideoConfig:              projectDetail.VideoConfig,
		}
		if len(infoParams.Ids) != 0 {
			params.BreakVideo = false
		}
		params.Name = filename + "-" + strconv.Itoa(int(info.Id))
		params.VideoPath = filepath.Join(savePath, params.Name+".mp4")
		params.VideoSubtitlePath = filepath.Join(params.SavePath, params.Name+"subtitle.mp4")
		err = source.DisposableSynthesisVideo(params)
		if err != nil {
			fmt.Println(err.Error(), "生成视频错误")
			continue
		}
		videoList = append(videoList, params.VideoPath)
		videoSubtitleList = append(videoSubtitleList, params.VideoSubtitlePath)
	}
	if projectDetail.ConcatVideo && len(infoParams.Ids) == 0 {
		outVideoPath := path.Join(projectPath, filename+".mp4")
		err = source.MergeVideoList(videoList, outVideoPath)
		if err != nil {
			return err
		}
		if projectDetail.OpenSubtitles {
			outVideoSubtitlePath := path.Join(projectPath, filename+"subtitle.mp4")
			err = source.MergeVideoList(videoSubtitleList, outVideoSubtitlePath)
			if err != nil {
				return err
			}
		}
	}
	return err
}
