package protocol

import (
	"errors"
	"github.com/creasty/defaults"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Config struct {
	Name string `yaml:"name" json:"name,omitempty" property:"name"`
	Ip   string `yaml:"ip" json:"ip,omitempty" property:"ip"`
	Port string `yaml:"port"  json:"port,omitempty" property:"port"`
}

func (*Config) Prefix() string {
	return "dubbo.protocols"
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
