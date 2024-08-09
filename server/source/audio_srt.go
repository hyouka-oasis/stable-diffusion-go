package source

import (
	"fmt"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

type AudioAndSrtParams struct {
	system.AudioConfig
	SavePath   string // 保存路径
	AudioPath  string // 音频路径
	Name       string // 名称
	Content    string // 内容
	Language   string // 语言
	BreakAudio bool   // 是否跳过
}

func windowCmdArgsConversion(value string) string {
	if runtime.GOOS == "windows" {
		return value + "%"
	}
	return value
}

// CreateAudioAndSrt 创建音频和字幕
func CreateAudioAndSrt(config AudioAndSrtParams, pythonName string) error {
	voice := config.Voice
	rate := config.Rate
	volume := config.Volume
	pitch := config.Pitch
	limit := config.SrtLimit
	audioPath := config.AudioPath
	audioSrtPath := path.Join(config.SavePath, config.Name+".srt")
	audioSrtMapPath := path.Join(config.SavePath, config.Name+"map.txt")
	_, err := os.Stat(audioPath)
	// 如果跳过，并且文件存在
	if config.BreakAudio && err == nil {
		return nil
	}
	args := []string{
		pythonName,
		"--content", config.Content,
		"--audi_srt_map_path", audioSrtMapPath,
		"--audio_path", audioPath,
		"--audio_srt_path", audioSrtPath,
		"--voice", voice, // 角色
		"--rate", rate, // 语速
		"--volume", volume, // 音量
		"--pitch", pitch, // 分贝
		"--pitch", pitch, // 分贝
		"--limit", strconv.Itoa(limit), // 分贝
		"--language", config.Language, // 语言
	}
	fmt.Println("输出对象", args)
	err = utils.ExecCommand("python", args)
	if err != nil {
		return nil
	}
	fmt.Println("字幕和音频生成完成")
	return nil
}

// MergeAudio 合并音频
func MergeAudio(mergeAudioList []string, outAudioPath string) error {
	if len(mergeAudioList) == 0 {
		return nil
	}
	// 构建 ffmpeg 命令
	var inputArgs []string
	for _, file := range mergeAudioList {
		inputArgs = append(inputArgs, "-i", file)
	}

	filterComplexArgs := make([]string, 0, len(mergeAudioList)+1)
	for i := 0; i < len(mergeAudioList); i++ {
		filterComplexArgs = append(filterComplexArgs, fmt.Sprintf("[%d:a]", i))
	}
	filterComplexArgs = append(filterComplexArgs, fmt.Sprintf("concat=n=%d:v=0:a=1[outAudio]", len(mergeAudioList)))

	outputArgs := []string{
		"-filter_complex",
		strings.Join(filterComplexArgs, ""),
		"-map",
		"[outAudio]",
		"-y",
		outAudioPath,
	}
	args := append(inputArgs, outputArgs...)
	err := utils.ExecCommand("ffmpeg", args)
	if err != nil {
		return err
	}
	fmt.Println("字幕和音频生成完成")
	return nil
}
