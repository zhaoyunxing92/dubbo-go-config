package application

import (
	"errors"
	"github.com/creasty/defaults"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

// Config application
type Config struct {
	Name    string `default:"zyx" yaml:"name" json:"name" validate:"required"`
	Module  string `yaml:"module" json:"module"`
	Version string `default:"1.0.0" yaml:"version" json:"version"`
	Owner   string `default:"zyx" yaml:"owner" json:"owner"`
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
