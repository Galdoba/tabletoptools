package galaxy

import "github.com/Galdoba/tabletoptools/pkg/dice"

type Galaxy struct {
	Dice *dice.Dicepool
	Seed string
	Grid map[HexCoord]SpaceHex
}

/*
Galaxy.Charts[Coords].GenerateStars(dice)
Galaxy.Charts[Coords].GenerateOrbits(dice)
Galaxy.Charts[Coords].GeneratePlanets(dice)
Galaxy.Charts[Coords].GenerateMoons(dice)
Galaxy.Charts[Coords].GenerateLife(dice)
Galaxy.Charts[Coords].GenerateSocial(dice)

*/
