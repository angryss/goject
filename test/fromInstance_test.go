package test

import (
	"github.com/angryss/goject"
	"github.com/angryss/goject/test/testdata"
	"testing"
)

func TestFromInstance(t *testing.T) {
	// Arrange
	var strategy = testdata.NewFootballStrategyConcrete()
	var player = testdata.NewFootballPlayerConcrete(strategy)
	var container = goject.NewServiceContainer(
		goject.NewServiceFromInstance(player),
	)

	// Act
	var instance = goject.GetServiceByType[*testdata.FootballPlayerConcrete](container)

	// Assert
	if instance != player {
		t.Errorf("Expected instance to be equal to player")
	}
}
