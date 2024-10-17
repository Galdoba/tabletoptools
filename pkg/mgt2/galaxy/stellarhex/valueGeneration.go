package stellarhex

import "github.com/Galdoba/tabletoptools/pkg/dice"

func GenerateDensity(dice *dice.Dicepool) int {
	density := Standard
	for _, val := range []int{dice.Flux(), dice.Flux(), dice.Flux()} {
		if val < 7 {
			density--
		}
		if val > 7 {
			density++
		}
	}
	return density
}
