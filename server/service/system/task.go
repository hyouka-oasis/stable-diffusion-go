package system

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	systemRequest "github/stable-diffusion-go/server/model/system/request"
)

type TaskService struct{}

// CreateTask 创建任务
func (s *TaskService) CreateTask(params system.Task) (task *system.Task, err error) {
	err = global.DB.Create(&params).Error
	if err != nil {
		return &params, err
	}
	if len(params.Errors) != 0 {
		err = global.DB.Create(&params.Errors).Error
		if err != nil {
			return &params, err
		}
	}
	return &params, nil
}

// UpdateTask 更新任务
func (s *TaskService) UpdateTask(params system.Task) error {
	db := global.DB.Model(&system.Task{})
	if params.Id != 0 {
		db = db.Where("id = ?", params.Id)
	} else if params.ProjectDetailId != 0 {
		db = db.Where("project_detail_id = ?", params.ProjectDetailId)
	}
	err := db.Updates(&params).Error
	if err != nil {
		return err
	}
	if len(params.Errors) != 0 {
		err = global.DB.Model(&system.TaskErrors{}).Where("task_id = ?", params.Id).Updates(&params.Errors).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// GetTaskList 获取任务列表
func (s *TaskService) GetTaskList(params systemRequest.TaskPageInfoRequest) (list interface{}, total int64, err error) {
	limit := params.PageSize
	offset := params.PageSize * (params.Page - 1)
	var taskList []*system.Task
	db := global.DB.Preload("Errors").Model(&system.Task{})
	db = db.Limit(limit).Offset(offset)
	OrderStr := "id desc"
	if params.TaskId != 0 {
		err = db.Order(OrderStr).Where("task_id = ?", params.TaskId).Find(&taskList).Error
	} else if params.ProjectDetailId != 0 {
		err = db.Order(OrderStr).Where("project_detail_id = ?", params.ProjectDetailId).Find(&taskList).Error
	} else {
		err = db.Order(OrderStr).Find(&taskList).Error
	}
	err = db.Count(&total).Error
	if err != nil {
		return taskList, total, err
	}
	return taskList, total, err
}

// GetTask 获取任务详情
func (s *TaskService) GetTask(params systemRequest.TaskPageInfoRequest) (task *system.Task, err error) {
	db := global.DB.Preload("Errors").Model(&system.Task{})
	if params.TaskId != 0 {
		err = db.Where("task_id = ?", params.TaskId).Find(&task).Error
	} else {
		err = db.Where("project_detail_id = ?", params.ProjectDetailId).Find(&task).Error
	}
	if err != nil {
		return task, err
	}
	return task, nil
}

// DeleteTaskWhereProjectDetailId 获取任务详情
func (s *TaskService) DeleteTaskWhereProjectDetailId(projectDetailId uint) error {
	err := global.DB.Where("project_detail_id = ?", projectDetailId).Delete(&system.Task{}).Error
	if err != nil {
		return err
	}
	return nil
}
