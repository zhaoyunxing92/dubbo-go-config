package service

import (
	"errors"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)
import "github.com/creasty/defaults"

//Config service
type Config struct {
	Id        string   `yaml:"id" json:"id"`
	Interface string   `yaml:"interface" json:"interface"`
	Registry  []string `yaml:"registry" json:"registry"`
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
