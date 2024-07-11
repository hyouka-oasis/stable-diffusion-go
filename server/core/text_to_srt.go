package core

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"log"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

func windowCmdArgsConversion(value string) string {
	if runtime.GOOS == "windows" {
		return value + "%"
	}
	return value
}

func TextToSrt() {
	// 异步处理文本
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		processTextToSrt()
	}()
	wg.Wait()
}

func processTextToSrt() {
	fmt.Println("开始生成字幕和音频")
	participleBookPath := global.OutParticipleBookPathBookPath
	maxAttempts := 10 // 设置最大尝试次数
	attempts := 0     // 初始化尝试次数计数器
	for attempts < maxAttempts {
		err := createVoiceSrt(participleBookPath)
		if err == nil {
			// 如果成功，则返回
			return
		}
		// 捕获到异常，打印错误信息，并决定是否重试
		fmt.Printf("尝试生成语音字幕时出错: %v\n", err)
		attempts++
		time.Sleep(10 * time.Second) // 等待一段时间后重试，避免立即重试
	}
	// 超过最大重试次数，返回错误
	log.Fatal("尝试生成语音字幕失败次数过多，停止重试")
}

func createVoiceSrt(filepath string) (err error) {
	voice := global.Config.Audio.Voice
	rate := global.Config.Audio.Rate
	volume := global.Config.Audio.Volume
	pitch := global.Config.Audio.Pitch
	audioPath := global.OutAudioPath
	audioSrtPath := global.OutAudioSrtPath
	voiceCaptionPythonPath := global.VoiceCaptionPath
	audioSrtMapPath := global.OutAudioSrtMapPath
	args := []string{
		voiceCaptionPythonPath,
		"--book_path", filepath,
		"--audi_srt_map_path", audioSrtMapPath,
		"--audio_path", audioPath,
		"--audio_srt_path", audioSrtPath,
		"--voice", voice, // 角色
		"--rate", windowCmdArgsConversion(rate), // 语速
		"--volume", windowCmdArgsConversion(volume), // 音量
		"--pitch", pitch, // 分贝
	}
	cmd := exec.Command("python", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("错误的执行代码", string(output))
		return err
	}
	fmt.Println("字幕和音频生成完成")
	return nil
}
