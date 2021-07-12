package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
	_ "zhaoyunxing92/dubbo-go-config/reference"
)

func TestYamlLoad(t *testing.T) {
	config := Load(
		WithGenre("yaml"),
		WithPrefix("dubbo"),
		WithPath("../conf/yaml"),
		WithName("application.yaml"),
	)

	application, _ := config.GetApplicationConfig()

	registries, _ := config.GetRegistriesConfig()

	providerConfig, _ := config.GetProviderConfig()

	assert.Equal(t, application.Name, "dubbo-go")

	assert.NotNil(t, registries)

	assert.NotNil(t, providerConfig)
}
