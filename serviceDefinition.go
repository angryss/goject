package goject

import (
	"log"
	"reflect"
)

type ServiceDefinition struct {
	key   string
	value *serviceValue
}

func NewGenericServiceFromType[K any, V any]() *ServiceDefinition {
	keyType := getGenericType[K](true)
	valueType := getGenericType[V](false)

	return newService(
		keyType.String(),
		nil,
		nil,
		LifecycleSingleton,
		nil,
		valueType,
	)
}

func NewGenericServiceFromInstance[T any](instance interface{}) *ServiceDefinition {
	keyType := getGenericType[T](true)
	return newService(
		keyType.String(),
		nil,
		nil,
		LifecycleSingleton,
		instance,
		nil,
	)
}

func NewServiceFromConstructor(constructor any) *ServiceDefinition {
	validateConstructor(constructor)
	constructorType := reflect.TypeOf(constructor)
	if constructorType.NumOut() != 1 {
		log.Fatalf("constructor must return only one value")
	}
	valueType := constructorType.Out(0)
	return newService(
		valueType.String(),
		constructor,
		nil,
		LifecycleSingleton,
		nil,
		nil)
}

func NewServiceFromInstance(instance interface{}) *ServiceDefinition {
	serviceType := reflect.TypeOf(instance)
	return newService(
		serviceType.String(),
		nil,
		nil,
		LifecycleSingleton,
		instance,
		nil)
}

func (s *ServiceDefinition) Lifecycle(lifecycle int) *ServiceDefinition {
	if lifecycle != LifecycleSingleton && lifecycle != LifecycleTransient {
		log.Fatalf("lifecycle must be either LifecycleSingleton or LifecycleTransient")
	}
	s.value.lifecycle = lifecycle
	return s
}

func (s *ServiceDefinition) MapTypes(mappings ...*TypeMapping) *ServiceDefinition {
	if s.value.constructor == nil && s.value.valueType == nil {
		log.Fatalf("service must have a constructor or value type to map types")
	}
	s.value.constructorArgMappings = mappings
	return s
}

func (s *ServiceDefinition) Name(key string) *ServiceDefinition {
	s.key = key
	return s
}

func newService(
	key string,
	constructor any,
	constructorArgMappings []*TypeMapping,
	lifecycle int,
	value interface{},
	valueType reflect.Type,
) *ServiceDefinition {
	return &ServiceDefinition{
		key: key,
		value: &serviceValue{
			constructor:            constructor,
			constructorArgMappings: constructorArgMappings,
			lifecycle:              lifecycle,
			value:                  value,
			valueType:              valueType,
		},
	}
}

func validateConstructor(constructor any) {
	if reflect.ValueOf(constructor).Kind() != reflect.Func {
		log.Fatalf("constructor must be a function")
	}
}
