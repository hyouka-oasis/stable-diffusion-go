package source

import (
	"bufio"
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/example"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func windowCmdArgsConversionPath(path string) (string, error) {
	var absPath string
	if runtime.GOOS == "windows" {
		projPath, err := filepath.Abs("./")
		if err != nil {
			fmt.Println("获取项目绝对路径失败:", err)
			return absPath, err
		}
		absPath, err = filepath.Rel(projPath, path)
		if err != nil {
			fmt.Println("转换相对路径失败:", err)
			return absPath, err
		}
		absPath = filepath.ToSlash(absPath)
		return absPath, err
	}
	return path, nil
}

// 获取字幕切片
func getAudioSrtMap() (audioSrtMap []string, err error) {
	// 指定的srt_map路径
	txtFilePath := global.CatchMergeConfig.AudioSrtMapPath

	file, err := os.Open(txtFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// 创建 scanner 读取文件内容
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// 去除方括号
		line = strings.TrimPrefix(line, "[")
		line = strings.TrimSuffix(line, "]")
		// 使用 strconv.Unquote 解析字符串
		parts := strings.Split(line, ",")
		for _, part := range parts {
			audioSrtMap = append(audioSrtMap, strings.TrimSpace(part))
		}
		if scannerError := scanner.Err(); scannerError != nil {
			fmt.Println("解析字符串失败:", scannerError)
			return
		}
	}

	// 检查是否有错误发生
	if scannerErr := scanner.Err(); err != nil {
		fmt.Println("读取文件内容失败:", scannerErr)
		return
	}

	return audioSrtMap, nil
}

// 获取图片列表
func getImagesMap() (picturePathList []string, err error) {
	imagesPath := global.OutImagesPath
	picturePathList, err = utils.GetPicturePaths(imagesPath, ".png")
	if err != nil {
		log.Fatalln("获取文件目录失败")
		return nil, err
	}
	// 对文件路径列表按照文件名中的数字大小进行排序
	sort.Slice(picturePathList, func(i, j int) bool {
		iNum, _ := utils.ExtractNumber(picturePathList[i], ".png")
		jNum, _ := utils.ExtractNumber(picturePathList[j], ".png")
		return iNum < jNum
	})
	return picturePathList, nil
}

// 转换字幕时间为float
func convertTimeToSeconds(timeStr string) (float64, error) {
	// Split the string by colon
	timeParts := strings.Split(timeStr, ":")

	// If milliseconds are included, split them by the comma
	var seconds, milliseconds float64
	if strings.Contains(timeParts[len(timeParts)-1], ".") {
		secAndMs := strings.Split(timeParts[len(timeParts)-1], ".")
		seconds, _ = strconv.ParseFloat(secAndMs[0], 64)
		milliseconds, _ = strconv.ParseFloat(secAndMs[1], 64)
	} else {
		seconds, _ = strconv.ParseFloat(timeParts[len(timeParts)-1], 64)
		milliseconds = 0
	}

	// Convert each part to an integer
	hours, _ := strconv.Atoi(timeParts[0])
	minutes, _ := strconv.Atoi(timeParts[1])

	// Calculate total seconds
	totalSeconds := float64(hours*3600+minutes*60) + seconds + milliseconds/1000

	return totalSeconds, nil
}

// 合并视频
func splicingVideo(catchVideoList []string) error {
	audioPath := global.OutAudioPath
	videoPath := global.OutVideoName
	// 打开一个文件用于写入
	file, err := os.Create(global.CatchMergeConfig.VideoCatchTxtPath)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return err
	}
	defer file.Close()

	for _, video := range catchVideoList {
		_, err := fmt.Fprintln(file, "file "+"'"+video+"'")
		if err != nil {
			fmt.Println("写入文件失败:", err)
			return err
		}
	}
	fmt.Println("开始合成视频")
	args := []string{
		"-y",
		"-f",
		"concat",
		"-safe",
		"0",
		"-i",
		global.CatchMergeConfig.VideoCatchTxtPath,
		"-i",
		audioPath,
		"-vsync",
		"cfr",
		"-pix_fmt",
		"yuv420p",
		videoPath,
	}
	err = utils.ExecCommand("ffmpeg", args)
	if err != nil {
		return err
	}
	return nil
}

