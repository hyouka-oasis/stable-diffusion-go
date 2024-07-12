package config

type CatchMergeConfig struct {
	AudioSrtMapPath    string // 输出的音频字幕map路径
	VideoCatchTxtPath  string // ffmpeg需要合成的视频列表txt
	VideoSubtitlesName string // 带有字幕的视频名称
}
