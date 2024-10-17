package basic

import (
	"errors"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/ehex"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/values"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

func RollSize(ctx profile.Profile, dice *dice.Dicepool) (string, error) {
	code := ""
	err := errors.New("roll was not made")
	switch ctx.Field(profile.GENERATION_METHOD) {
	case method.Basic:
		code, err = basicSizeRoll(ctx, dice)
	}
	if err := ctx.Inject(profile.KEY_Size, code); err != nil {
		return "", nil
	}
	return code, err
}

func basicSizeRoll(ctx profile.Profile, dice *dice.Dicepool) (string, error) {
	r := dice.Sroll("2d6-2")
	for r > 9 {
		r2 := dice.Sroll("1d6")
		if r2 < 4 {
			break
		}
		r++
	}
	r = values.BoundInt(r, 0, 15)
	return ehex.ToCode(r), nil
}
