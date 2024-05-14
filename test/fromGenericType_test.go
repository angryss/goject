package test

import (
	"github.com/angryss/goject"
	"github.com/angryss/goject/test/testdata"
	"testing"
)

func TestFromGenericPointerType(t *testing.T) {
	// Arrange
	container := goject.NewServiceContainer(
		goject.NewGenericServiceFromType[testdata.FootballStrategy, *testdata.FootballStrategyLoose](),
		goject.NewGenericServiceFromType[*testdata.FootballPlayerLoose, *testdata.FootballPlayerLoose](),
	)

	// Act
	player := goject.GetServiceByType[*testdata.FootballPlayerLoose](container)

	// Assert
	if player == nil {
		t.Errorf("expected player to have a value")
	}

	if player.Strategy == nil {
		t.Errorf("expected player.Strategy to have a value")
	}
}

func TestFromGenericValueType(t *testing.T) {
	// Arrange
	container := goject.NewServiceContainer(
		goject.NewGenericServiceFromType[testdata.FootballStrategy, testdata.FootballStrategyLoose](),
		goject.NewGenericServiceFromType[*testdata.FootballPlayerLoose, *testdata.FootballPlayerLoose](),
	)

	// Act
	player := goject.GetServiceByType[*testdata.FootballPlayerLoose](container)

	// Assert
	if player == nil {
		t.Errorf("expected player to have a value")
	}

	if player.Strategy == nil {
		t.Errorf("expected player.Strategy to have a value")
	}
}
