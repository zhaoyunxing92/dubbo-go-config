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

	application, err := conf.GetApplicationConfig()
	assert.Nil(t, err)

	registries, err := conf.GetRegistriesConfig()
	assert.Nil(t, err)

	providerConfig, err := conf.GetProviderConfig()

	assert.Nil(t, err)

	assert.Equal(t, application.Name, "dubbo-go")

	assert.NotNil(t, registries)

	assert.NotNil(t, providerConfig)
}