// 创建单个视频
func createAnimatedSegment(imagePath string, duration float64, animation string, videoPath string, audioPath string, width int, height int) error {
	animationSpeed := 1.2 // TODO 需要做成配置
	initialZoom := 1.0
	imageWidth := float64(width)
	imageHeight := float64(height)
	offsetTime := float64(26)
	zoomSteps := (animationSpeed - initialZoom) / (offsetTime * duration)
	leftRightMove := (imageWidth*animationSpeed - imageWidth - offsetTime) / (offsetTime * duration)
	upDownMove := (imageHeight*animationSpeed - imageHeight - offsetTime - offsetTime) / (offsetTime * duration)
	ffmpegWidthAndHeight := strconv.Itoa(int(imageWidth)) + "x" + strconv.Itoa(int(imageHeight))
	var scale string
	if animation == "shrink" {
		scale = fmt.Sprintf("scale=-2:ih*10,zoompan=z='if(lte(zoom,%f),%f,max(zoom-%.19f,1))':x='iw/2-(iw/zoom/2)':y='ih/2-(ih/zoom/2)':d=%f*%f:s=%s",
			initialZoom, animationSpeed, zoomSteps, duration, offsetTime, ffmpegWidthAndHeight)
	} else if animation == "left_move" {
		scale = fmt.Sprintf("scale=-2:ih*10,zoompan='%f':x='if(lte(on,-1),(iw-iw/zoom)/2,x+%.13f)':y='if(lte(on,1),(ih-ih/zoom)/2,y)':d=%f*%f:s=%s",
			animationSpeed, leftRightMove*10, duration, offsetTime, ffmpegWidthAndHeight)
	} else if animation == "right_move" {
		scale = fmt.Sprintf("scale=-2:ih*10,zoompan='%f':x='if(lte(on,1),(iw/zoom)/2,x-%.13f)':y='if(lte(on,1),(ih-ih/zoom)/2,y)':d=%f*%f:s=%s",
			animationSpeed, leftRightMove*10, duration, offsetTime, ffmpegWidthAndHeight)
	} else if animation == "up_move" {
		scale = fmt.Sprintf("scale=-2:ih*10,zoompan='%f':x='if(lte(on,1),(iw-iw/zoom)/2,x)':y='if(lte(on,-1),(ih-ih/zoom)/2,y+%.13f)':d=%f*%f:s=%s",
			animationSpeed, upDownMove*10, duration, offsetTime, ffmpegWidthAndHeight)
	} else if animation == "down_move" {
		scale = fmt.Sprintf("scale=-2:ih*10,zoompan='%f':x='if(lte(on,1),(iw-iw/zoom)/2,x)':y='if(lte(on,1),(ih/zoom)/2,y-%.13f)':d=%f*%f:s=%s",
			animationSpeed, upDownMove*10, duration, offsetTime, ffmpegWidthAndHeight)
	} else {
		scale = fmt.Sprintf("scale=-2:ih*10,zoompan=z='min(zoom+%.19f,%f)*if(gte(zoom,1),1,0)+if(lt(zoom,1),1,0)':x='iw/2-(iw/zoom/2)':y='ih/2-(ih/zoom/2)':d=%f*%f:s=%s",
			zoomSteps, animationSpeed, duration, offsetTime, ffmpegWidthAndHeight)
	}
	args := []string{
		"-y",
		"-r",
		fmt.Sprintf("%f", offsetTime),
		"-loop",
		"1",
		"-t",
		fmt.Sprintf("%.3f", duration),
		"-i",
		imagePath,
		"-i",
		audioPath,
		"-filter_complex",
		scale,
		"-vframes",
		fmt.Sprintf("%d", int(offsetTime*duration)),
		"-c:v",
		"libx264",
		"-pix_fmt",
		"yuv420p",
		videoPath,
	}
	err := utils.ExecCommand("ffmpeg", args)
	if err != nil {
		return err
	}
	return nil
}

