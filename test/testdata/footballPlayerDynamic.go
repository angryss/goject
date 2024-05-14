package testdata

import "github.com/angryss/goject"

type FootballPlayerDynamic struct {
	Strategy FootballStrategy
}

func NewFootballPlayerDynamic(container *goject.ServiceContainer) *FootballPlayerDynamic {
	var player = &FootballPlayerDynamic{}
	player.Strategy = goject.GetServiceByType[FootballStrategy](container)
	return player
}
