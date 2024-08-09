package config

type Config struct {
	Video  Video  `yaml:"video" mapstructure:"video"`
	Sqlite Sqlite `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Local  Local  `mapstructure:"local" json:"local" yaml:"local"`
}
