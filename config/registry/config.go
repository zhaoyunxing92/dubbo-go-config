package registry

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

import "github.com/creasty/defaults"

type Config struct {
	Protocol string `default:"zookpeer" yaml:"protocol" json:"protocol"`
	Timeout  string `default:"10s" yaml:"timeout" json:"timeout"`
	Group    string `default:"dubbo" yaml:"group" json:"group"`
	Address  string `default:"127.0.0.1:2181" yaml:"address" json:"address"`
	Register bool
}

func (*Config) Prefix() string {
	return "dubbo.registries"
}

func (c *Config) DefaultSetter() error {
	return defaults.Set(c)
}

func (c *Config) Validate(valid *validator.Validate, trans translator.Translator) error {
	if err := valid.Struct(c); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Translate(trans))
		}
		return errors.New(strings.Join(slice, ","))
	}
	return nil
}
