package config

type System struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`                // 环境
	Addr    int    `mapstructure:"addr" json:"addr" yaml:"addr"`             // 端口值
	OssType string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
}
