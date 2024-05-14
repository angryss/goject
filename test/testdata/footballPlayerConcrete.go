package testdata

type FootballPlayerConcrete struct {
	Strategy *FootballStrategyConcrete
}

func NewFootballPlayerConcrete(strategy *FootballStrategyConcrete) *FootballPlayerConcrete {
	return &FootballPlayerConcrete{
		Strategy: strategy,
	}
}
