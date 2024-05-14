package test

import (
	"github.com/angryss/goject"
	"github.com/angryss/goject/test/testdata"
	"testing"
)

func TestConstructorTypeMapping(t *testing.T) {
	// Arrange
	var container = goject.NewServiceContainer(
		goject.NewServiceFromConstructor(testdata.NewFootballStrategyLoose).Name("NewFootballStrategyLoose"),
		goject.NewServiceFromConstructor(testdata.NewFootballPlayerLoose).
			Name("NewFootballPlayerLoose").
			MapTypes(goject.NewTypeMapping(0, "", "NewFootballStrategyLoose")),
	)

	// Act
	player := goject.GetServiceByName[*testdata.FootballPlayerLoose](container, "NewFootballPlayerLoose")

	// Assert
	if player == nil {
		t.Errorf("'player' should not be nil")
		return
	}

	if player.Strategy == nil {
		t.Errorf("'player.Strategy' should not be nil")
	}
}

func TestStructTypeMapping(t *testing.T) {
	// Arrange
	var container = goject.NewServiceContainer(
		goject.NewGenericServiceFromType[testdata.FootballStrategy, *testdata.FootballStrategyLoose]().Name("NewFootballStrategyLoose"),
		goject.NewGenericServiceFromType[*testdata.FootballPlayerLoose, *testdata.FootballPlayerLoose]().
			Name("NewFootballPlayerLoose").
			MapTypes(goject.NewTypeMapping(-1, "testdata.FootballStrategy", "NewFootballStrategyLoose")),
	)

	// Act
	player := goject.GetServiceByName[*testdata.FootballPlayerLoose](container, "NewFootballPlayerLoose")

	// Assert
	if player == nil {
		t.Errorf("'player' should not be nil")
		return
	}

	if player.Strategy == nil {
		t.Errorf("'player.Strategy' should not be nil")
	}
}
