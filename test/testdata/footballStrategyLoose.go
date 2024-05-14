package testdata

import "fmt"

type FootballStrategyLoose struct {
}

func NewFootballStrategyLoose() FootballStrategy {
	return &FootballStrategyLoose{}
}

func (f *FootballStrategyLoose) Play() {
	fmt.Println("loose football strategy")
}
