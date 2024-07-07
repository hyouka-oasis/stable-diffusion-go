package config

type StableDiffusionConfig struct {
	Url    string `yaml:"url"`    // url
	Height string `yaml:"height"` // 高
	Width  string `yaml:"width"`  // 宽
	Lora   string `yaml:"lora"`   // lora
}
