package test

import (
	"testing"

	"github.com/angryss/goject"
	"github.com/angryss/goject/test/testdata"
)

func TestFromConstructor(t *testing.T) {
	// Arrange
	var container = goject.NewServiceContainer(
		goject.NewServiceFromConstructor(testdata.NewFootballStrategyConcrete),
		goject.NewServiceFromConstructor(testdata.NewFootballPlayerConcrete),
	)

	// Act
	player := goject.GetServiceByType[*testdata.FootballPlayerConcrete](container)

	// Assert
	if player == nil {
		t.Errorf("'player' should not be nil")
		return
	}

	if player.Strategy == nil {
		t.Errorf("'player.Strategy' should not be nil")
	}
}
