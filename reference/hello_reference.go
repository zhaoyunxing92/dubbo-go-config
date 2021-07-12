package reference

import (
	"context"
	"github.com/apache/dubbo-go/config"
)

func init() {
	config.SetProviderService(new(HelloService))
}

type HelloService struct {
	// say hello
	Say func(ctx context.Context, req []interface{}) error
}

func (hs *HelloService) Reference() string {
	return "hello-service"
}
