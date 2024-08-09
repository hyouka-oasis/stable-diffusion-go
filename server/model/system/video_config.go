package system

type VideoConfig struct {
	ProjectDetailId uint   `json:"projectDetailId"`                                    // 项目详情Id
	InfoId          uint   `json:"infoId"`                                             // 单个列表ID
	FontSize        int    `json:"fontSize" form:"fontSize" gorm:"default:12"`         // 字体大小
	FontColor       string `json:"fontColor" form:"fontColor" gorm:"default:'FFFFFF'"` // 字体颜色
	FontFile        string `json:"fontFile" form:"fontFile"`                           // 字体文件
	Position        int    `json:"position" form:"position" gorm:"default:2"`          // 字幕位置
	//StrokeColor     string  `json:"strokeColor"`                         // 描边颜色
	//StrokeWidth     string  `json:"strokeWidth"`                         // 描边宽度
	//Kerning         int     `json:"kerning"`                             // 文字间距
	AnimationSpeed float64 `json:"animationSpeed" form:"animationSpeed" gorm:"default:1.2"`  // 动画速度
	AnimationName  string  `json:"animationName" form:"animationName" gorm:"default:random"` // 动画名称默认随机
}

func (VideoConfig) TableName() string {
	return "video_config"
}
