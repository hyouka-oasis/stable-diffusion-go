package config

type StableDiffusionConfig struct {
	Url    string `yaml:"url"`    // url
	Height int    `yaml:"height"` // 高
	Width  int    `yaml:"width"`  // 宽
	Lora   string `yaml:"lora"`   // lora
}
