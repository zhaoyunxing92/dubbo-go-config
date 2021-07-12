package root

import (
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)
import (
	"zhaoyunxing92/dubbo-go-config/config/application"
	"zhaoyunxing92/dubbo-go-config/config/registry"
	"zhaoyunxing92/dubbo-go-config/config/service"
)

// Config root config
type Config struct {
	// Application config
	Application application.Config `yaml:"application" json:"application"`
	// Registries config
	Registries map[string]registry.Config `yaml:"registries" json:"registries"`
	// Services config
	Services map[string]service.Config `yaml:"services" json:"services"`
	// viper
	v *viper.Viper
	// validate
	validate *validator.Validate
	// trans translator
	trans translator.Translator
}

func (c *Config) GetApplicationConfig() (application.Config, error) {
	// set default
	config := c.Application
	if err := config.DefaultSetter(); err != nil {
		return application.Config{}, err
	}
	//validate value
	if err := config.Validate(c.validate, c.trans); err != nil {
		return application.Config{}, err
	}
	return config, nil
}

func (c *Config) SetViper(v *viper.Viper) {
	c.v = v
}

func (c *Config) SetValidate(validate *validator.Validate) {
	c.validate = validate
}

func (c *Config) SetTranslator(trans translator.Translator) {
	c.trans = trans
}
