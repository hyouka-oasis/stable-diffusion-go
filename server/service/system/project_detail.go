package system

import (
	"bufio"
	"errors"
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/model/system/request"
	"github/stable-diffusion-go/server/source"
	"github/stable-diffusion-go/server/utils"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
)

type ProjectDetailService struct{}

// UploadProjectDetailFile 上传文件并且处理分词
func (s *ProjectDetailService) UploadProjectDetailFile(id uint, file *multipart.FileHeader, saveType string, whetherParticiple string) (err error) {
	if saveType != "create" && saveType != "push" {
		return errors.New("请选择需要创建还是覆盖")
	}
	var projectDetail system.ProjectDetail
	// 这里只更新name
	err = global.DB.Model(&system.ProjectDetail{}).Where("id = ?", id).Updates(&system.ProjectDetail{
		FileName: file.Filename,
	}).Preload("ParticipleConfig").First(&projectDetail).Error
	if err != nil {
		return err
	}
	if saveType == "create" {
		var infoList []system.Info
		err = global.DB.Model(&system.Info{}).Where("project_detail_id = ?", id).Find(&infoList).Error
		if err != nil {
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
			return errors.New("删除原有项目失败:" + err.Error())
		}
	}
	filePath := global.Config.Local.Path + "/" + file.Filename
	outParticipleBookPathBookPath := global.Config.Local.Path + "/" + "participleBook.txt"
	err = utils.UploadFileToLocal(file, filePath)
	if err != nil {
		return errors.New("处理文件失败:" + err.Error())
	}
	splitTextError := source.SplitText(projectDetail, whetherParticiple)
	if splitTextError != nil {
		return errors.New("进行分词失败:" + splitTextError.Error())
	}
	// 打开文件
	var participleBook *os.File
	participleBook, err = os.Open(outParticipleBookPathBookPath)
	if err != nil {
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
		return errors.New("写入列表失败:" + err.Error())
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
	return err
}

// GetProjectDetail 获取项目详情
func (s *ProjectDetailService) GetProjectDetail(config system.ProjectDetail) (detail system.ProjectDetail, err error) {
	err = global.DB.Preload("ParticipleConfig").Preload("AudioConfig").Preload("InfoList").Preload("InfoList.StableDiffusionImages").Preload("InfoList.AudioConfig").Model(&system.ProjectDetail{}).Where("project_id = ?", config.ProjectId).First(&detail).Error
	return
}

// UpdateProjectDetail 更新项目详情
func (s *ProjectDetailService) UpdateProjectDetail(config request.UpdateProjectDetailRequestParams) (err error) {
	fmt.Println(&config)
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&system.ProjectDetail{}).Where("id = ?", config.Id).Updates(&config).Error
		if err != nil {
			return err
		}
		err = tx.Model(&system.ProjectDetail{}).Where("id = ?", config.Id).Update("break_audio", config.BreakAudio).Error
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
			return err
		}
		return err
	})
}
