# Overview
**goject** is a simple and light-weight dependency injection framework for Go. It simplifies the management of dependencies in your Go projects by providing a range of features that promote loose coupling and testability. Manual dependency management should be sufficient for simple projects. goject is well-suited for projects that perform a lot of different tasks, such as a CLI built using the Cobra framework.

## Features

- Dependency Graph Initialization: goject automatically initializes the dependency graph, resolving dependencies and ensuring that all required instances are created and properly injected.
- Constructor-Based Injection: goject allows you to define dependencies using constructor functions, making it easy to create and inject instances of your structs.
- Typed Injection: goject allows you to define dependencies as interfaces and structs. It will use its own internal reflection-based constructor.
- Interface-to-Struct Mappings: You can map interfaces to concrete struct types, enabling you to inject dependencies based on their interfaces rather than their concrete implementations.
- Lifecycle Management: goject supports both singleton and transient lifecycle management for constructor-injected dependencies. You can specify whether a dependency should be shared across the application (singleton) or created fresh for each injection (transient).
- Self-Injection: goject has the ability to inject itself into your structs, allowing for dynamic dependency resolution. This is particularly useful when you need to resolve dependencies at runtime based on specific conditions or configurations.

# Installation
To install goject, use the following command:
```bash
go get github.com/angryss/goject
```

# Documentation
- [The ServiceContainer](doc/001_serviceContainer.md)
- [Constructor Services](doc/002_constructorServices.md)
- [Instance Services](doc/003_instanceServices.md)
- [Typed Services](doc/004_typedServices.md)
- [Service Lifecycle](doc/005_serviceLifecycle.md)
- [Type Mappings](doc/006_typeMappings.md)
