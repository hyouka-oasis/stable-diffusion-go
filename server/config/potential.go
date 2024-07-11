package config

type Potential struct {
	MinLength string `yaml:"minLength"` // 最小长度
	MaxLength string `yaml:"maxLength"` // 最大长度
	MaxWords  string `yaml:"maxWords"`  // 最大长度
	MinWords  string `yaml:"minWords"`  // 最大长度
}
