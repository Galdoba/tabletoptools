package systemorbit

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method/basic"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

type Orbit struct {
	AU          *float64
	HZCO        *float64
	Temperature string
}

type OrbitalData interface {
	Zone() string
}

func New() *Orbit {
	o := Orbit{}
	return &o
}

func (o *Orbit) Roll(ctx profile.Profile, dice *dice.Dicepool) error {
	gm := ctx.Field(profile.GENERATION_METHOD)
	switch gm {
	default:
		return fmt.Errorf("unknown method '%v'")
	case method.Basic:
		if o.Temperature != "" {
			return fmt.Errorf("basic roll failed: hydrosphere was rolled")
		}
		code, err := basic.RollTemperature(ctx, dice)
		if err != nil {
			return fmt.Errorf("basic roll failed: %v", err)
		}
		o.Temperature = code

	}
	return nil
}
