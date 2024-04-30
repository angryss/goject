# Service Lifecycle
goject supports two different lifecycle patterns for your services: Singleton and transient. Singleton is the default lifecycle. Use the `Lifecycle` method to change the lifecycle during registration.

With the singleton lifecycle, goject will only ever create one instance of the service per container. Subsequent requests will return the same instance.

With the transient lifecycle, goject will create a new instance of the service for each request. Subsequent requests will return new instances.

## Example
```go
    var container = goject.NewServiceContainer(
        goject.NewServiceFromConstructor(testdata.NewFootballStrategyConcrete).Lifecycle(goject.LifecycleTransient),
        goject.NewServiceFromConstructor(testdata.NewFootballPlayerConcrete).Lifecycle(goject.LifecycleTransient),
    )
```