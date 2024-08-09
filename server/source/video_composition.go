package source

import (
	"fmt"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/example"
	"github/stable-diffusion-go/server/model/system"
	"github/stable-diffusion-go/server/utils"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type DisposableSynthesisVideoParams struct {
	system.Info
	system.VideoConfig
	example.ExaFileUploadAndDownload
	SavePath          string
	VideoPath         string
	VideoSubtitlePath string
	Width             int
	Height            int
	BreakVideo        bool // 是否跳过存在的音频
	OpenSubtitles     bool // 是否开启字幕
}

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

// 创建单个视频
func createAnimatedSegment(params DisposableSynthesisVideoParams, duration float64, audioPath string) error {
	animationSpeed := params.AnimationSpeed
	initialZoom := 1.0
	imageWidth := float64(params.Width)
	imageHeight := float64(params.Height)
	offsetTime := float64(25)
	zoomSteps := (animationSpeed - initialZoom) / (offsetTime * duration)
	leftRightMove := (imageWidth*animationSpeed - imageWidth - offsetTime) / (offsetTime * duration)
	upDownMove := (imageHeight*animationSpeed - imageHeight - offsetTime - offsetTime) / (offsetTime * duration)
	ffmpegWidthAndHeight := strconv.Itoa(int(imageWidth)) + "x" + strconv.Itoa(int(imageHeight))
	var scale string
	if params.AnimationName == "shrink" {
		scale = fmt.Sprintf("scale=-2:ih*10,zoompan=z='if(lte(zoom,%f),%f,max(zoom-%.19f,1))':x='iw/2-(iw/zoom/2)':y='ih/2-(ih/zoom/2)':d=%f*%f:s=%s",
			initialZoom, animationSpeed, zoomSteps, duration, offsetTime, ffmpegWidthAndHeight)
	} else if params.AnimationName == "left_move" {
		scale = fmt.Sprintf("scale=-2:ih*10,zoompan='%f':x='if(lte(on,-1),(iw-iw/zoom)/2,x+%.13f)':y='if(lte(on,1),(ih-ih/zoom)/2,y)':d=%f*%f:s=%s",
			animationSpeed, leftRightMove*10, duration, offsetTime, ffmpegWidthAndHeight)
	} else if params.AnimationName == "right_move" {
		scale = fmt.Sprintf("scale=-2:ih*10,zoompan='%f':x='if(lte(on,1),(iw/zoom)/2,x-%.13f)':y='if(lte(on,1),(ih-ih/zoom)/2,y)':d=%f*%f:s=%s",
			animationSpeed, leftRightMove*10, duration, offsetTime, ffmpegWidthAndHeight)
	} else if params.AnimationName == "up_move" {
		scale = fmt.Sprintf("scale=-2:ih*10,zoompan='%f':x='if(lte(on,1),(iw-iw/zoom)/2,x)':y='if(lte(on,-1),(ih-ih/zoom)/2,y+%.13f)':d=%f*%f:s=%s",
			animationSpeed, upDownMove*10, duration, offsetTime, ffmpegWidthAndHeight)
	} else if params.AnimationName == "down_move" {
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
		params.Url,
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
		params.VideoPath,
	}
	err := utils.ExecCommand("ffmpeg", args)
	if err != nil {
		return err
	}
	return nil
}

// 创建字幕视频
func createSubtitleVideo(params DisposableSynthesisVideoParams, srtPath string) error {
	fontFile := params.FontFile
	fontSize := strconv.Itoa(params.FontSize)
	var fontColor string
	if global.Config.Video.FontColor != "" {
		fontColor = global.Config.Video.FontColor
	} else {
		fontColor = strings.Replace(params.FontColor, "#", "", -1)
	}
	fontPosition := strconv.Itoa(params.Position)
	audioSrtPath := srtPath
	//+ "StrokeColor=&HFFFFFF" + "StrokeWidth=50" + "Kerning=10"
	subtitleStyle := "Fontname=" + fontFile + "," + "Fontsize=" + fontSize + "," + "PrimaryColour=&H" + fontColor + "," + "Alignment=" + fontPosition + ","
	subtitleVideoName, err := windowCmdArgsConversionPath(params.VideoSubtitlePath)
	if err != nil {
		fmt.Println("Error getting relative out path:", err)
		return err
	}
	videoPath, err := windowCmdArgsConversionPath(params.VideoPath)
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

// DisposableSynthesisVideo 生成视频
func DisposableSynthesisVideo(params DisposableSynthesisVideoParams) (err error) {
	_, err = os.Stat(params.Url)
	if err != nil {
		return err
	}
	videoPath := params.VideoPath
	videoSubtitlePath := params.VideoSubtitlePath
	_, videoPathError := os.Stat(videoPath)
	_, videoSubtitlePathError := os.Stat(videoSubtitlePath)
	// 如果跳过存在视频
	if params.BreakVideo {
		// 如果开启字幕
		if params.OpenSubtitles {
			if videoPathError == nil && videoSubtitlePathError == nil {
				return nil
			}
		} else {
			if videoPathError == nil {
				return nil
			}
		}
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
	if params.AnimationName == "random" {
		params.AnimationName = global.Animations[rand.Intn(len(global.Animations))]
	}
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
	err = createAnimatedSegment(params, duration, audioPath)
	if err != nil {
		return err
	}
	if params.OpenSubtitles {
		err = createSubtitleVideo(params, srtPath)
		if err != nil {
			return err
		}
	}
	return err
}

// MergeVideoList 合并带字幕和不带字幕的视频
func MergeVideoList(videoList []string, outVideoPath string) error {
	if len(videoList) == 0 {
		return nil
	}
	// 构建 ffmpeg 命令
	var inputArgs []string
	for _, path := range videoList {
		inputArgs = append(inputArgs, "-i", path)
	}
	filterComplexArgs := make([]string, 0, len(videoList)+1)
	for i := 0; i < len(videoList); i++ {
		filterComplexArgs = append(filterComplexArgs, fmt.Sprintf("[%d:v]", i), fmt.Sprintf("[%d:a]", i))
	}
	filterComplexArgs = append(filterComplexArgs, fmt.Sprintf("concat=n=%d:v=1:a=1[outVideo][outAudio]", len(videoList)))

	outputArgs := []string{
		"-filter_complex",
		strings.Join(filterComplexArgs, ""),
		"-map",
		"[outVideo]",
		"-map",
		"[outAudio]",
		"-c:v",
		"libx264",
		"-crf",
		"23",
		"-y",
		outVideoPath,
	}
	args := append(inputArgs, outputArgs...)
	err := utils.ExecCommand("ffmpeg", args)
	if err != nil {
		return err
	}
	return err
}
