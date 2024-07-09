package config

type Audio struct {
	Voice    string `yaml:"voice"`     // 角色名称
	Rate     string `yaml:"rate"`      // 语速
	Volume   string `yaml:"volume"`    // 音量
	Pitch    string `yaml:"pitch"`     // 分贝
	ProxyUrl string `yaml:"proxy_url"` // 代理路径
}
