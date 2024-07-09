package core

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

func windowCmdArgsConversion(value string) string {
	if os.Getenv("OSTYPE") == "windows" {
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
	filepath := global.OutBookPath
	maxAttempts := 10 // 设置最大尝试次数
	attempts := 0     // 初始化尝试次数计数器
	for attempts < maxAttempts {
		err := createVoiceSrt(filepath)
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
	bookMp3Path := global.OutBookMp3Path
	bookMp3SrtPath := global.OutBookMp3SrtPath
	//bookPath := global.OutBookPath
	voiceCaptionPythonPath := global.VoiceCaptionPath
	args := []string{
		voiceCaptionPythonPath,
		"--text_path", filepath,
		"--mp3_path", bookMp3Path,
		"--srt_path", bookMp3SrtPath,
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
	fmt.Println("代码执行完成")
	return nil
	//if len(proxyURL) > 0 {
	//	connOptions = append(connOptions, edge_tts.SetProxy(proxyURL))
	//}
	//
	//conn, err := edge_tts.NewCommunicate(
	//	fileContent,
	//	connOptions...,
	//)
	//if err != nil {
	//	return err
	//}
	//audioData, err := conn.Stream()
	//if err != nil {
	//	return err
	//}
	//if len(bookMp3Path) > 0 {
	//	writeMediaErr := os.WriteFile(bookMp3Path, audioData, 0644)
	//	if writeMediaErr != nil {
	//		return writeMediaErr
	//	}
	//	return nil
	//}
	//// write mp3 file's binary data to stdout
	//_, err = io.Copy(os.Stdout, bytes.NewReader(audioData))
	//if err != nil {
	//	return err
	//}
	return nil
}
