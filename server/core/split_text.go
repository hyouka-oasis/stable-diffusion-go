package core

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"sync"
)

func ProcessText(path string) (err error) {
	// 异步处理文本
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err = startCmd(path)
		if err != nil {
			return
		}
		defer wg.Done()
	}()
	wg.Wait()
	return nil
}

func startCmd(path string) error {
	participlePythonPath := global.ParticiplePythonPath
	outParticipleBookPathBookPath := global.OutParticipleBookPath
	maxWords := global.Config.Potential.MaxWords
	minWords := global.Config.Potential.MinWords
	args := []string{
		participlePythonPath,
		"--book_path", path,
		"--participle_book_path", outParticipleBookPathBookPath,
		"--max_word", maxWords,
		"--min_word", minWords,
	}
	err := ExecCommand("python", args)
	if err != nil {
		return err
	}
	fmt.Println("进行分词完成")
	return nil
}
