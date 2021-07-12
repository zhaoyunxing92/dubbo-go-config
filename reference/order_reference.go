package reference

import (
	"context"
	"github.com/apache/dubbo-go/config"
)

func init() {
	config.SetProviderService(new(OrderService))
}

type OrderService struct {
	// GetOrders
	GetOrders func(ctx context.Context, req []interface{}) error
}

func (os *OrderService) Reference() string {
	return "order-service"
}
