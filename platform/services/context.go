package services

import (
	"context"
	"reflect"
)

type ServiceKeyType string

const ServiceKey ServiceKeyType = "services"

type serviceMap map[reflect.Type]reflect.Value

func NewServiceContext(ctx context.Context) context.Context {
	if ctx.Value(ServiceKey) == nil {
		return context.WithValue(ctx, ServiceKey, make(serviceMap))
	} else {
		return ctx
	}
}
