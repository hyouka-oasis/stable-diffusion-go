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
		if project != (system.Project{}) {
			return errors.New("存在相同名称项目")
		}
		err = tx.Create(&config).Error
		if err != nil {
			return err
		}
		projectDetail := system.ProjectDetail{
			ProjectId: config.Id,
		}
		// 同时创建项目详情
		err = tx.Create(&projectDetail).Error
		if err != nil {
			return err
		}
		participleConfig := system.ParticipleConfig{
			ProjectDetailId: projectDetail.Id,
		}
		// 同时创建项目音频设置
		err = tx.Create(&participleConfig).Error
		if err != nil {
			return err
		}
		audioConfig := system.AudioConfig{
			ProjectDetailId: projectDetail.Id,
		}
		// 同时创建项目详情分词
		err = tx.Create(&audioConfig).Error
		if err != nil {
			return err
		}
		return err
	})
}

// DeleteProject 删除项目
func (s *ProjectService) DeleteProject(config system.Project) (err error) {
	var entity system.Project
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.First(&entity, "id = ?", config.Id).Error // 根据id查询api记录
		if errors.Is(err, gorm.ErrRecordNotFound) {        // 记录不存在
			return err
		}
		err = tx.Delete(&entity).Error
		if err != nil {
			return err
		}
		var projectDetail system.ProjectDetail
		err = tx.Model(&system.ProjectDetail{}).Where("project_id = ?", config.Id).First(&projectDetail).Error
		if err != nil {
			return err
		}
		err = tx.Where("project_id = ?", config.Id).Delete(&system.ProjectDetail{}).Error
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
		// 返回 nil 提交事务
		return nil
	})
	return err
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
	err = db.Order(OrderStr).Find(&projectList).Error
	return projectList, total, err
}
