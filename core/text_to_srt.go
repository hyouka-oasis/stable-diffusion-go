package core

import (
	"ComicTweetsGo/global"
	"bytes"
	"errors"
	"fmt"
	"github.com/wujunwei928/edge-tts-go/edge_tts"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

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
	bookPath := global.BookPath
	bookName := global.Config.Book.Name
	file, err := os.ReadFile(bookPath)
	if err != nil {
		log.Fatal("读取文件失败")
		return
	}
	maxAttempts := 10 // 设置最大尝试次数
	attempts := 0     // 初始化尝试次数计数器
	for attempts < maxAttempts {
		err = createVoiceSrt(bookName, string(file))
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

func createVoiceSrt(name string, fileContent string) (err error) {
	if len(fileContent) <= 0 {
		return errors.New("错误的内容")
	}
	voice := global.Config.Audio.Voice
	rate := global.Config.Audio.Rate
	volume := global.Config.Audio.Volume
	pitch := global.Config.Audio.Pitch
	proxyURL := global.Config.Audio.ProxyUrl
	bookMp3Path := global.BookMp3Path
	connOptions := []edge_tts.CommunicateOption{
		edge_tts.SetVoice(voice),
		edge_tts.SetRate(rate),
		edge_tts.SetVolume(volume),
		edge_tts.SetPitch(pitch),
		edge_tts.SetReceiveTimeout(20),
	}
	if len(proxyURL) > 0 {
		connOptions = append(connOptions, edge_tts.SetProxy(proxyURL))
	}

	conn, err := edge_tts.NewCommunicate(
		fileContent,
		connOptions...,
	)
	if err != nil {
		return err
	}
	audioData, err := conn.Stream()
	if err != nil {
		return err
	}
	if len(bookMp3Path) > 0 {
		writeMediaErr := os.WriteFile(bookMp3Path, audioData, 0644)
		if writeMediaErr != nil {
			return writeMediaErr
		}
		return nil
	}
	// write mp3 file's binary data to stdout
	_, err = io.Copy(os.Stdout, bytes.NewReader(audioData))
	if err != nil {
		return err
	}
	return nil
}
