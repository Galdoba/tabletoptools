package lawlevel

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/ehex"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method/basic"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

type LawLevel struct {
	Code  string
	Value *int
}

func New() *LawLevel {
	s := LawLevel{}
	return &s
}

func (a *LawLevel) Roll(ctx profile.Profile, dice *dice.Dicepool) error {
	gm := ctx.Field(profile.GENERATION_METHOD)
	switch gm {
	default:
		return fmt.Errorf("unknown method '%v'")
	case method.Basic:
		if a.Code != "" {
			return fmt.Errorf("basic roll failed: law level was rolled")
		}
		code, err := basic.RollLawLevel(ctx, dice)
		if err != nil {
			return fmt.Errorf("basic roll failed: %v", err)
		}
		a.Code = code
		v := ehex.ValueOf(a.Code)
		a.Value = &v
	}

	return nil
}
