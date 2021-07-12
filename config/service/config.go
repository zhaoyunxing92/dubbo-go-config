package service

import "github.com/creasty/defaults"

//Config service
type Config struct {
	Id       string   `yaml:"id" json:"id"`
	Registry []string `yaml:"registry" json:"registry"`
	Register bool     `default:"true" yaml:"register" json:"register"`
}

func NewServiceConfig() *Config {
	conf := new(Config)
	_ = defaults.Set(conf)
	return conf
}

func (c *Config) DefaultSetter() error {
	return defaults.Set(c)
}