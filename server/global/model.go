package global

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	Id        uint           `gorm:"primarykey" json:"id" form:"id"` // 主键ID
	CreatedAt time.Time      `json:"createdAt"`                      // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`                      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                 // 删除时间
}
