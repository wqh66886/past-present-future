package model

type Auth struct {
	SecretKey   string `yaml:"secret_key" mapstructure:"secret_key"`
	ExpireTime  int    `yaml:"expire_time" mapstructure:"expire_time"`
	RefreshTime int    `yaml:"refresh_time" mapstructure:"refresh_time"`
}
