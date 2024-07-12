package global

import (
	"github/stable-diffusion-go/server/config"
)

var (
	Config                config.Config           // yaml配置文件
	BookPath              string                  // 源文件路径
	OutParticiplePath     string                  // 分词的路径
	OutImagesPath         string                  // 输出图片的路径
	OutBookJsonPath       string                  // 输出的prompt路径
	OutParticipleBookPath string                  // 通过分割后的文本路径
	CatchMergeConfig      config.CatchMergeConfig // 缓存下来并且需要删除的文件配置
)

// 音频字幕类型
var (
	OutAudioPath    string // 输出的MP3路径
	OutAudioSrtPath string // 输出的字幕路径
)

// 视频类型
var (
	OutVideoPath string // 输出视频的路径
	OutVideoName string // 最终合成的视频
)

// 脚本类型
var (
	VoiceCaptionPath     string // 执行生成srt的Python路径
	ParticiplePythonPath string // 执行生成分词的Python路径
)

// Animations 动画列表
var Animations = []string{"shrink", "magnify", "left_move", "right_move", "up_move", "down_move"}
