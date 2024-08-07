package config

type StableDiffusionConfig struct {
	Url      string `yaml:"url"`                                // url
	Height   int    `yaml:"height"`                             // 高
	Width    int    `yaml:"width"`                              // 宽
	Lora     string `yaml:"lora"`                               // lora
	ArgsJson string `yaml:"args_json" mapstructure:"args_json"` // 额外的配置
}
