package global

import (
	"github/stable-diffusion-go/server/config"
)

var (
	Config                        config.Config // yaml配置文件
	OutPath                       string        // 输出的目录
	OutImagesPath                 string        // 输出图片的路径
	OutVideoPath                  string        // 输出视频的路径
	OriginBookPath                string        // 源文件路径
	OutBookJsonPath               string        // 输出的prompt路径
	OutParticipleBookPathBookPath string        // 通过分割后的文本路径
	OutAudioPath                  string        // 输出的MP3路径
	OutAudioSrtPath               string        // 输出的字幕路径
	OutAudioSrtMapPath            string        // 输出的音频字幕map路径
	VoiceCaptionPath              string        // 执行生成srt的Python路径
	ParticiplePythonPath          string        // 执行生成分词的Python路径
)

// Animations 动画列表
var Animations = []string{"shrink", "magnify", "left_move", "right_move", "up_move", "down_move"}
