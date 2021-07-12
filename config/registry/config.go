package registry

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

import "github.com/creasty/defaults"

type Config struct {
	Protocol string `yaml:"protocol" json:"protocol"`
	Timeout  string `yaml:"timeout" json:"timeout"`
	Group    string `yaml:"group" json:"group"`
	Address  string `yaml:"address" json:"address"`
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
