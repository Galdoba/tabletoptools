package traveller

type Options func(*creationOpt)

type creationOpt struct {
	race string
	dice DiceRoller
}

func defaultCreationOptions() creationOpt {
	return creationOpt{
		race: "Vilani",
		dice: nil,
	}
}

func WithRace(race string) Options {
	return func(opt *creationOpt) {
		opt.race = race
	}
}

func WithDice(dice DiceRoller) Options {
	return func(opt *creationOpt) {
		opt.dice = dice
	}
}
