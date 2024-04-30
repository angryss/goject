package goject

import (
	"reflect"
)

type ServiceContainer struct {
	services map[string]*ServiceDefinition
}

func NewServiceContainer(services ...*ServiceDefinition) *ServiceContainer {
	container := &ServiceContainer{
		services: make(map[string]*ServiceDefinition),
	}

	// register the container into itself
	containerService := NewServiceFromInstance(container)
	container.services[reflect.TypeOf(container).String()] = containerService

	// register the services
	for _, service := range services {
		container.services[service.key] = service
	}

	return container
}

func (ctr *ServiceContainer) getService(key string) interface{} {
	service, exists := ctr.services[key]
	if !exists {
		return nil
	}
	return service.value.getValue(ctr)
}
