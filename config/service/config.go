package service

import (
	conf "github.com/apache/dubbo-go/config"
	"github.com/creasty/defaults"
	"zhaoyunxing92/dubbo-go-config/config/service/method"
)

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
	Methods map[string]*method.Config `yaml:"methods" json:"methods"`
}

//GetServiceConfig service config
func GetServiceConfig(registryIds []string, services map[string]*Config) map[string]*Config {
	if services == nil || len(services) <= 0 {
		services = make(map[string]*Config, len(conf.GetAllProviderService()))
	}
	// 遍历全部服务
	for key := range conf.GetAllProviderService() {
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
		//获取service下的方法
		services[key] = svc
	}
	return services
}
