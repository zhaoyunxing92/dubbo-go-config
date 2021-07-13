package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
	_ "zhaoyunxing92/dubbo-go-config/reference"
)

func TestYamlLoad(t *testing.T) {
	conf := Load(
		WithGenre("toml"),
		WithCache(false),
		WithPath("../conf/toml"),
		WithName("application.toml"),
	)

	application, _ := conf.GetApplicationConfig()

	registries, _ := conf.GetRegistriesConfig()

	providerConfig, _ := conf.GetProviderConfig()


	assert.Equal(t, application.Name, "dubbo-go")

	assert.NotNil(t, registries)

	assert.NotNil(t, providerConfig)
}
