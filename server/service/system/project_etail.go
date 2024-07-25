package system

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"gorm.io/gorm"
)

type ProjectDetailService struct{}

// UpdateProjectDetailFile d
func (s *ProjectDetailService) UpdateProjectDetailFile(id int, config system.ProjectDetail) (err error) {
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Model(&system.ProjectDetail{}).Where("id = ?", id).Updates(&config).Error
		if err != nil {
			return err
		}
		err = global.DB.Model(&system.ProjectDetailPotential{}).Where("project_detail_id = ?", id).Updates(&config.Potential).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetProjectDetail 获取项目详情
func (s *ProjectDetailService) GetProjectDetail(config system.ProjectDetail) (detail system.ProjectDetail, err error) {
	err = global.DB.Preload("Potential").Model(&system.ProjectDetail{}).Where("project_id = ?", config.ProjectId).First(&detail).Error
	return
}
