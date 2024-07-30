package source

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"strconv"
	"sync"
)

func SplitText(projectDetail system.ProjectDetail) (err error) {
	// 异步处理文本
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err = startCmd(projectDetail)
		if err != nil {
			return
		}
		defer wg.Done()
	}()
	wg.Wait()
	return nil
}

func startCmd(projectDetail system.ProjectDetail) error {
	participlePythonPath := global.ParticiplePythonPath
	bookPath := global.Config.Local.Path + "/" + projectDetail.FileName
	outParticipleBookPathBookPath := global.Config.Local.Path + "/" + "participleBook.txt"
	maxWords := projectDetail.ParticipleConfig.MaxWords
	minWords := projectDetail.ParticipleConfig.MinWords
	//participlePythonPath := global.ParticiplePythonPath
	//bookPath := global.BookPath
	//outParticipleBookPathBookPath := global.OutParticipleBookPath
	//maxWords := global.Config.ParticipleConfig.MaxWords
	//minWords := global.Config.ParticipleConfig.MinWords
	args := []string{
		participlePythonPath,
		"--book_path", bookPath,
		"--participle_book_path", outParticipleBookPathBookPath,
		"--max_word", strconv.Itoa(maxWords),
		//"--max_word", maxWords,
		"--min_word", strconv.Itoa(minWords),
		//"--min_word", minWords,
	}
	fmt.Println(args)
	err := utils.ExecCommand("python", args)
	if err != nil {
		return err
	}
	fmt.Println("进行分词完成")
	return nil
}
