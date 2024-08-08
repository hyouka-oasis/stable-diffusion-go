package source

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"strconv"
)

func SplitText(projectDetail system.ProjectDetail, whetherParticiple string, pythonName string) error {
	bookPath := global.Config.Local.Path + "/" + projectDetail.FileName
	outParticipleBookPathBookPath := global.Config.Local.Path + "/" + global.ParticipleBookName
	maxWords := projectDetail.ParticipleConfig.MaxWords
	minWords := projectDetail.ParticipleConfig.MinWords

	args := []string{
		pythonName,
		"--book_path", bookPath,
		"--participle_book_path", outParticipleBookPathBookPath,
		"--max_word", strconv.Itoa(maxWords),
		"--min_word", strconv.Itoa(minWords),
		"--whether_participle", whetherParticiple,
	}
	err := utils.ExecCommand("python", args)
	if err != nil {
		return err
	}
	fmt.Println("进行分词完成")
	return nil
}
