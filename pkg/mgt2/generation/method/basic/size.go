package basic

import (
	"errors"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/ehex"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/values"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

func RollSize(context profile.Profile, dice *dice.Dicepool) (string, error) {
	code := ""
	err := errors.New("roll was not made")
	switch context.Field(profile.GENERATION_METHOD) {
	case method.Basic:
		code, err = basicSizeRoll(context, dice)
	}
	return code, err
}

func basicSizeRoll(context profile.Profile, dice *dice.Dicepool) (string, error) {
	r := dice.Sroll("2d6-2")
	for r > 9 {
		r2 := dice.Sroll("1d6")
		if r2 < 6 {
			break
		}
		r++
	}
	r = values.BoundInt(r, 0, 15)
	return ehex.ToCode(r), nil
}
