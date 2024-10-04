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

func RollAtmosphere(context profile.Profile, dice *dice.Dicepool) (string, error) {
	code := ""
	err := errors.New("roll was not made")
	switch context.Field(profile.GENERATION_METHOD) {
	case method.Basic:
		code, err = basicAtmosphereRoll(context, dice)
		if err != nil {
			return code, fmt.Errorf("atmosphere roll failed: %v", err)
		}
	}
	return code, err
}

func basicAtmosphereRoll(context profile.Profile, dice *dice.Dicepool) (string, error) {

	prequisites, err := getBasicPrequisites(context, profile.KEY_Atmo)
	if err != nil {
		return "", fmt.Errorf("prequisites aquisition failed: %v", err)
	}
	sizeDM := prequisites[profile.KEY_Size]
	valid := false
	r := -1
	for !valid {
		r = dice.Sroll("2d6") - 7 + sizeDM
		switch sizeDM {
		case 0, 1:
			if r > 0 {
				continue
			}
		case 2, 3, 4, 5:
			if r > sizeDM+1 {
				continue
			}
		}
		valid = true
	}
	r = values.BoundInt(r, 0, 17)
	return ehex.ToCode(r), err
}
