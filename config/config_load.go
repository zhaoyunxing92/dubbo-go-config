package config

import (
	"flag"
	"fmt"
)
import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"zhaoyunxing92/dubbo-go-config/config/root"
)

// config config
type config struct {
	// config file name default application
	name string
	// config file type default yaml
	genre string
	// config file path default ./conf
	path string
	// cache file default ture
	cache bool
}
type optionFunc func(*config)

func (fn optionFunc) apply(vc *config) {
	fn(vc)
}

type Option interface {
	apply(vc *config)
}

var bc *root.Config

func Load(opts ...Option) *root.Config {
	// pares CommandLine
	parseCommandLine()
	// get config
	conf := getConfig()

	for _, opt := range opts {
		opt.apply(conf)
	}

	v := viper.New()
	v.AddConfigPath(conf.path)
	v.SetConfigName(conf.name)
	v.SetConfigType(conf.genre)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	bc = new(root.Config)
	if err := v.UnmarshalKey(bc.Prefix(), &bc); err != nil {
		fmt.Println(err)
	}

	validate := validator.New()
	uni := translator.New(en.New(), zh.New())
	trans, _ := uni.GetTranslator("zh")
	_ = zh_trans.RegisterDefaultTranslations(validate, trans)

	bc.SetViper(v)
	bc.SetValidate(validate)
	bc.SetTranslator(trans)

	// cache file
	//if conf.cache {
	//	_ = bc.WriteConfig()
	//}
	return bc
}

func getConfig() *config {
	return &config{
		viper.GetString("name"),
		viper.GetString("genre"),
		viper.GetString("path"),
		viper.GetBool("cache"),
	}
}

func GetBaseConfig() *root.Config {
	return bc
}

//parseCommandLine parse command line
func parseCommandLine() {
	flag.Bool("cache", true, "config file cache")
	flag.String("name", "application.yaml", "config file name")
	flag.String("genre", "yaml", "config file type")
	flag.String("path", "./conf", "config file path default")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
}

func WithGenre(genre string) Option {
	return optionFunc(func(conf *config) {
		conf.genre = genre
	})
}

func WithPath(path string) Option {
	return optionFunc(func(conf *config) {
		conf.path = path
	})
}

func WithName(name string) Option {
	return optionFunc(func(conf *config) {
		conf.name = name
	})
}

func WithCache(cache bool) Option {
	return optionFunc(func(conf *config) {
		conf.cache = cache
	})
}
