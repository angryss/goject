# Instance Services
You can register instances of services with the container using the `NewServiceFromInstance` and `NewGenericServiceFromInstance` functions. The difference is that `NewGenericServiceFromInstance` function will allow you to map an interface to a struct.

## Example
```go
    var strategy = testdata.NewFootballStrategyLoose()
    var container = goject.NewServiceContainer(
        // Register the strategy using the generic function that maps
        //   an interface type to the instance type.
        goject.NewGenericServiceFromInstance[testdata.FootballStrategy](strategy),
        goject.NewServiceFromConstructor(testdata.NewFootballPlayerLoose),
    )

    var instance = goject.GetServiceByType[*testdata.FootballPlayerLoose](container)

```