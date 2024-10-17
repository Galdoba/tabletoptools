package basic

import (
	"errors"
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/ehex"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/values"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

func RollTemperature(ctx profile.Profile, dice *dice.Dicepool) (string, error) {
	code := ""
	err := errors.New("roll was not made")
	switch ctx.Field(profile.GENERATION_METHOD) {
	case method.Basic:
		code, err = basicTemperatureRoll(ctx, dice)
		if err != nil {
			return code, fmt.Errorf("hydrosphere roll failed: %v", err)
		}
	}
	return code, err

}

func basicTemperatureRoll(ctx profile.Profile, dice *dice.Dicepool) (string, error) {
	prequisites, err := getBasicPrequisites(ctx, profile.KEY_Atmo_Temp)
	if err != nil {
		return "", fmt.Errorf("prequisites aquisition failed: %v", err)
	}
	atmo := prequisites[profile.KEY_Atmo]
	dm := 0
	switch atmo {
	case 0, 1:
		return "X", nil
	case 2, 3:
		dm += -2
	case 4, 5, 14:
		dm += -1
	case 6, 7:
		dm += 0
	case 8, 9:
		dm += 1
	case 10, 13, 15:
		dm += 2
	case 11, 12:
		dm += +6
	}
	r := dice.Sroll("2d6") + dm
	r = values.BoundInt(r, 0, 10)
	return ehex.ToCode(r), err
}
