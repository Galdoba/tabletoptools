package continuation

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

func RollDiameterKm(ctx profile.Profile, dice *dice.Dicepool) (int, error) {
	size, sizeInt, err := ctx.Ehex(profile.KEY_Size)
	if err != nil {
		return 0, fmt.Errorf("continuation.RollDiameterKm failed: %v", err)
	}
	minimum := 0
	maximum := 0
	switch size {
	case "0", "R":
		return 0, nil
	case "S":
		minimum = 400
		maximum = 799
	case "1":
		minimum = 800
		maximum = 2399
	default:
		minimum = (sizeInt-1)*1600 + 800
		maximum = minimum + 1599
	}
	diameter := 50000
	increaseD3, increaseD6 := diameterIncreaseMaps()
	atempt := 1
	for diameter > maximum {

		r1 := dice.Sroll("1d3")
		r2 := dice.Sroll("1d6")
		r3 := dice.Sroll("1d100")
		diameter = minimum + increaseD3[r1] + increaseD6[r2] + r3
		fmt.Println("diameter atempt", atempt, ":", diameter)
	}
	if err := ctx.Inject(profile.KEY_Size_Dkm, fmt.Sprintf("%v", diameter)); err != nil {
		return 0, fmt.Errorf("continuation.RollDiameterKm failed: %v", err)
	}
	return diameter, nil
}

func diameterIncreaseMaps() (map[int]int, map[int]int) {
	increaseD3 := make(map[int]int)
	increaseD3[1] = 0
	increaseD3[2] = 600
	increaseD3[3] = 1200
	increaseD6 := make(map[int]int)
	increaseD6[1] = 0
	increaseD6[2] = 100
	increaseD6[3] = 200
	increaseD6[4] = 300
	increaseD6[5] = 400
	increaseD6[6] = 500
	return increaseD3, increaseD6
}
