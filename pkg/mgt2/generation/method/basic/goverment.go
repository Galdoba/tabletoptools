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

func RollGoverment(context profile.Profile, dice *dice.Dicepool) (string, error) {
	code := ""
	err := errors.New("roll was not made")
	switch context.Field(profile.GENERATION_METHOD) {
	case method.Basic:
		code, err = basicGovermentRoll(context, dice)
		if err != nil {
			return code, fmt.Errorf("goverment roll failed: %v", err)
		}
	}
	return code, err
}

func basicGovermentRoll(context profile.Profile, dice *dice.Dicepool) (string, error) {
	prequisites, err := getBasicPrequisites(context, profile.KEY_Govr)
	if err != nil {
		return "", fmt.Errorf("prequisites aquisition failed: %v", err)
	}
	popDM := prequisites[profile.KEY_Pops]
	if popDM == 0 {
		return "0", nil
	}
	r := dice.Sroll("2d6") - 7 + popDM
	r = values.BoundInt(r, 0, 15)
	return ehex.ToCode(r), err
}
