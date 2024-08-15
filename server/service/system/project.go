package system

import (
	"errors"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/common/request"
	"github/stable-diffusion-go/server/model/system"
	"gorm.io/gorm"
)

type ProjectService struct{}

// CreateProject 新增项目
func (s *ProjectService) CreateProject(config system.Project) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var project system.Project
		err = tx.Model(&system.Project{}).Where("name = ?", config.Name).First(&project).Error
		if err == nil {
			return errors.New("存在相同名称项目")
		}
		err = tx.Create(&config).Error
		if err != nil {
			return err
		}
		return err
	})
}

// UpdateProject 更新项目
func (s *ProjectService) UpdateProject(project system.Project) (err error) {
	err = global.DB.Model(&system.Project{}).Where("id = ?", project.Id).Update("name", project.Name).Error
	if err != nil {
		return err
	}
	return err
}

// DeleteProject 删除项目
func (s *ProjectService) DeleteProject(projectId uint) (err error) {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Delete(&system.Project{}, "id = ?", projectId).Error
		if err != nil {
			return err
		}
		var projectDetailList []system.ProjectDetail
		err = tx.Model(&system.ProjectDetail{}).Where("project_id = ?", projectId).Find(&projectDetailList).Error
		if err != nil {
			return err
		}
		for _, projectDetail := range projectDetailList {
			err = tx.Delete(&system.ProjectDetail{}, "id = ?", projectDetail.Id).Error
			if err != nil {
				return err
			}
			err = tx.Where("project_detail_id = ?", projectDetail.Id).Delete(&system.VideoConfig{}).Error
			if err != nil {
				return err
			}
			err = tx.Where("project_detail_id = ?", projectDetail.Id).Delete(&system.ParticipleConfig{}).Error
			if err != nil {
				return err
			}
			err = tx.Where("project_detail_id = ?", projectDetail.Id).Delete(&system.AudioConfig{}).Error
			if err != nil {
				return err
			}
			err = tx.Where("project_detail_id = ?", projectDetail.Id).Delete(&system.Info{}).Error
			if err != nil {
				return err
			}
			err = tx.Where("project_detail_id = ?", projectDetail.Id).Delete(&system.StableDiffusionImages{}).Error
			if err != nil {
				return err
			}
			err = tx.Where("project_detail_id = ?", projectDetail.Id).Delete(&system.StableDiffusionSettings{}).Error
			if err != nil {
				return err
			}
			err = tx.Where("project_detail_id = ?", projectDetail.Id).Delete(&system.StableDiffusionOverrideSettings{}).Error
			if err != nil {
				return err
			}
		}
		return err
	})
}

// GetProjectList 获取项目列表
func (s *ProjectService) GetProjectList(project system.Project, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&system.Project{})
	var projectList []system.Project
	if project.Name != "" {
		db = db.Where("name LIKE ?", "%"+project.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return projectList, total, err
	}
	db = db.Limit(limit).Offset(offset)
	OrderStr := "id desc"
	//if order != "" {
	//	orderMap := make(map[string]bool, 5)
	//	orderMap["id"] = true
	//	orderMap["path"] = true
	//	orderMap["api_group"] = true
	//	orderMap["description"] = true
	//	orderMap["method"] = true
	//	if !orderMap[order] {
	//		err = fmt.Errorf("非法的排序字段: %v", order)
	//		return apiList, total, err
	//	}
	//	OrderStr = order
	//	if desc {
	//		OrderStr = order + " desc"
	//	}
	//}
	err = db.Preload("List").Order(OrderStr).Find(&projectList).Error
	return projectList, total, err
}
