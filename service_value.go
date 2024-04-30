package goject

import (
	"log"
	"reflect"
)

const (
	LifecycleSingleton = 1
	LifecycleTransient = 2
)

type serviceValue struct {
	constructor            any
	constructorArgMappings []*TypeMapping
	lifecycle              int
	value                  interface{}
	valueType              reflect.Type
}

func (val *serviceValue) createFromConstructor(container *ServiceContainer) interface{} {
	if val.lifecycle == LifecycleSingleton && val.value != nil {
		return val.value
	}

	constructorType := reflect.TypeOf(val.constructor)
	numArgs := constructorType.NumIn()
	var args []reflect.Value
	if numArgs > 0 {
		args = make([]reflect.Value, numArgs)
	}

	for i := 0; i < numArgs; i++ {
		argType := constructorType.In(i)
		argTypeName := val.getConstructorArgumentTypeName(argType, i)
		argValue := container.getService(argTypeName)
		if argValue == nil {
			log.Fatalf("unable to resolve service %s", argTypeName)
		}

		args[i] = reflect.ValueOf(argValue)
	}

	constructorVal := reflect.ValueOf(val.constructor)
	response := constructorVal.Call(args)
	result := response[0].Interface()
	if val.lifecycle == LifecycleSingleton {
		val.value = result
	}

	return result
}

func (val *serviceValue) createFromType(container *ServiceContainer) interface{} {
	if val.lifecycle == LifecycleSingleton && val.value != nil {
		return val.value
	}

	valueType := val.valueType
	if valueType.Kind() == reflect.Pointer {
		valueType = valueType.Elem()
	}
	serviceInstance := reflect.New(valueType)
	serviceInterface := serviceInstance.Interface()
	serviceElem := serviceInstance.Elem()

	numField := serviceElem.NumField()
	for idx := 0; idx < numField; idx++ {
		field := serviceElem.Field(idx)

		if !field.CanSet() {
			break
		}

		fieldType := field.Type()
		if !validateKeyType(fieldType) {
			break
		}

		fieldValidateType := fieldType
		if fieldValidateType.Kind() == reflect.Pointer {
			fieldValidateType = fieldValidateType.Elem()
		}

		if fieldValidateType.Kind() != reflect.Struct && fieldValidateType.Kind() != reflect.Interface {
			break
		}

		fieldName := val.getConstructorArgumentTypeName(fieldType, idx)
		fieldValue := container.getService(fieldName)
		if fieldValue != nil {
			field.Set(reflect.ValueOf(fieldValue))
		}
	}

	if val.lifecycle == LifecycleSingleton {
		val.value = serviceInterface
	}

	return serviceInterface
}

func (val *serviceValue) getConstructorArgumentTypeName(
	argType reflect.Type,
	idx int,
) string {
	if val.constructorArgMappings == nil {
		return argType.String()
	}

	for argIdx, mapping := range val.constructorArgMappings {
		if argType.String() == mapping.TypeName || argIdx == idx {
			return mapping.ServiceKey
		}
	}

	return argType.String()
}

func (val *serviceValue) getValue(container *ServiceContainer) interface{} {
	if val.lifecycle == LifecycleSingleton && val.value != nil {
		return val.value
	}

	if val.constructor != nil {
		return val.createFromConstructor(container)
	}

	if val.valueType != nil {
		return val.createFromType(container)
	}

	return nil
}
