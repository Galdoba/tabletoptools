package population

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/ehex"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method/basic"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

type Population struct {
	Code  string
	Value *int
	PCR   *int
}

func New() *Population {
	s := Population{}
	return &s
}

func (s *Population) Roll(context profile.Profile, dice *dice.Dicepool) error {
	gm := context.Field(profile.GENERATION_METHOD)
	switch gm {
	default:
		return fmt.Errorf("unknown method '%v'")
	case method.Basic:
		if s.Code != "" {
			return fmt.Errorf("basic roll failed: population was rolled")
		}
		code, err := basic.RollPopulation(context, dice)
		if err != nil {
			return fmt.Errorf("basic roll failed: %v", err)
		}
		s.Code = code
		v := ehex.ValueOf(s.Code)
		s.Value = &v
	}
	return nil
}

func (s *Population) Roll_PCR(ctx profile.Profile, dice *dice.Dicepool) error {
	gm := ctx.Field(profile.GENERATION_METHOD)
	switch gm {
	default:
		return fmt.Errorf("unknown method '%v'")
	case method.Continuation:
		if s.PCR != nil {
			return fmt.Errorf("continuation roll failed: pcr was rolled")
		}
		code, err := basic.RollPopulation(ctx, dice)
		if err != nil {
			return fmt.Errorf("basic roll failed: %v", err)
		}
		s.Code = code
		v := ehex.ValueOf(s.Code)
		s.Value = &v
	}
	return nil
}
