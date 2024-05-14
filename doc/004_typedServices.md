# Typed Services
A typed service allows you to pass in a struct, or an interface and a struct, and goject will automatically instantiate the struct type. You can use struct values or struct pointers. Both are supported. goject will try to automatically assign values to fields in the struct if it recognizes the types.

## Example
```go
    container := goject.NewServiceContainer(
        // Map an interface to a struct pointer
        goject.NewGenericServiceFromType[testdata.FootballStrategy, *testdataFootballStrategyLoose](),
    
        // Map a struct pointer to the same struct pointer
        goject.NewGenericServiceFromType[*testdata.FootballPlayerLoose, *testdata.FootballPlayerLoose](),
    )

    // Create an instance of *testdata.FootballPlayerLoose. goject will
    // automatically populate the Strategy field of the 
    // FootballPlayerLoose struct.
    player := goject.GetServiceByType[*testdata.FootballPlayerLoose](container)

```