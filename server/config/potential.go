package config

type Potential struct {
	MinLength int  `yaml:"min_length"` // 最小长度
	MaxLength int  `yaml:"max_length"` // 最大长度
	MaxWords  int  `yaml:"max_words"`  // 最大长度
	MinWords  int  `yaml:"min_words"`  // 最大长度
	Split     bool `yaml:"split"`      // 是否切割
}
