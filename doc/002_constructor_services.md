# Constructor Services
You can create your services from constructor functions using the `goject.NewServiceFromConstructor` function when creating the container. You must first create a constructor function so you can pass it to the `NewServiceFromConstructor` function.

The type of the service will be inferred by the return value of the constructor function.

## Example
```go
    // Constructor function for the FootballStrategyConcrete struct
    func NewFootballStrategyConcrete() *FootballStrategyConcrete {
        return &FootballStrategyConcrete{}
    }

    // Constructor function for the FootballPlayerConcrete struct
    func NewFootballPlayerConcrete(strategy *FootballStrategyConcrete) *FootballPlayerConcrete {
        return &FootballPlayerConcrete{
            Strategy: strategy,
        }
    }

    // Create the container and register the constructor functions
    var container = goject.NewServiceContainer(
        goject.NewServiceFromConstructor(testdata.NewFootballStrategyConcrete),
        goject.NewServiceFromConstructor(testdata.NewFootballPlayerConcrete),
    )

    // Retrieve a player from the container
    player := goject.GetServiceByType[*testdata.FootballPlayerConcrete](container)
```