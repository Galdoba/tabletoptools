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

func RollHydrosphere(context profile.Profile, dice *dice.Dicepool) (string, error) {
	code := ""
	err := errors.New("roll was not made")
	switch context.Field(profile.GENERATION_METHOD) {
	case method.Basic:
		code, err = basicHydrosphereRoll(context, dice)
		if err != nil {
			return code, fmt.Errorf("hydrosphere roll failed: %v", err)
		}
	}
	return code, err
}

func basicHydrosphereRoll(context profile.Profile, dice *dice.Dicepool) (string, error) {
	prequisites, err := getBasicPrequisites(context, profile.KEY_Hydr)
	if err != nil {
		return "", fmt.Errorf("prequisites aquisition failed: %v", err)
	}
	size := prequisites[profile.KEY_Size]
	atmo := prequisites[profile.KEY_Atmo]
	temp := prequisites[profile.KEY_Temperature]
	switch size {
	case 0, 1:
		return "0", nil
	}
	dm := 0
	switch atmo {
	case 0, 1, 10, 11, 12, 13, 14, 15, 16, 17:
		dm += -4
	}
	if atmo != 14 && atmo != 15 {
		switch temp {
		case 10, 11:
			dm += -2
		default:
			if temp >= 12 {
				dm += -6
			}
		}
	}
	r := dice.Sroll("2d6") - 7 + atmo
	r = values.BoundInt(r, 0, 10)
	return ehex.ToCode(r), err
}
