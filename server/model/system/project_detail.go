package system

import (
	"github/stable-diffusion-go/server/global"
)

type ProjectDetail struct {
	global.Model
	ProjectId             uint                    `json:"projectId" form:"projectId"` // 项目Id
	FileName              string                  `json:"fileName"`                   // 文件名称
	ParticipleConfig      ParticipleConfig        `json:"participleConfig" form:"participleConfig"`
	InfoList              []Info                  `json:"infoList"`
	StableDiffusionConfig StableDiffusionSettings `json:"stableDiffusionConfig"`             // api调用参数
	Language              string                  `json:"language" gorm:"default:zh"`        // 语言
	AudioConfig           AudioConfig             `json:"audioConfig"`                       // 音频配置
	VideoConfig           VideoConfig             `json:"videoConfig"`                       // 视频配置
	OpenSubtitles         bool                    `json:"openSubtitles" gorm:"default:true"` // 是否开启字幕
	BreakAudio            bool                    `json:"breakAudio" gorm:"default:true"`    // 是否跳过存在的音频
	BreakVideo            bool                    `json:"breakVideo" gorm:"default:true"`    // 是否跳过存在的视频
	ConcatAudio           bool                    `json:"concatAudio" gorm:"default:false"`  // 是否合并音频
	ConcatVideo           bool                    `json:"concatVideo" gorm:"default:false"`  // 是否合并视频
	OpenContext           bool                    `json:"openContext" gorm:"default:true"`   // 是否开启上下文模式
	PromptText            string                  `json:"promptText"`                        // 自定义训练prompt地址
}

func (ProjectDetail) TableName() string {
	return "project_detail"
}
