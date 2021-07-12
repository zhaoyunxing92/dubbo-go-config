package config

import (
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Config interface {
	// DefaultSetter set default value
	DefaultSetter() error
	// Validate value
	Validate(valid *validator.Validate, trans translator.Translator) error
}
