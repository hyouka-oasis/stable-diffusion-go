package config

type System struct {
	Addr    int    `mapstructure:"addr" json:"addr" yaml:"addr"`             // 端口值
	OssType string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"` // Oss类型
}
