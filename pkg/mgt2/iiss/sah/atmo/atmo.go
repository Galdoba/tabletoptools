package atmo

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/ehex"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method/basic"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

type Atmosphere struct {
	Code  string
	Value *int
}

func New() *Atmosphere {
	s := Atmosphere{}
	return &s
}

func (a *Atmosphere) Roll(ctx profile.Profile, dice *dice.Dicepool) error {
	gm := ctx.Field(profile.GENERATION_METHOD)
	switch gm {
	default:
		return fmt.Errorf("unknown method '%v'")
	case method.Basic:
		if a.Code != "" {
			return fmt.Errorf("basic roll failed: atmosphere was rolled")
		}
		code, err := basic.RollAtmosphere(ctx, dice)
		if err != nil {
			return fmt.Errorf("basic roll failed: %v", err)
		}
		a.Code = code
		v := ehex.ValueOf(a.Code)
		a.Value = &v
	}

	return nil
}
