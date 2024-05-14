package goject

import (
	"log"
	"reflect"
)

func GetServiceByName[T any](container *ServiceContainer, name string) T {
	service := container.getService(name)
	return service.(T)
}

func GetServiceByType[T any](container *ServiceContainer) T {
	serviceType := getGenericType[T](true)
	service := container.getService(serviceType.String())
	return service.(T)
}

func getGenericDefault[T any]() T {
	return *new(T)
}

func getGenericType[T any](key bool) reflect.Type {
	funcType := reflect.TypeOf(getGenericDefault[T])
	genericType := funcType.Out(0)
	if key && !validateKeyType(genericType) {
		log.Fatalf("key type must be a pointer, interface, or struct type")
	}

	if !key && !validateValueType(genericType) {
		log.Fatalf("value must be a pointer or struct type")
	}
	return genericType
}

func validateKeyType(keyType reflect.Type) bool {
	kind := keyType.Kind()
	if kind == reflect.Pointer {
		kind = keyType.Elem().Kind()
	}
	return kind == reflect.Interface || kind == reflect.Struct
}

func validateValueType(valueType reflect.Type) bool {
	kind := valueType.Kind()
	if kind == reflect.Pointer {
		kind = valueType.Elem().Kind()
	}
	return kind == reflect.Struct
}
