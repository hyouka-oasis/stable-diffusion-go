package core

import (
	"github/stable-diffusion-go/server/global"
	"os"
	"path/filepath"
)

func InitGlobalConfig() {
	bookName := global.Config.Book.Name
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	global.BookPath = filepath.Join(wd, bookName+".txt")                                    // 源文件地址
	global.OutParticiplePath = filepath.Join(wd, bookName+"/participle/")                   // 输出的路径
	global.OutImagesPath = filepath.Join(wd, bookName+"/images/")                           // 图片输出路径
	global.OutVideoPath = filepath.Join(wd, bookName+"/video/")                             // 图片输出路径
	global.OutParticipleBookPath = filepath.Join(global.OutParticiplePath, bookName+".txt") // 输出的文本路径
	global.OutBookJsonPath = filepath.Join(global.OutParticiplePath, bookName+".json")      // 转换的prompt路径
	global.OutAudioPath = filepath.Join(global.OutParticiplePath, bookName+".mp3")          // 输出的mp3路径
	global.OutAudioSrtPath = filepath.Join(global.OutParticiplePath, bookName+".srt")       // 输出的字幕路径

	// 视频配置
	global.OutVideoName = filepath.Join(global.OutVideoPath, bookName+".mp4") // 最终输出的视频文件

	// 需要删除的文件都存放在根目录
	global.CatchMergeConfig.AudioSrtMapPath = filepath.Join(wd, bookName+"/"+bookName+"map.txt")          // 输出的字幕时间列表
	global.CatchMergeConfig.VideoCatchTxtPath = filepath.Join(wd, bookName+"/"+bookName+"video_map.txt")  // 输出的字幕时间列表
	global.CatchMergeConfig.VideoSubtitlesName = filepath.Join(wd, bookName+"/"+bookName+"subtitles.mp4") // 带字幕的视频

	// python脚本
	global.VoiceCaptionPath = filepath.Join(wd, "python_core/voice_caption.py")  // 进行字幕转换的python代码
	global.ParticiplePythonPath = filepath.Join(wd, "python_core/participle.py") // 进行文本分词的python代码

}
