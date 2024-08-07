package source

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/python_core"
	"github/stable-diffusion-go/server/utils"
	"os"
	"strconv"
	"sync"
)

func SplitText(projectDetail system.ProjectDetail, whetherParticiple string) (err error) {
	// 异步处理文本
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err = startCmd(projectDetail, whetherParticiple)
		if err != nil {
			return
		}
		defer wg.Done()
	}()
	wg.Wait()
	return nil
}

func startCmd(projectDetail system.ProjectDetail, whetherParticiple string) error {
	tmpFile, err := os.CreateTemp(".", "participle-*.py")
	if err != nil {
		fmt.Println("创建python文件失败:", err)
		return err
	}
	bookPath := global.Config.Local.Path + "/" + projectDetail.FileName
	outParticipleBookPathBookPath := global.Config.Local.Path + "/" + global.ParticipleBookName
	maxWords := projectDetail.ParticipleConfig.MaxWords
	minWords := projectDetail.ParticipleConfig.MinWords
	_, err = tmpFile.Write([]byte(python_core.PythonParticiplePythonPath))
	if err != nil {
		fmt.Println("写入python内容失败", err)
		return err
	}
	args := []string{
		tmpFile.Name(),
		"--book_path", bookPath,
		"--participle_book_path", outParticipleBookPathBookPath,
		"--max_word", strconv.Itoa(maxWords),
		"--min_word", strconv.Itoa(minWords),
		"--whether_participle", whetherParticiple,
	}
	fmt.Println(args, tmpFile.Name())
	err = utils.ExecCommand("python", args)
	tmpFile.Close()
	os.Remove(tmpFile.Name())
	if err != nil {
		return err
	}
	fmt.Println("进行分词完成")
	return nil
}
