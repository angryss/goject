package test

import (
	"github.com/angryss/goject"
	"github.com/angryss/goject/test/testdata"
	"testing"
)

func TestSelfInjection(t *testing.T) {
	// Arrange
	var container = goject.NewServiceContainer(
		goject.NewServiceFromConstructor(testdata.NewFootballStrategyLoose),
		goject.NewServiceFromConstructor(testdata.NewFootballPlayerDynamic),
	)

	// Act
	player := goject.GetServiceByType[*testdata.FootballPlayerDynamic](container)

	// Assert
	if player == nil {
		t.Errorf("'player' should not be nil")
		return
	}

	if player.Strategy == nil {
		t.Errorf("'player.Strategy' should not be nil")
	}
}
