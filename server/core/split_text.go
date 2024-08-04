package core

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/utils"
	"sync"
)

func SplitText() (err error) {
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
	bookPath := global.BookPath
	outParticipleBookPathBookPath := global.OutParticipleBookPath
	maxWords := global.Config.Participle.MaxWords
	minWords := global.Config.Participle.MinWords
	args := []string{
		participlePythonPath,
		"--book_path", bookPath,
		"--participle_book_path", outParticipleBookPathBookPath,
		"--max_word", maxWords,
		"--min_word", minWords,
		"--whether_participle", "no",
	}
	fmt.Println(args)
	err := utils.ExecCommand("python", args)
	if err != nil {
		return err
	}
	fmt.Println("进行分词完成")
	return nil
}
