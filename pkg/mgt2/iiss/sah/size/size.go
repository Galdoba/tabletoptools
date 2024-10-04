package size

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/ehex"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method/basic"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method/continuation"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

type Size struct {
	Code       string
	Value      *int
	DiameterKm *int
	Density    *float64
	Gravity    *float64
	Mass       *float64
}

func New() *Size {
	s := Size{}
	return &s
}

func (s *Size) Roll(ctx profile.Profile, dice *dice.Dicepool) error {
	gm := ctx.Field(profile.GENERATION_METHOD)
	switch gm {
	default:
		return fmt.Errorf("unknown method '%v'")
	case method.Basic:
		if s.Code != "" {
			return fmt.Errorf("basic roll failed: size was rolled")
		}
		code, err := basic.RollSize(ctx, dice)
		if err != nil {
			return fmt.Errorf("basic roll failed: %v", err)
		}
		s.Code = code
		v := ehex.ValueOf(s.Code)
		s.Value = &v
	case method.Continuation:
		fmt.Println("continuation: xxxxxxxxxxxxxxxxxxxxxx")
		if s.DiameterKm == nil {
			dkm, err := continuation.RollDiameterKm(ctx, dice)
			if err != nil {
				return fmt.Errorf("continuation roll failed: %v", err)
			}
			s.DiameterKm = &dkm
		}

	}
	return nil
}
