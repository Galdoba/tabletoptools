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

func RollLawLevel(context profile.Profile, dice *dice.Dicepool) (string, error) {
	code := ""
	err := errors.New("roll was not made")
	switch context.Field(profile.GENERATION_METHOD) {
	case method.Basic:
		code, err = basicLawLevelRoll(context, dice)
		if err != nil {
			return code, fmt.Errorf("law level roll failed: %v", err)
		}
	}
	return code, err
}

func basicLawLevelRoll(context profile.Profile, dice *dice.Dicepool) (string, error) {
	prequisites, err := getBasicPrequisites(context, profile.KEY_Laws)
	if err != nil {
		return "", fmt.Errorf("prequisites aquisition failed: %v", err)
	}
	if prequisites[profile.KEY_Pops] == 0 {
		return "0", nil
	}
	govrDM := prequisites[profile.KEY_Govr]
	r := dice.Sroll("2d6") - 7 + govrDM
	r = values.BoundInt(r, 0, 33)
	return ehex.ToCode(r), err
}
