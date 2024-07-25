package config

type Config struct {
	Book                  Book                  `yaml:"book" mapstructure:"book"`
	TranslateConfig       TranslateConfig       `yaml:"translate_config" mapstructure:"translate_config"`
	Ollama                OllamaConfig          `yaml:"ollama" mapstructure:"ollama"`
	Potential             Potential             `yaml:"potential" mapstructure:"potential"`
	StableDiffusionConfig StableDiffusionConfig `yaml:"stable_diffusion" mapstructure:"stable_diffusion"`
	Audio                 Audio                 `yaml:"audio" mapstructure:"audio"`
	Video                 Video                 `yaml:"video" mapstructure:"video"`
	Sqlite                Sqlite                `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Zap                   Zap                   `mapstructure:"zap" json:"zap" yaml:"zap"`
	System                System                `mapstructure:"system" json:"system" yaml:"system"`
	JWT                   JWT                   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Local                 Local                 `mapstructure:"local" json:"local" yaml:"local"`
}
