package test

import (
	"testing"

	"github.com/angryss/goject"
	"github.com/angryss/goject/test/testdata"
)

func TestFromGenericInstance(t *testing.T) {
	// Arrange
	var strategy = testdata.NewFootballStrategyLoose()
	var container = goject.NewServiceContainer(
		goject.NewGenericServiceFromInstance[testdata.FootballStrategy](strategy),
		goject.NewServiceFromConstructor(testdata.NewFootballPlayerLoose),
	)

	// Act
	var instance = goject.GetServiceByType[*testdata.FootballPlayerLoose](container)

	// Assert
	if instance.Strategy != strategy {
		t.Errorf("expected the injected strategy to be the same as the one passed to the container")
	}
}
