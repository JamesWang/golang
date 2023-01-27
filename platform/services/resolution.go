package services

import (
	"context"
	"errors"
	"reflect"
)

func GetService(target interface{}) error {
	return GetServiceForContext(context.Background(), target)
}

func GetServiceForContext(ctx context.Context, target interface{}) (err error) {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() == reflect.Ptr && targetValue.Elem().CanSet() {
		err = resolveServiceFromValue(ctx, targetValue)
	} else {
		err = errors.New("Type cannot be used as target")
	}
	return
}
