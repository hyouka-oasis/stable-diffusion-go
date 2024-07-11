package core

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"log"
	"os/exec"
	"sync"
)

func ProcessText() (err error) {
	// 异步处理文本
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err = startCmd()
		if err != nil {
			return
		}
		defer wg.Done()
	}()
	wg.Wait()
	return nil
}

func startCmd() error {
	participlePythonPath := global.ParticiplePythonPath
	bookPath := global.OriginBookPath
	outParticipleBookPathBookPath := global.OutParticipleBookPathBookPath
	maxWords := global.Config.Potential.MaxWords
	minWords := global.Config.Potential.MinWords
	args := []string{
		participlePythonPath,
		"--book_path", bookPath,
		"--participle_book_path", outParticipleBookPathBookPath,
		"--max_word", maxWords,
		"--min_word", minWords,
	}
	cmd := exec.Command("python", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln("错误的执行代码", string(output))
		return err
	}
	fmt.Println("进行分词完成")
	return nil
}
