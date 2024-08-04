package system

type AudioConfig struct {
	ProjectDetailId uint   `json:"projectDetailId"`                                       // 项目详情Id
	InfoId          uint   `json:"infoId"`                                                // 单个列表ID
	Voice           string `json:"voice" form:"voice" gorm:"default:'zh-CN-YunxiNeural'"` // 角色名称
	Rate            string `json:"rate" form:"rate" gorm:"default:'+0%'"`                 // 语速
	Volume          string `json:"volume" form:"volume" gorm:"default:'+100%'"`           // 音量
	Pitch           string `json:"pitch" form:"pitch" gorm:"default:'+0Hz'"`              // 分贝
	SrtLimit        int    `json:"srtLimit" form:"srtLimit" gorm:"default:15"`            //字幕每一行最大的长度
}

func (AudioConfig) TableName() string {
	return "audio_config"
}
