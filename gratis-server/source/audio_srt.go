package source

import (
	"fmt"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/python_core"
	"github/stable-diffusion-go/server/utils"
	"os"
	"path"
	"runtime"
	"strconv"
)

type AudioAndSrtParams struct {
	system.AudioConfig
	SavePath   string //保存路径
	Name       string //名称
	Content    string //内容
	Language   string // 语言
	BreakAudio bool   // 是否跳过
}

func windowCmdArgsConversion(value string) string {
	if runtime.GOOS == "windows" {
		return value + "%"
	}
	return value
}

func CreateAudioAndSrt(config AudioAndSrtParams) error {
	tmpFile, err := os.CreateTemp(".", "participle-*.py")
	if err != nil {
		fmt.Println("创建python文件失败:", err)
		return err
	}
	voice := config.Voice
	rate := config.Rate
	volume := config.Volume
	pitch := config.Pitch
	limit := config.SrtLimit
	audioPath := path.Join(config.SavePath, config.Name+".mp3")
	audioSrtPath := path.Join(config.SavePath, config.Name+".srt")
	audioSrtMapPath := path.Join(config.SavePath, config.Name+"map.txt")
	_, err = os.Stat(audioPath)
	// 如果跳过，并且文件存在
	if config.BreakAudio && err == nil {
		return nil
	}
	_, err = tmpFile.Write([]byte(python_core.PythonVoiceCaptionPath))
	if err != nil {
		fmt.Println("写入python内容失败", err)
		return err
	}
	//filepath := path.Join(global.Config.Local.StorePath, "participleBook.txt")
	args := []string{
		tmpFile.Name(),
		"--content", config.Content,
		"--participle_book_path", "",
		"--audi_srt_map_path", audioSrtMapPath,
		"--audio_path", audioPath,
		"--audio_srt_path", audioSrtPath,
		"--voice", voice, // 角色
		"--rate", windowCmdArgsConversion(rate), // 语速
		"--volume", windowCmdArgsConversion(volume), // 音量
		"--pitch", pitch, // 分贝
		"--pitch", pitch, // 分贝
		"--limit", strconv.Itoa(limit), // 分贝
		"--language", config.Language, // 语言
	}
	err = utils.ExecCommand("python", args)
	if err != nil {
		return err
	}
	fmt.Println("字幕和音频生成完成")
	return nil
}
