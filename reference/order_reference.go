package reference

import (
	"context"
	"zhaoyunxing92/dubbo-go-config/config/service"
)

func init() {
	service.SetProviderService(new(OrderService))
}

type OrderService struct {
	// GetOrders
	GetOrders func(ctx context.Context, req []interface{}) error
}

func (OrderService) Reference() string {
	return "order-service"
}
