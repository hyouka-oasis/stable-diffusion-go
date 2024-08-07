package config

type Book struct {
	Name     string `yaml:"name"`     // 小说名
	Language string `yaml:"language"` // 语言
	Batch    bool   `yaml:"batch"`    // 是否批量处理
}
