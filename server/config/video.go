package config

type Video struct {
	ImagemagickPath string  `yaml:"imagemagickPath"` // 角色名称
	FontSize        string  `yaml:"fontSize"`        // 字幕字体大小
	FontColor       string  `yaml:"fontColor"`       // 字体颜色
	FontFile        string  `yaml:"fontFile"`        // 字体文件
	StrokeColor     string  `yaml:"strokeColor"`     // 描边颜色
	StrokeWidth     string  `yaml:"strokeWidth"`     // 描边宽
	Kerning         string  `yaml:"kerning"`         // 文字间距
	Position        string  `yaml:"position"`        // 字幕位置 越小越靠上 越大越靠下 0 - 1 开启ffmpeg后 6=上 10中 2下
	Animation       string  `yaml:"animation"`       // 动画
	AnimationSpeed  float64 `yaml:"animationSpeed"`  // 动画速度
	Subtitles       bool    `yaml:"subtitles"`       // 字幕
}
