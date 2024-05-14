package test

import (
	"github.com/angryss/goject"
	"github.com/angryss/goject/test/testdata"
	"testing"
)

func TestTransientLifecycle(t *testing.T) {
	// Arrange
	var container = goject.NewServiceContainer(
		goject.NewServiceFromConstructor(testdata.NewFootballStrategyConcrete).Lifecycle(goject.LifecycleTransient),
		goject.NewServiceFromConstructor(testdata.NewFootballPlayerConcrete).Lifecycle(goject.LifecycleTransient),
	)

	// Act
	var player1 = goject.GetServiceByType[*testdata.FootballPlayerConcrete](container)
	var player2 = goject.GetServiceByType[*testdata.FootballPlayerConcrete](container)

	// Assert
	if player1 == player2 {
		t.Error("player1 and player2 should be different instances")
	}
}
