package testdata

import "fmt"

type FootballStrategyConcrete struct {
}

func NewFootballStrategyConcrete() *FootballStrategyConcrete {
	return &FootballStrategyConcrete{}
}

func (f *FootballStrategyConcrete) Play() {
	fmt.Println("concrete football strategy")
}
