package root

import (
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"zhaoyunxing92/dubbo-go-config/config/provider"
	"zhaoyunxing92/dubbo-go-config/config/service"
)
import (
	"zhaoyunxing92/dubbo-go-config/config/application"
	"zhaoyunxing92/dubbo-go-config/config/registry"
)

// Config root config
type Config struct {
	// Application config
	Application *application.Config `yaml:"application" json:"application"`
	// Registries config
	Registries map[string]*registry.Config `yaml:"registries" json:"registries"`
	// Provider config
	Provider *provider.Config `yaml:"provider" json:"provider"`
	// viper
	v *viper.Viper
	// validate
	validate *validator.Validate
	// trans translator
	trans translator.Translator
}

func (Config) Prefix() string {
	return "dubbo"
}
func (c *Config) GetApplicationConfig() (*application.Config, error) {
	conf := c.Application
	if err := conf.DefaultSetter(); err != nil {
		return &application.Config{}, err
	}
	//validate value
	if err := conf.Validate(c.validate, c.trans); err != nil {
		return &application.Config{}, err
	}
	return conf, nil
}

func (c *Config) GetRegistriesConfig() (map[string]*registry.Config, error) {
	registries := c.Registries

	if len(registries) <= 0 {
		reg := new(registry.Config)
		if err := reg.DefaultSetter(); err != nil {
			return nil, err
		}
		registries = make(map[string]*registry.Config, 1)
		registries["default"] = reg
		return registries, nil
	}
	for _, reg := range registries {
		if err := reg.DefaultSetter(); err != nil {
			return nil, err
		}
		if err := reg.Validate(c.validate, c.trans); err != nil {
			return nil, err
		}
	}
	return registries, nil
}

func (c *Config) GetRegistryIds() []string {
	registriesConfig, _ := c.GetRegistriesConfig()
	ids := make([]string, 0, len(registriesConfig))
	for key := range registriesConfig {
		ids = append(ids, key)
	}
	return ids
}

//GetProviderConfig services config
func (c *Config) GetProviderConfig() (*provider.Config, error) {
	var (
		services map[string]*service.Config
		err      error
	)
	pro := c.Provider
	if err = pro.DefaultSetter(); err != nil {
		return &provider.Config{}, err
	}

	if services, err = service.GetServiceConfig(c.GetRegistryIds(), pro.Services, c.validate); err != nil {
		return nil, err
	}
	// services config
	pro.Services = services
	if err = pro.Validate(c.validate, c.trans); err != nil {
		return nil, err
	}
	return pro, nil
}

func (c *Config) WriteConfig() error {
	c.v.Set(c.Prefix(), c)
	return c.v.WriteConfig()
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
