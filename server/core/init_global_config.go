package core

import (
	"encoding/json"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/utils"
	"os"
	"path/filepath"
)

// 兼容json文件中的宽度和高度
func initStableDiffusionJsonWidthAndHeight() {
	// 读取 JSON 文件内容
	jsonContent, err := os.ReadFile(global.Config.StableDiffusionConfig.ArgsJson)

	// 读取失败
	if err != nil {
		initStableDiffusionImageWidthAndHeight()
	} else {
		var jsonData map[string]interface{}
		err = json.Unmarshal(jsonContent, &jsonData)
		if err != nil {
			initStableDiffusionImageWidthAndHeight()
			return
		}
		// 合并 Stable Diffusion 配置参数
		for key, value := range jsonData {
			if key == "width" && global.Config.StableDiffusionConfig.Width == 0 {
				global.Config.StableDiffusionConfig.Width = utils.GetInterfaceToInt(value)
			}
			if key == "height" && global.Config.StableDiffusionConfig.Height == 0 {
				global.Config.StableDiffusionConfig.Height = utils.GetInterfaceToInt(value)
			}
		}
	}
}

// 给默认高度宽度
func initStableDiffusionImageWidthAndHeight() {
	// 校验当前宽是不是0
	if global.Config.StableDiffusionConfig.Width == 0 {
		global.Config.StableDiffusionConfig.Width = 512
	}
	if global.Config.StableDiffusionConfig.Height == 0 {
		global.Config.StableDiffusionConfig.Height = 512
	}
}

// 初始化输出路径的值
func initOutPath(bookName string, pwd string) {
	global.OutParticiplePath = filepath.Join(pwd, bookName+"/participle/")                  // 输出的路径
	global.OutImagesPath = filepath.Join(pwd, bookName+"/images/")                          // 图片输出路径
	global.OutVideoPath = filepath.Join(pwd, bookName+"/video/")                            // 图片输出路径
	global.OutParticipleBookPath = filepath.Join(global.OutParticiplePath, bookName+".txt") // 输出的文本路径
	global.OutBookJsonPath = filepath.Join(global.OutParticiplePath, bookName+".json")      // 转换的prompt路径
	global.OutAudioPath = filepath.Join(global.OutParticiplePath, bookName+".mp3")          // 输出的mp3路径
	global.OutAudioSrtPath = filepath.Join(global.OutParticiplePath, bookName+".srt")       // 输出的字幕路径

	// 视频配置
	global.OutVideoName = filepath.Join(global.OutVideoPath, bookName+".mp4") // 最终输出的视频文件
}

// 初始化python脚本目录
func initPythonPath(pwd string) {
	// python脚本
	global.VoiceCaptionPath = filepath.Join(pwd, "python_core/voice_caption.py")  // 进行字幕转换的python代码
	global.ParticiplePythonPath = filepath.Join(pwd, "python_core/participle.py") // 进行文本分词的python代码
}

func initDeleteCatchPath(bookName string, pwd string) {
	// 需要删除的文件都存放在根目录
	global.CatchMergeConfig.AudioSrtMapPath = filepath.Join(pwd, bookName+"/"+bookName+"map.txt")          // 输出的字幕时间列表
	global.CatchMergeConfig.VideoCatchTxtPath = filepath.Join(pwd, bookName+"/"+bookName+"video_map.txt")  // 输出的字幕时间列表
	global.CatchMergeConfig.VideoSubtitlesName = filepath.Join(pwd, bookName+"/"+bookName+"subtitles.mp4") // 带字幕的视频
}

// InitGlobalConfig 初始化值
func InitGlobalConfig(bookName string) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	global.BookPath = filepath.Join(pwd, bookName+".txt") // 源文件地址
	// 初始化输出路径
	initOutPath(bookName, pwd)
	// 初始化最后需要删除的目录
	initDeleteCatchPath(bookName, pwd)
	// 初始化python目录
	initPythonPath(pwd)
	if global.Config.StableDiffusionConfig.ArgsJson != "" {
		initStableDiffusionJsonWidthAndHeight()
	} else {
		global.Config.StableDiffusionConfig.ArgsJson = "stable_diffusion.json"
	}
}
