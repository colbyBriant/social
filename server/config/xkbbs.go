package config

type XkBbs struct {
	Appid string `mapstructure:"appid" json:"appid" yaml:"appid"` // appid
	Level string `mapstructure:"level" json:"level" yaml:"level"` // 级别
}
