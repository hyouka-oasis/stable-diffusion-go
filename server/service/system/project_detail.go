package system

import (
	"bufio"
	"errors"
	"fmt"
	"github/stable-diffusion-go/server/config"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/python_core"
	"github/stable-diffusion-go/server/source"
	"github/stable-diffusion-go/server/utils"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
	"path/filepath"
)

type ProjectDetailService struct{}

// UploadProjectDetailFile 上传文件并且处理分词
func (s *ProjectDetailService) UploadProjectDetailFile(id uint, file *multipart.FileHeader, saveType string, whetherParticiple string) (err error) {
	err = taskService.DeleteTaskWhereProjectDetailId(id)
	if err != nil {
		return err
	}
	var taskErrors []system.TaskErrors
	if saveType != "create" && saveType != "update" {
		return errors.New("请选择需要创建还是覆盖")
	}
	tmpFile, err := os.Create(filepath.Join(config.ExecutePath, "participle.py"))
	if err != nil {
		taskErrors = append(taskErrors, system.TaskErrors{Error: "创建python文件失败:" + err.Error()})
		fmt.Println("创建python文件失败:", err)
		return err
	}
	_, err = tmpFile.Write([]byte(python_core.PythonParticiplePythonPath))
	if err != nil {
		taskErrors = append(taskErrors, system.TaskErrors{Error: "写入python内容失败:" + err.Error()})
		fmt.Println("写入python内容失败", err)
		return err
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())
	var projectDetail system.ProjectDetail
	// 这里只更新name
	err = global.DB.Model(&system.ProjectDetail{}).Where("id = ?", id).Updates(&system.ProjectDetail{
		FileName: file.Filename,
	}).Preload("ParticipleConfig").First(&projectDetail).Error
	if err != nil {
		taskErrors = append(taskErrors, system.TaskErrors{Error: "更新name失败:" + err.Error()})
		return err
	}
	if saveType == "create" {
		var infoList []system.Info
		err = global.DB.Model(&system.Info{}).Where("project_detail_id = ?", id).Find(&infoList).Error
		if err != nil {
			taskErrors = append(taskErrors, system.TaskErrors{Error: "查找列表失败:" + err.Error()})
			return errors.New("查找列表失败:" + err.Error())
		}
		for _, info := range infoList {
			err = global.DB.Model(&system.AudioConfig{}).Where("info_id = ?", info.Id).Delete(&system.AudioConfig{}).Error
			if err != nil {
				return errors.New("删除音频配置失败:" + err.Error())
			}
			err = global.DB.Model(&system.StableDiffusionImages{}).Where("info_id = ?", info.Id).Delete(&system.StableDiffusionImages{}).Error
			if err != nil {
				return errors.New("删除图片失败:" + err.Error())
			}
		}
		// 将原有的全部删除
		err = global.DB.Delete(&system.Info{}, "project_detail_id = ?", id).Error
		if err != nil {
			taskErrors = append(taskErrors, system.TaskErrors{Error: "删除原有项目失败:" + err.Error()})
			return errors.New("删除原有项目失败:" + err.Error())
		}
	}
	filePath := global.Config.Local.Path + "/" + file.Filename
	outParticipleBookPathBookPath := global.Config.Local.Path + "/" + global.ParticipleBookName
	err = utils.UploadFileToLocal(file, filePath)
	if err != nil {
		taskErrors = append(taskErrors, system.TaskErrors{Error: "处理文件失败:" + err.Error()})
		return errors.New("处理文件失败:" + err.Error())
	}
	splitTextError := source.SplitText(projectDetail, whetherParticiple, tmpFile.Name())
	if splitTextError != nil {
		taskErrors = append(taskErrors, system.TaskErrors{Error: "进行分词失败:" + err.Error()})
		return errors.New("进行分词失败:" + splitTextError.Error())
	}
	// 打开文件
	var participleBook *os.File
	participleBook, err = os.Open(outParticipleBookPathBookPath)
	if err != nil {
		taskErrors = append(taskErrors, system.TaskErrors{Error: "打开文件失败:" + err.Error()})
		return errors.New("打开文件失败:" + err.Error())
	}
	defer participleBook.Close() // 确保文件在函数退出时被关闭
	// 创建扫描器
	scanner := bufio.NewScanner(participleBook)
	var infoList []system.Info
	// 逐行读取并输出
	for scanner.Scan() {
		infoList = append(infoList, system.Info{
			ProjectDetailId: id,
			Text:            scanner.Text(),
		})
	}
	err = global.DB.Model(&system.Info{}).Create(&infoList).Error
	if err != nil {
		taskErrors = append(taskErrors, system.TaskErrors{Error: "写入列表失败:" + err.Error()})
		return errors.New("写入列表失败:" + err.Error())
	}
	// 创建任务
	task := system.Task{
		ProjectDetailId: projectDetail.Id,
		Progress:        0,
		Status:          system.START,
		Errors:          taskErrors,
	}
	if len(taskErrors) != 0 {
		task.Status = system.REJECTED
	}
	_, err = taskService.CreateTask(task)
	if err != nil {
		return err
	}
	err = global.DB.Model(&system.Info{}).Where("project_detail_id = ?", id).Find(&infoList).Error
	if err != nil {
		return errors.New("写入列表失败:" + err.Error())
	}
	var audioConfigList []system.AudioConfig
	for _, info := range infoList {
		audioConfigList = append(audioConfigList, system.AudioConfig{
			InfoId: info.Id,
		})
	}
	err = global.DB.Model(&system.AudioConfig{}).Create(&audioConfigList).Error
	if err != nil {
		return errors.New("更新音频配置失败:" + err.Error())
	}
	os.Remove(filePath)
	os.Remove(outParticipleBookPathBookPath)
	err = taskService.UpdateTask(system.Task{
		ProjectDetailId: id,
		Status:          system.RESOLVED,
	})
	if err != nil {
		return err
	}
	return err
}

