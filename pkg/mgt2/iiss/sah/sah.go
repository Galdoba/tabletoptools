package sah

import (
	"errors"
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/sah/atmo"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/sah/hydr"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/sah/size"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

type SAH struct {
	Size        *size.Size
	Atmosphere  *atmo.Atmosphere
	Hydrosphere *hydr.Hydrosphere
	Temperature string
	MSTL        string
}

func New() *SAH {
	sah := SAH{}
	sah.Size = size.New()
	sah.Atmosphere = atmo.New()
	sah.Hydrosphere = hydr.New()
	return &sah
}

func (sah *SAH) GenerateMissing(context profile.Profile, dice *dice.Dicepool) error {
	gm := context.Field(profile.GENERATION_METHOD)
	switch gm {
	case method.Basic:
		err := errors.New("no rolls wes made")
		funcs := basicRollFuncMap(sah)
		for _, key := range methodKeys(gm) {
			if err = funcs[key](context, dice); err != nil {
				return fmt.Errorf("%v generation failed: %v", key, err)
			}
		}
	case method.Continuation:
		err := errors.New("no rolls wes made")
		funcs := continuationRollFuncMap(sah)
		for _, key := range methodKeys(gm) {
			if err = funcs[key](context, dice); err != nil {
				return fmt.Errorf("%v generation failed: %v", key, err)
			}
		}

	}
	return nil
}

func methodKeys(genMethod string) []string {
	switch genMethod {
	case method.Basic:
		return basicKeys()
	case method.Continuation:
		return continuationKeys()
	}
	panic(fmt.Sprintf("method keys for '%v' not implemented"))
	return nil
}

func basicKeys() []string {
	return []string{profile.KEY_Size, profile.KEY_Atmo, profile.KEY_Hydr}
}

func continuationKeys() []string {
	return []string{profile.KEY_Size_Dkm}
}

func basicRollFuncMap(sah *SAH) map[string]func(profile.Profile, *dice.Dicepool) error {
	funcMap := make(map[string]func(profile.Profile, *dice.Dicepool) error)
	funcMap[profile.KEY_Size] = sah.Size.Roll
	funcMap[profile.KEY_Atmo] = sah.Atmosphere.Roll
	funcMap[profile.KEY_Hydr] = sah.Hydrosphere.Roll

	return funcMap
}

func continuationRollFuncMap(sah *SAH) map[string]func(profile.Profile, *dice.Dicepool) error {
	funcMap := make(map[string]func(profile.Profile, *dice.Dicepool) error)
	funcMap[profile.KEY_Size_Dkm] = sah.Size.Roll
	funcMap[profile.KEY_Atmo] = sah.Atmosphere.Roll
	funcMap[profile.KEY_Hydr] = sah.Hydrosphere.Roll

	return funcMap
}

func ValueOf(sah *SAH, key string) string {
	switch key {
	case profile.KEY_Size:
		return fmt.Sprintf("%v", sah.Size.Code)
	case profile.KEY_Size_Dkm:
		return fmt.Sprintf("%v", sah.Size.DiameterKm)
	case profile.KEY_Size_D:
		return fmt.Sprintf("%v", sah.Size.Density)
	case profile.KEY_Size_G:
		return fmt.Sprintf("%v", sah.Size.Gravity)
	case profile.KEY_Size_M:
		return fmt.Sprintf("%v", sah.Size.Mass)
	case profile.KEY_Atmo:
		return fmt.Sprintf("%v", sah.Atmosphere.Code)
	case profile.KEY_Atmo_msTL:
		return fmt.Sprintf("%v", sah.Atmosphere.Temperature)
	case profile.KEY_Hydr:
		return fmt.Sprintf("%v", sah.Hydrosphere.Code)
	}
	return ""
}
