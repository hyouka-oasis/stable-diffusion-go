package system

import (
	"errors"
	"fmt"
	"github/stable-diffusion-go/server/config"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/python_core"
	"github/stable-diffusion-go/server/source"
	"github/stable-diffusion-go/server/utils"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type AudioSrtService struct{}

// CreateAudioAndSrt 生成字幕和音频
func (s *AudioSrtService) CreateAudioAndSrt(params systemRequest.AudioSrtRequestParams) error {
	tmpFile, err := os.Create(filepath.Join(config.ExecutePath, "voice-caption.py"))
	if err != nil {
		fmt.Println("创建python文件失败:", err)
		return err
	}
	_, err = tmpFile.Write([]byte(python_core.PythonVoiceCaptionPath))
	if err != nil {
		fmt.Println("写入python内容失败", err)
		return nil
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())
	var settings system.Settings
	err = global.DB.Model(&system.Settings{}).First(&settings).Error
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
	if params.InfoId != 0 {
		err = global.DB.Preload("AudioConfig").Model(&system.Info{}).Where("id = ?", params.InfoId).Find(&infoList).Error
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
	projectPath := path.Join(settings.SavePath, project.Name+"-"+strconv.Itoa(int(project.Id)), filename)
	err = utils.EnsureDirectory(projectPath)
	if err != nil {
		return errors.New("创建目录失败:" + err.Error())
	}
	err = taskService.DeleteTaskWhereProjectDetailId(projectDetail.Id)
	if err != nil {
		return err
	}
	var taskErrors []system.TaskErrors
	// 创建任务
	task := system.Task{
		ProjectDetailId: projectDetail.Id,
		Progress:        0,
		Status:          system.START,
		Errors:          taskErrors,
		Message:         "正在进行文本转语音",
	}
	systemTask, err := taskService.CreateTask(task)
	if err != nil {
		return err
	}
	var audioPathList []string
	for index, info := range infoList {
		savePath := path.Join(projectPath, strconv.Itoa(int(info.Id)))
		var audioConfig source.AudioAndSrtParams
		audioConfig.SavePath = savePath
		audioConfig.Language = projectDetail.Language
		audioConfig.Content = info.Text
		audioConfig.BreakAudio = projectDetail.BreakAudio
		if params.AudioConfig.InfoId != 0 {
			audioConfig.BreakAudio = false
		}
		err = utils.EnsureDirectory(savePath)
		if err != nil {
			taskErrors = append(taskErrors, system.TaskErrors{Error: "文字转语音失败:" + err.Error()})
			continue
		}
		if info.AudioConfig == (system.AudioConfig{}) {
			audioConfig.AudioConfig = projectDetail.AudioConfig
		} else {
			audioConfig.AudioConfig = info.AudioConfig
		}
		audioConfig.Name = filename + "-" + strconv.Itoa(int(info.Id))
		audioConfig.AudioPath = path.Join(audioConfig.SavePath, audioConfig.Name+".mp3")
		err = taskService.UpdateTask(system.Task{
			Model: global.Model{
				Id: systemTask.Id,
			},
			Progress: float64(index+1) / float64(len(infoList)),
			Message:  strconv.Itoa(index+1) + "/" + strconv.Itoa(len(infoList)),
		})
		if err != nil {
			return err
		}
		err = source.CreateAudioAndSrt(audioConfig, tmpFile.Name())
		if err != nil {
			taskErrors = append(taskErrors, system.TaskErrors{Error: "文字转语音失败:" + err.Error()})
			continue
		}
		audioPathList = append(audioPathList, audioConfig.AudioPath)
	}
	err = taskService.UpdateTask(system.Task{
		Model: global.Model{
			Id: systemTask.Id,
		},
		Status:   system.RESOLVED,
		Progress: 1,
	})
	if err != nil {
		return err
	}
	// 只有合并音频和infoId不存在时代表生成全部音频
	if projectDetail.ConcatAudio && params.InfoId == 0 {
		outAudioPath := path.Join(projectPath, filename+".mp3")
		err = source.MergeAudio(audioPathList, outAudioPath)
		if err != nil {
			return err
		}
	}
	return err
}
