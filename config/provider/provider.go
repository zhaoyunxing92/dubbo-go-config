package provider

import (
	"errors"
	"github.com/creasty/defaults"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
	"zhaoyunxing92/dubbo-go-config/config/service"
)

type Config struct {
	// 是否检查
	Check bool `default:"true" yaml:"check" json:"check"`
	// 是否注册
	Register bool `default:"true" yaml:"register" json:"register"`
	// Services config
	Services map[string]*service.Config `yaml:"services" json:"services"`
}

func (c *Config) Prefix() string {
	return "dubbo.provider"
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
