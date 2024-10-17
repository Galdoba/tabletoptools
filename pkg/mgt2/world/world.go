package world

import (
	"fmt"
	"time"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/sah"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/social"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/systemorbit"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/tradecode"
)

const (
	genarationMethod = "generation_method"
)

type World struct {
	UID          int64
	Name         string
	Profile      profile.Profile
	OrbitalData  *systemorbit.Orbit
	PhysicalData *sah.SAH
	SocialData   *social.Social
}

func New(name string) *World {
	w := World{}
	w.UID = time.Now().UnixNano()
	w.Name = name
	w.Profile = profile.New()
	w.OrbitalData = systemorbit.New()
	w.PhysicalData = sah.New()
	w.SocialData = social.New()
	return &w
}

func (w *World) SetGenerationMethod(m string) error {
	w.Profile.Delete(genarationMethod)
	switch m {
	default:
		return fmt.Errorf("generation_method = %v: unimplemented", m)
	case method.Basic, method.Continuation:
		return w.Profile.Inject(genarationMethod, m)
	}
}

func (w *World) GenerateMissingData(dice *dice.Dicepool) error {
	genM := w.Profile.Field(profile.GENERATION_METHOD)
	switch genM {
	default:
		return fmt.Errorf("mathod '%v' unimplemented", genM)
	case method.Basic:
		if err := w.OrbitalData.GenerateMissing(w.Profile, dice); err != nil {
			return err
		}
		if err := w.PhysicalData.GenerateMissing(w.Profile, dice); err != nil {
			return err
		}
		if err := w.SocialData.GenerateMissing(w.Profile, dice); err != nil {
			return err
		}
		tc := tradecode.TradeCodes(w.Profile)
		if err := tradecode.Inject(w.Profile, tc...); err != nil {
			return err
		}
		w.SetGenerationMethod(method.Continuation)
	case method.Continuation:
		if err := w.OrbitalData.GenerateMissing(w.Profile, dice); err != nil {
			return err
		}
		if err := w.PhysicalData.GenerateMissing(w.Profile, dice); err != nil {
			return err
		}
	}
	return nil
}

type basicGeneration struct {
	function func(string, *dice.Dicepool) error
}

func BasicGenerationMethods() []func(string, *dice.Dicepool) error {
	return nil
}
