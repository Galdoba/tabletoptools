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
	Code        string
	Value       *int
	Temperature string
	MSTL        string
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
	if err := ctx.Inject(profile.KEY_Atmo, a.Code); err != nil {
		return err
	}

	if err := a.defineMinimumSustainableTL(ctx); err != nil {
		return err
	}
	if err := a.rollTemperature(ctx, dice); err != nil {
		return err
	}

	return nil
}

func (a *Atmosphere) defineMinimumSustainableTL(ctx profile.Profile) error {
	_, atmo, err := ctx.Ehex(profile.KEY_Atmo)
	if err != nil {
		return fmt.Errorf("failed to define MSTL: %v", err)
	}
	mstl := 0
	switch atmo {
	case 0, 1, 10, 15, 16, 17:
		mstl = 8
	case 2, 3, 13, 14:
		mstl = 5
	case 4, 7, 9:
		mstl = 3
	case 11:
		mstl = 9
	case 12:
		mstl = 10
	}
	a.MSTL = ehex.ToCode(mstl)
	return ctx.Inject(profile.KEY_Atmo_msTL, a.MSTL)
}

func (a *Atmosphere) rollTemperature(ctx profile.Profile, dice *dice.Dicepool) error {
	temp, err := basic.RollTemperature(ctx, dice)
	if err != nil {
		return fmt.Errorf("temperature roll failed: %v", err)
	}
	a.Temperature = temp
	return ctx.Inject(profile.KEY_Atmo_Temp, temp)
}
