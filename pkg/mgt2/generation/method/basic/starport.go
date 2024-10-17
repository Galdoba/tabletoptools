package basic

import (
	"errors"
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/values"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

func RollStarport(context profile.Profile, dice *dice.Dicepool) (string, error) {
	code := ""
	err := errors.New("roll was not made")
	switch context.Field(profile.GENERATION_METHOD) {
	case method.Basic:
		code, err = basicStarportRoll(context, dice)
		if err != nil {
			return code, fmt.Errorf("law level roll failed: %v", err)
		}
	}
	return code, err
}

func basicStarportRoll(context profile.Profile, dice *dice.Dicepool) (string, error) {
	prequisites, err := getBasicPrequisites(context, profile.KEY_Port)
	if err != nil {
		return "", fmt.Errorf("prequisites aquisition failed: %v", err)
	}
	if prequisites[profile.KEY_Pops] == 0 {
		return "X", nil
	}
	pops := prequisites[profile.KEY_Pops]
	dm := 0
	switch pops {
	case 1, 2:
		dm = -2
	case 3, 4:
		dm = -1
	case 8, 9:
		dm = 1
	case 10, 11, 12:
		dm = 2
	}

	r := dice.Sroll("2d6") + dm
	r = values.BoundInt(r, 0, 33)
	switch r {
	case 0, 1, 2:
		return "X", nil
	case 3, 4:
		return "E", nil
	case 5, 6:
		return "D", nil
	case 7, 8:
		return "C", nil
	case 9, 10:
		return "B", nil
	}
	return "A", err
}
