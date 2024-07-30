package system

import (
	"bufio"
	"errors"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/source"
	"github/stable-diffusion-go/server/utils"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
)

type ProjectDetailService struct{}

// UploadProjectDetailFile 上传文件并且处理分词
func (s *ProjectDetailService) UploadProjectDetailFile(id uint, config system.ProjectDetail, file *multipart.FileHeader) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&system.ProjectDetail{}).Where("id = ?", id).Updates(&config).Error
		if err != nil {
			return err
		}
		//var projectDetailParticiple system.ProjectDetailParticiple
		// 先查找是否存在分词配置
		//err = tx.Model(&system.ProjectDetailParticiple{}).Where("project_detail_id = ?", id).First(projectDetailParticiple).Error
		// 如果不存在则保存
		//if err != nil {
		//	err = tx.Create(&config.ParticipleConfig).Error
		//	if err != nil {
		//		return errors.New("更新分词配置失败" + err.Error())
		//	}
		//}
		// 如果存在则更新
		err = tx.Model(&system.ProjectDetailParticiple{}).Where("project_detail_id = ?", id).Updates(&config.ParticipleConfig).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&system.ProjectDetailInfo{}, "project_detail_id = ?", id).Error
		if err != nil {
			return errors.New("删除原有项目失败:" + err.Error())
		}
		filePath := global.Config.Local.Path + "/" + file.Filename
		outParticipleBookPathBookPath := global.Config.Local.Path + "/" + "participleBook.txt"
		err = utils.UploadFileToLocal(file, filePath)
		if err != nil {
			return errors.New("处理文件失败:" + err.Error())
		}
		splitTextError := source.SplitText(config)
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
		var projectFormDetail []system.ProjectDetailInfo
		// 逐行读取并输出
		for scanner.Scan() {
			projectFormDetail = append(projectFormDetail, system.ProjectDetailInfo{
				ProjectDetailId: id,
				Text:            scanner.Text(),
			})
		}
		err = tx.Model(&system.ProjectDetailInfo{}).Create(&projectFormDetail).Error
		if err != nil {
			return errors.New("写入列表失败:" + err.Error())
		}
		return nil
	})
}

// GetProjectDetail 获取项目详情
func (s *ProjectDetailService) GetProjectDetail(config system.ProjectDetail) (detail system.ProjectDetail, err error) {
	err = global.DB.Preload("ParticipleConfig").Preload("ProjectDetailInfoList").Model(&system.ProjectDetail{}).Where("project_id = ?", config.ProjectId).First(&detail).Error
	return
}

// UpdateProjectDetail 更新项目详情
func (s *ProjectDetailService) UpdateProjectDetail(config system.ProjectDetail) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&system.ProjectDetail{}).Where("id = ?", config.Id).Updates(&config).Error
		if err != nil {
			return err
		}
		// 更新分词
		if config.ParticipleConfig != (system.ProjectDetailParticiple{}) {
			err = tx.Model(&system.ProjectDetailParticiple{}).Where("project_detail_id = ?", config.Id).Updates(&config.ParticipleConfig).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}