// GetProjectDetail 获取项目详情
func (s *ProjectDetailService) GetProjectDetail(config system.ProjectDetail) (detail system.ProjectDetail, err error) {
	err = global.DB.Preload("StableDiffusionConfig").Preload("StableDiffusionConfig.OverrideSettings").Preload("VideoConfig").Preload("ParticipleConfig").Preload("AudioConfig").Preload("InfoList").Preload("InfoList.StableDiffusionImages").Preload("InfoList.AudioConfig").Preload("InfoList.VideoConfig").Model(&system.ProjectDetail{}).Where("id = ?", config.Id).First(&detail).Error
	return
}

// UpdateProjectDetail 更新项目详情
func (s *ProjectDetailService) UpdateProjectDetail(config request.UpdateProjectDetailRequestParams) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&system.ProjectDetail{}).Where("id = ?", config.Id).Updates(&config).Error
		if err != nil {
			return err
		}
		err = tx.Model(&system.ProjectDetail{}).Where("id = ?", config.Id).Update("open_context", config.OpenContext).Update("open_subtitles", config.OpenSubtitles).Update("break_audio", config.BreakAudio).Update("break_video", config.BreakVideo).Update("concat_audio", config.ConcatAudio).Update("concat_video", config.ConcatVideo).Error
		if err != nil {
			return err
		}
		// 更新分词
		if config.ParticipleConfig != (system.ParticipleConfig{}) {
			err = tx.Model(&system.ParticipleConfig{}).Where("project_detail_id = ?", config.Id).Updates(&config.ParticipleConfig).Error
			if err != nil {
				return err
			}
		}
		// 更新分词
		if config.AudioConfig != (system.AudioConfig{}) {
			err = tx.Model(&system.AudioConfig{}).Where("project_detail_id = ?", config.Id).Updates(&config.AudioConfig).Error
			if err != nil {
				return err
			}
		}
		// 更新视频配置
		if config.VideoConfig != (system.VideoConfig{}) {
			err = tx.Model(&system.VideoConfig{}).Where("project_detail_id = ?", config.Id).Updates(&config.VideoConfig).Update("open_animation", config.VideoConfig.OpenAnimation).Error
			if err != nil {
				return err
			}
		}
		// 更新stable-diffusion配置
		if config.StableDiffusionConfig != (system.StableDiffusionSettings{}) {
			err = tx.Model(&system.StableDiffusionSettings{}).Where("project_detail_id = ?", config.Id).Updates(&config.StableDiffusionConfig).Error
			if err != nil {
				return err
			}
		}
		// 更新stable-diffusion配置
		if config.StableDiffusionConfig.OverrideSettings != (system.StableDiffusionOverrideSettings{}) {
			err = tx.Model(&system.StableDiffusionOverrideSettings{}).Where("project_detail_id = ?", config.Id).Updates(&config.StableDiffusionConfig.OverrideSettings).Error
			if err != nil {
				return err
			}
		}
		return err
	})
}

