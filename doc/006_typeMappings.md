# Type Mappings
You can create string names for types that your register into the ServiceContainer by using the `MapTypes` function.

Use the 'Name' function on the object to be injected. Use the `MapTypes`  mappings on the receiver.

You can map a type using either an index position or a source type name. For a constructor, the index position refers to the argument number. For a generic type, the index position refers to the field number. Both start with 0 as the first field. 

## Example
```go
    var container = goject.NewServiceContainer(
        goject.NewServiceFromConstructor(testdata.NewFootballStrategyLoose).Name("NewFootballStrategyLoose"),
        goject.NewServiceFromConstructor(testdata.NewFootballPlayerLoose).
            Name("NewFootballPlayerLoose").
            MapTypes(goject.NewTypeMapping(0, "", "NewFootballStrategyLoose")),
    )

    // Act
    player := goject.GetServiceByName[*testdata.FootballPlayerLoose](container, "NewFootballPlayerLoose")
```