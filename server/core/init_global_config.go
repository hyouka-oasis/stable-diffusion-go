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
	global.OutPath = filepath.Join(wd, bookName+"/participle/")                   //输出的路径
	global.OutImagesPath = filepath.Join(wd, bookName+"/images/")                 // 图片输出路径
	global.OriginBookPath = filepath.Join(wd, bookName+".txt")                    //原路径
	global.VoiceCaptionPath = filepath.Join(wd, "python_core/voice_caption.py")   //python代码
	global.OutBookPath = filepath.Join(global.OutPath, bookName+".txt")           //输出的文本路径
	global.OutBookJsonPath = filepath.Join(global.OutPath, bookName+".json")      // 转换的prompt路径
	global.OutAudioPath = filepath.Join(global.OutPath, bookName+".mp3")          //输出的mp3路径
	global.OutAudioSrtPath = filepath.Join(global.OutPath, bookName+".srt")       //输出的字幕路径
	global.OutAudioSrtMapPath = filepath.Join(global.OutPath, bookName+"map.txt") //输出的字幕时间列表
}