// DeleteProjectDetail 删除项目详情
func (s *ProjectDetailService) DeleteProjectDetail(projectDetailId uint) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Delete(&system.ProjectDetail{}, "id = ?", projectDetailId).Error
		if err != nil {
			return err
		}
		err = tx.Where("project_detail_id = ?", projectDetailId).Delete(&system.VideoConfig{}).Error
		if err != nil {
			return err
		}
		err = tx.Where("project_detail_id = ?", projectDetailId).Delete(&system.ParticipleConfig{}).Error
		if err != nil {
			return err
		}
		err = tx.Where("project_detail_id = ?", projectDetailId).Delete(&system.AudioConfig{}).Error
		if err != nil {
			return err
		}
		err = tx.Where("project_detail_id = ?", projectDetailId).Delete(&system.Info{}).Error
		if err != nil {
			return err
		}
		err = tx.Where("project_detail_id = ?", projectDetailId).Delete(&system.StableDiffusionImages{}).Error
		if err != nil {
			return err
		}
		// 返回 nil 提交事务
		return err
	})
}

// CreateProjectDetail 创建项目详情
func (s *ProjectDetailService) CreateProjectDetail(projectId uint) (projectDetail system.ProjectDetail, err error) {
	return projectDetail, global.DB.Transaction(func(tx *gorm.DB) error {
		projectDetail = system.ProjectDetail{
			ProjectId: projectId,
		}
		// 同时创建项目详情
		err = tx.Create(&projectDetail).Error
		if err != nil {
			return err
		}
		participleConfig := system.ParticipleConfig{
			ProjectDetailId: projectDetail.Id,
		}
		// 同时创建项目详情分词配置
		err = tx.Create(&participleConfig).Error
		if err != nil {
			return err
		}
		audioConfig := system.AudioConfig{
			ProjectDetailId: projectDetail.Id,
		}
		// 同时创建项目详情音频配置
		err = tx.Create(&audioConfig).Error
		if err != nil {
			return err
		}
		videoConfig := system.VideoConfig{
			ProjectDetailId: projectDetail.Id,
		}
		// 同时创建项目详情视频配置
		err = tx.Create(&videoConfig).Error
		if err != nil {
			return err
		}
		stableDiffusionSettings := system.StableDiffusionSettings{
			ProjectDetailId: projectDetail.Id,
		}
		// 同时创建项目stable-diffusion配置
		err = tx.Create(&stableDiffusionSettings).Error
		if err != nil {
			return err
		}
		stableDiffusionOverrideSettings := system.StableDiffusionOverrideSettings{
			ProjectDetailId:           projectDetail.Id,
			StableDiffusionSettingsId: stableDiffusionSettings.Id,
		}
		// 同时创建项目stable-diffusion配置
		err = tx.Create(&stableDiffusionOverrideSettings).Error
		if err != nil {
			return err
		}
		return err
	})
}
