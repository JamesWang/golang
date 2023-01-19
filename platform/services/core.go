package services

import (
	"context"
	"fmt"
	"reflect"
)

type BindingMap struct {
	factoryFunc reflect.Value
	lifecycle
}

var services = make(map[reflect.Type]BindingMap)

func addService(life lifecycle, factoryFunc interface{}) (err error) {
	factoryFuncType := reflect.TypeOf(factoryFunc)
	if factoryFuncType.Kind() == reflect.Func && factoryFuncType.NumOut() == 1 {
		services[factoryFuncType.Out(0)] = BindingMap{
			factoryFunc: reflect.ValueOf(factoryFunc),
			lifecycle:   life,
		}
	} else {
		err = fmt.Errorf("type cannot be used as service: %v", factoryFuncType)
	}
	return
}

var contextReference = (*context.Context)(nil)
var contextReferenceType = reflect.TypeOf(contextReference).Elem()

func resolveServiceFromValue(ctx context.Context, val reflect.Value) (err error) {
	serviceType := val.Elem().Type()
	if serviceType == contextReferenceType {
		val.Elem().Set(reflect.ValueOf(ctx))
	} else if binding, found := services[serviceType]; found {
		if binding.lifecycle == Scoped {
			resolveScopedService(ctx, val, binding)
		} else {
			val.Elem().Set(invokeFunction(ctx, binding.factoryFunc)[0])
		}
	} else {
		err = fmt.Errorf("cannot find service %v", serviceType)
	}
	return
}

func resolveScopedService(ctx context.Context, val reflect.Value, binding BindingMap) {
	sMap, ok := ctx.Value(ServiceKey).(serviceMap)
	if ok {
		serviceVal, ok := sMap[val.Type()]
		if !ok {
			serviceVal = invokeFunction(ctx, binding.factoryFunc)[0]
			sMap[val.Type()] = serviceVal
		}
		val.Elem().Set(serviceVal)
	} else {
		val.Elem().Set(invokeFunction(ctx, binding.factoryFunc)[0])
	}
	return
}

func invokeFunction(ctx context.Context, f reflect.Value, otherArgs ...interface{}) []reflect.Value {
	return f.Call(resolveServiceFromArguments(ctx, f, otherArgs...))
}

func resolveServiceFromArguments(ctx context.Context, f reflect.Value, otherArgs ...interface{}) []reflect.Value {
	params := make([]reflect.Value, f.Type().NumIn())
	i := 0
	if otherArgs != nil {
		for ; i < len(otherArgs); i++ {
			params[i] = reflect.ValueOf(otherArgs[i])
		}
	}
	for ; i < len(params); i++ {
		pType := f.Type().In(i)
		pVal := reflect.New(pType)
		err := resolveServiceFromValue(ctx, pVal)
		if err != nil {
			panic(err)
		}
		params[i] = pVal.Elem()
	}
	return params
}
