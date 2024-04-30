package testdata

type FootballPlayerLoose struct {
	Strategy FootballStrategy
}

func NewFootballPlayerLoose(strategy FootballStrategy) *FootballPlayerLoose {
	return &FootballPlayerLoose{
		Strategy: strategy,
	}
}
