package system

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
)

type ProjectDetailParticipleListService struct{}

// DeleteProjectDetailParticipleListItem 删除单条记录
func (s *ProjectDetailParticipleListService) DeleteProjectDetailParticipleListItem(id uint) error {
	err := global.DB.Delete(&system.ProjectDetailParticipleList{}, "id = ?", id).Error
	return err
}
