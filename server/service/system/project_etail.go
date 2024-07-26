package system

import (
	"bufio"
	"errors"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
)

type ProjectDetailService struct{}

// UpdateProjectDetailFile 上传文件并且处理分词
func (s *ProjectDetailService) UpdateProjectDetailFile(id uint, config system.ProjectDetail, file *multipart.FileHeader) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&system.ProjectDetail{}).Where("id = ?", id).Updates(&config).Error
		if err != nil {
			return err
		}
		err = tx.Model(&system.ProjectDetailParticiple{}).Where("project_detail_id = ?", id).Updates(&config.Participle).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&system.ProjectDetailParticipleList{}, "project_detail_id = ?", id).Error
		if err != nil {
			return errors.New("删除原有项目失败:" + err.Error())
		}
		filePath := global.Config.Local.Path + "/" + file.Filename
		outParticipleBookPathBookPath := global.Config.Local.Path + "/" + "participleBook.txt"
		err = utils.SplitTextUploadFileToLocal(file, filePath, config)
		if err != nil {
			return errors.New("处理文件失败:" + err.Error())
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
		var projectFormDetail []system.ProjectDetailParticipleList
		// 逐行读取并输出
		for scanner.Scan() {
			projectFormDetail = append(projectFormDetail, system.ProjectDetailParticipleList{
				ProjectDetailId: id,
				Text:            scanner.Text(),
			})
		}
		err = tx.Model(&system.ProjectDetailParticipleList{}).Create(&projectFormDetail).Error
		if err != nil {
			return errors.New("写入列表失败:" + err.Error())
		}
		return nil
	})
	return err
}

// GetProjectDetail 获取项目详情
func (s *ProjectDetailService) GetProjectDetail(config system.ProjectDetail) (detail system.ProjectDetail, err error) {
	err = global.DB.Preload("Participle").Preload("ParticipleList").Model(&system.ProjectDetail{}).Where("project_id = ?", config.ProjectId).First(&detail).Error
	return
}
