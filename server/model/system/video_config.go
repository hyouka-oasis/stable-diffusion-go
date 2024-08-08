package system

type VideoConfig struct {
	ProjectDetailId uint    `json:"projectDetailId"`              // 项目详情Id
	InfoId          uint    `json:"infoId"`                       // 单个列表ID
	FontSize        int     `json:"fontSize" gorm:"default:12"`   // 字体大小
	FontColor       string  `json:"rate" gorm:"default:'FFFFFF'"` // 字体颜色
	FontFile        string  `json:"fontFile"`                     // 字体文件
	StrokeColor     string  `json:"strokeColor"`                  // 描边颜色
	StrokeWidth     string  `json:"strokeWidth"`                  // 描边宽度
	Kerning         int     `json:"kerning"`                      // 文字间距
	AnimationSpeed  float64 `json:"animationSpeed"`               // 动画速度
}

func (VideoConfig) TableName() string {
	return "video_config"
}
