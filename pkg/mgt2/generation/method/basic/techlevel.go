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

func RollTechLevel(context profile.Profile, dice *dice.Dicepool) (string, error) {
	code := ""
	err := errors.New("roll was not made")
	switch context.Field(profile.GENERATION_METHOD) {
	case method.Basic:
		code, err = basicTechLevelRoll(context, dice)
		if err != nil {
			return code, fmt.Errorf("hydrosphere roll failed: %v", err)
		}
	}
	return code, err
}

func basicTechLevelRoll(context profile.Profile, dice *dice.Dicepool) (string, error) {
	prequisites, err := getBasicPrequisites(context, profile.KEY_TL)
	if err != nil {
		return "", fmt.Errorf("prequisites aquisition failed: %v", err)
	}
	size := prequisites[profile.KEY_Size]
	atmo := prequisites[profile.KEY_Atmo]
	hydr := prequisites[profile.KEY_Hydr]
	pops := prequisites[profile.KEY_Pops]
	govr := prequisites[profile.KEY_Govr]
	port := prequisites[profile.KEY_Port]
	if pops == 0 {
		return "0", nil
	}
	dm := 0
	switch port {
	case ehex.ValueOf("A"):
		dm += 6
	case ehex.ValueOf("B"):
		dm += 4
	case ehex.ValueOf("C"):
		dm += 2
	case ehex.ValueOf("X"):
		dm += -4
	}
	switch size {
	case 0, 1:
		dm += 2
	case 2, 3, 4:
		dm += 1
	}
	switch atmo {
	case 0, 1, 2, 3, 10, 11, 12, 13, 14, 15, 16, 17:
		dm += 1
	}
	switch hydr {
	case 0, 9:
		dm += 1
	case 10:
		dm += 2
	}
	switch pops {
	case 1, 2, 3, 4, 5, 8:
		dm += 1
	case 9:
		dm += 2
	case 10:
		dm += 2
	case 11:
		dm += 3
	case 12:
		dm += 4
	}
	switch govr {
	case 0, 5:
		dm += 1
	case 7:
		dm += 1
	case 13, 14:
		dm += -2
	}

	r := dice.Sroll("1d6") + dm
	r = values.BoundInt(r, 0, 33)
	return ehex.ToCode(r), err
}
