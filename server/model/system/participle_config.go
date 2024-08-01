package system

type ParticipleConfig struct {
	ProjectDetailId uint `json:"projectDetailId"`                             // 项目详情Id
	MinLength       int  `json:"minLength" form:"minLength" gorm:"default:2"` // 名称最小长度
	MaxLength       int  `json:"maxLength" form:"maxLength" gorm:"default:4"` // 名称最大长度
	MinWords        int  `json:"minWords" form:"minWords" gorm:"default:30"`  // 每张图片的最小文字数量
	MaxWords        int  `json:"maxWords" form:"maxWords" gorm:"default:40"`  // 每张图片的最大文字数量
}

func (ParticipleConfig) TableName() string {
	return "participle_config"
}
