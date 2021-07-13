package reference

import (
	"context"
	"zhaoyunxing92/dubbo-go-config/config/service"
)

func init() {
	service.SetProviderService(new(HelloService))
}

type HelloService struct {
	// say hello
	Say func(ctx context.Context, req []interface{}) error
}

func (HelloService) Reference() string {
	return "hello-service"
}
