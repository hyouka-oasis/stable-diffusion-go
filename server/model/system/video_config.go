package system

type VideoConfig struct {
	ProjectDetailId uint   `json:"projectDetailId"`              // 项目详情Id
	InfoId          uint   `json:"infoId"`                       // 单个列表ID
	FontSize        int    `json:"fontSize" gorm:"default:12"`   // 角色名称
	FontColor       string `json:"rate" gorm:"default:'FFFFFF'"` // 语速
}

func (VideoConfig) TableName() string {
	return "video_config"
}
