package service

import (
	"errors"
	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"strings"
	"zhaoyunxing92/dubbo-go-config/common"
	"zhaoyunxing92/dubbo-go-config/config/service/method"
)

var (
	proServices = map[string]common.RPCService{} // service name -> service
)

// SetProviderService is called by init() of implement of RPCService
func SetProviderService(service common.RPCService) {
	proServices[service.Reference()] = service
}

func GetAllProviderService() map[string]common.RPCService {
	return proServices
}

//Config service
type Config struct {
	Id string `yaml:"id" json:"id"`
	// 注册中心
	Registry []string `yaml:"registry" json:"registry"`
	// 是否需要注册
	Register bool `yaml:"register" json:"register"`
	// 负载
	LoadBalance string `default:"random" yaml:"load-balance" json:"load-balance,omitempty"`
	// 方法
	Methods map[string]*method.Config `validate:"required" yaml:"methods" json:"methods"`
}

func (c *Config) Validate(valid *validator.Validate) error {
	if err := valid.Struct(c); err != nil {
		errs := err.(validator.ValidationErrors)
		var slice []string
		for _, msg := range errs {
			slice = append(slice, msg.Error())
		}
		return errors.New(strings.Join(slice, ","))
	}
	return valid.Struct(c)
}

//GetServiceConfig service config
func GetServiceConfig(registryIds []string, services map[string]*Config,
	valid *validator.Validate) (map[string]*Config, error) {

	if services == nil || len(services) <= 0 {
		services = make(map[string]*Config, len(proServices))
	}
	// 遍历全部服务
	for key := range proServices {
		var (
			svc   *Config
			exist bool
		)
		//存在配置了使用用户的配置
		if svc, exist = services[key]; exist {
			if len(svc.Registry) <= 0 {
				svc.Registry = registryIds
			}
		} else {
			//使用默认配置
			svc = new(Config)
			svc.Register = true
			svc.Registry = registryIds
		}
		_ = defaults.Set(svc)
		svc.Id = key
		if err := svc.Validate(valid); err != nil {
			return nil, err
		}
		//获取service下的方法
		services[key] = svc
	}
	return services, nil
}
