# ServiceContainer
The ServiceContainer is used to manage the lifecycle of services in your application. A service can be a struct or a pointer to a struct. An interface can optionally be mapped to a service. 

Create the container in your main function and add services to it. To create the container, use the `NewServiceContainer` function.

Note that the container can inject itself as a dependency to your services. To do this, add a constructor argument of the type `*goject.ServiceContainer`.

### Example
```go
func main() {
    // Create and populate the container
    var container = goject.NewServiceContainer(
        goject.NewServiceFromConstructor(testdata.NewFootballStrategyConcrete),
        goject.NewServiceFromConstructor(testdata.NewFootballPlayerConcrete),
    )

    // Retrieve a service from the container
    var player := goject.GetServiceFromType[*testdata.FootballPlayerConcrete(container)
    
```

