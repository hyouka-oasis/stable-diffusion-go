package source

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"path"
	"runtime"
	"strconv"
)

func windowCmdArgsConversion(value string) string {
	if runtime.GOOS == "windows" {
		return value + "%"
	}
	return value
}

func CreateAudioAndSrt(savePath string, name string, projectDetail system.ProjectDetail) error {
	voiceCaptionPythonPath := global.VoiceCaptionPath
	voice := projectDetail.AudioConfig.Voice
	rate := projectDetail.AudioConfig.Rate
	volume := projectDetail.AudioConfig.Volume
	pitch := projectDetail.AudioConfig.Pitch
	limit := projectDetail.AudioConfig.SrtLimit
	language := projectDetail.Language
	audioPath := path.Join(savePath, name+".mp3")
	audioSrtPath := path.Join(savePath, name+".srt")
	audioSrtMapPath := path.Join(savePath, name+"map.txt")
	filepath := path.Join(global.Config.Local.StorePath, "participleBook.txt")
	args := []string{
		voiceCaptionPythonPath,
		"--participle_book_path", filepath,
		"--audi_srt_map_path", audioSrtMapPath,
		"--audio_path", audioPath,
		"--audio_srt_path", audioSrtPath,
		"--voice", voice, // 角色
		"--rate", windowCmdArgsConversion(rate), // 语速
		"--volume", windowCmdArgsConversion(volume), // 音量
		"--pitch", pitch, // 分贝
		"--pitch", pitch, // 分贝
		"--limit", strconv.Itoa(limit), // 分贝
		"--language", language, // 语言
	}
	fmt.Println(args)
	err := utils.ExecCommand("python", args)
	if err != nil {
		return err
	}
	fmt.Println("字幕和音频生成完成")
	return nil
}
