package hydr

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/ehex"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method/basic"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

type Hydrosphere struct {
	Code  string
	Value *int
}

func New() *Hydrosphere {
	s := Hydrosphere{}
	return &s
}

func (a *Hydrosphere) Roll(context profile.Profile, dice *dice.Dicepool) error {
	gm := context.Field(profile.GENERATION_METHOD)
	switch gm {
	default:
		return fmt.Errorf("unknown method '%v'")
	case method.Basic:
		if a.Code != "" {
			return fmt.Errorf("basic roll failed: hydrosphere was rolled")
		}
		code, err := basic.RollHydrosphere(context, dice)
		if err != nil {
			return fmt.Errorf("basic roll failed: %v", err)
		}
		a.Code = code
		v := ehex.ValueOf(a.Code)
		a.Value = &v
	}
	return context.Inject(profile.KEY_Hydr, a.Code)
}
