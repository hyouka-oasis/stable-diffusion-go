package system

import (
	"errors"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/example"
	"github/stable-diffusion-go/server/model/system"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/source"
	"strconv"
)

type StableDiffusionImagesService struct{}

// StableDiffusionTextToImage 批量文字转图片
func (s *StableDiffusionImagesService) StableDiffusionTextToImage(params systemRequest.StableDiffusionRequestParams) (images []string, err error) {
	var settings system.Settings
	err = global.DB.Model(&system.Settings{}).Preload("StableDiffusionConfig").First(&settings).Error
	if err != nil {
		return images, errors.New("请先配置")
	}
	if settings.StableDiffusionConfig.Url == "" {
		return images, errors.New("stable-diffusion-url不能为空")
	}
	var projectDetail system.ProjectDetail
	err = global.DB.Preload("StableDiffusionConfig").Preload("StableDiffusionConfig.OverrideSettings").Model(&system.ProjectDetail{}).Where("id = ?", params.ProjectDetailId).First(&projectDetail).Error
	if err != nil {
		return images, errors.New("获取项目详情失败")
	}
	err = taskService.DeleteTaskWhereProjectDetailId(projectDetail.Id)
	if err != nil {
		return
	}
	var taskErrors []system.TaskErrors
	// 创建任务
	task := system.Task{
		ProjectDetailId: projectDetail.Id,
		Progress:        0,
		Status:          system.START,
		Errors:          taskErrors,
	}
	systemTask, err := taskService.CreateTask(task)
	if err != nil {
		return
	}
	for index, infoId := range params.Ids {
		if err != nil {
			taskErrors = append(taskErrors, system.TaskErrors{Error: "生成图片错误" + err.Error()})
			continue
		}
		// 异步处理翻译
		var info system.Info
		// 查到单个的列表
		err = global.DB.Model(&system.Info{}).Where("id = ?", infoId).Find(&info).Error
		if err != nil {
			taskErrors = append(taskErrors, system.TaskErrors{Error: "生成图片错误" + err.Error()})
			continue
		}
		if info.Prompt == "" {
			projectDetail.StableDiffusionConfig.Prompt = info.Text
		} else {
			projectDetail.StableDiffusionConfig.Prompt = info.Prompt
		}
		if info.NegativePrompt != "" {
			projectDetail.StableDiffusionConfig.NegativePrompt = info.NegativePrompt
		}
		err = taskService.UpdateTask(system.Task{
			Model: global.Model{
				Id: systemTask.Id,
			},
			Progress: float64(index+1) / float64(len(params.Ids)),
			Message:  strconv.Itoa(index+1) + "/" + strconv.Itoa(len(params.Ids)),
		})
		if err != nil {
			continue
		}
		apiUrl := settings.StableDiffusionConfig.Url + "/sdapi/v1/txt2img"
		stableDiffusionImages, generateError := source.StableDiffusionGenerateImage(apiUrl, projectDetail.StableDiffusionConfig)
		if generateError != nil {
			taskErrors = append(taskErrors, system.TaskErrors{Error: "生成图片错误" + generateError.Error()})
			err = generateError
			continue
		}
		images = stableDiffusionImages
	}
	err = taskService.UpdateTask(system.Task{
		Model: global.Model{
			Id: systemTask.Id,
		},
		Status:   system.RESOLVED,
		Progress: 1,
		Errors:   taskErrors,
	})
	if err != nil {
		return
	}
	return images, err
}

// DeleteStableDiffusionImage 批量删除图片
func (s *StableDiffusionImagesService) DeleteStableDiffusionImage(params request.IdsReq) error {
	err := global.DB.Model(&system.StableDiffusionImages{}).Error
	if err != nil {
		return err
	}
	for _, id := range params.Ids {
		var stableDiffusionImages system.StableDiffusionImages
		err = global.DB.Where("id = ?", id).First(&stableDiffusionImages).Error
		if err != nil {
			return err
		}
		var info system.Info
		err = global.DB.Where("id = ?", stableDiffusionImages.InfoId).First(&info).Error
		if err != nil {
			return err
		}
		if info.StableDiffusionImageId == id {
			err = global.DB.Model(&info).Update("stable_diffusion_image_id", 0).Error
		}
		err = global.DB.Where("id = ?", id).Delete(&system.StableDiffusionImages{}).Error
		if err != nil {
			return err
		}
		err = global.DB.Where("id = ?", stableDiffusionImages.FileId).Delete(&example.ExaFileUploadAndDownload{}).Error
		if err != nil {
			return err
		}
	}
	return err
}

// AddStableDiffusionImage 添加图片
func (s *StableDiffusionImagesService) AddStableDiffusionImage(params system.StableDiffusionImages) error {
	err := global.DB.Model(&system.StableDiffusionImages{}).Create(&params).Error
	return err
}