// 创建字幕视频
func createSubtitleVideo(srtPath string, videoPath string, subtitleVideoName string) error {
	fontName := strings.Split(global.Config.Video.FontFile, ".")[0]
	fontSize := global.Config.Video.FontSize
	fontColor := global.Config.Video.FontColor
	fontPosition := global.Config.Video.Position
	audioSrtPath := srtPath
	subtitleStyle := "FontName=" + fontName + "," + "Fontsize=" + fontSize + "," + "PrimaryColour=&H" + fontColor + "," + "Alignment=" + fontPosition + "WrapStyle=0"
	subtitleVideoName, err := windowCmdArgsConversionPath(subtitleVideoName)
	if err != nil {
		fmt.Println("Error getting relative out path:", err)
		return err
	}
	videoPath, err = windowCmdArgsConversionPath(videoPath)
	if err != nil {
		fmt.Println("Error getting relative video path:", err)
		return err
	}
	audioSrtPath, err = windowCmdArgsConversionPath(audioSrtPath)
	if err != nil {
		fmt.Println("Error getting relative SRT path:", err)
		return err
	}
	args := []string{
		"-i",
		videoPath,
		"-vf",
		"subtitles=" + audioSrtPath + ":force_style='" + subtitleStyle + "'",
		"-c:a",
		"copy",
		"-y",
		subtitleVideoName,
	}
	err = utils.ExecCommand("ffmpeg", args)
	if err != nil {
		return err
	}
	return nil
}

type DisposableSynthesisVideoParams struct {
	system.Info
	example.ExaFileUploadAndDownload
	SavePath string
	Width    int
	Height   int
}

// DisposableSynthesisVideo 生成视频
func DisposableSynthesisVideo(params DisposableSynthesisVideoParams) (err error) {
	_, err = os.Stat(params.Url)
	if err != nil {
		return err
	}
	srtMapPath := filepath.Join(params.SavePath, params.Name+"map.txt")
	_, err = os.Stat(srtMapPath)
	if err != nil {
		return err
	}
	srtPath := filepath.Join(params.SavePath, params.Name+".srt")
	_, err = os.Stat(srtPath)
	if err != nil {
		return err
	}
	audioPath := filepath.Join(params.SavePath, params.Name+".mp3")
	_, err = os.Stat(audioPath)
	if err != nil {
		return err
	}
	// TODO 随机选择一个动画效果 需要做成配置
	selectedAnimation := global.Animations[rand.Intn(len(global.Animations))]
	srtDuration, srtDurationError := utils.GetAudioSrtMap(srtMapPath)
	if srtDurationError != nil {
		return srtDurationError
	}
	srtDuration = strings.Replace(srtDuration, ",", ".", -1)
	srtDuration = strings.Replace(srtDuration, "'", "", -1)
	// 将时间字符串转换为时间结构体
	duration, durationErr := convertTimeToSeconds(srtDuration)
	if durationErr != nil {
		return durationErr
	}
	videoPath := filepath.Join(params.SavePath, params.Name+".mp4")
	err = createAnimatedSegment(params.Url, duration, selectedAnimation, videoPath, audioPath, params.Width, params.Height)
	if err != nil {
		return err
	}
	videoSubtitlePath := filepath.Join(params.SavePath, params.Name+"subtitle.mp4")
	err = createSubtitleVideo(srtPath, videoPath, videoSubtitlePath)
	if err != nil {
		return err
	}
	//err = splicingVideo(catchVideoList)
	//if err != nil {
	//	return err
	//}
	//subtitles := global.Config.Video.Subtitles
	//if subtitles {
	//	err = createSubtitleVideo()
	//	if err != nil {
	//		return err
	//	}
	//}
	//return err
	return nil
}
