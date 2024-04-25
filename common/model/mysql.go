package model

type Mysql struct {
	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	Database          string `yaml:"database"`
	MaxPoolSize       int    `yaml:"max_pool_size" mapstructure:"max_pool_size"`
	MaxOpenConns      int    `yaml:"max_open_conns" mapstructure:"max_open_conns"`
	MaxLifeTime       int    `yaml:"max_life_time" mapstructure:"max_life_time"`
	ConnectionTimeout int    `yaml:"connection_timeout" mapstructure:"connection_timeout"`
}
