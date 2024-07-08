package config

type Config struct {
	Book                  Book                  `yaml:"book" mapstructure:"book"`
	TranslateConfig       TranslateConfig       `yaml:"translate_config" mapstructure:"translate_config"`
	Ollama                OllamaConfig          `yaml:"ollama" mapstructure:"ollama"`
	Potential             Potential             `yaml:"potential" mapstructure:"potential"`
	StableDiffusionConfig StableDiffusionConfig `yaml:"stable_diffusion" mapstructure:"stable_diffusion"`
	Audio                 Audio                 `yaml:"audio" mapstructure:"audio"`
}
