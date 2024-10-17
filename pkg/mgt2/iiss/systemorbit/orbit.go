package systemorbit

import (
	"errors"
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

type Orbit struct {
	AU        *float64
	HZCO      *float64
	TidalLock bool
}

type OrbitalData interface {
	Zone() string
}

func New() *Orbit {
	o := Orbit{}
	return &o
}

func (o *Orbit) GenerateMissing(context profile.Profile, dice *dice.Dicepool) error {
	gm := context.Field(profile.GENERATION_METHOD)
	switch gm {
	case method.Basic:
		err := errors.New("no rolls wes made")
		funcs := basicRollFuncMap(o)
		for _, key := range methodKeys(gm) {
			// data := context.Field(key)
			// if data != "" {
			// 	continue
			// }
			// fn := funcs[key]
			// if fn == nil {
			// 	fmt.Println("not implemented for key", key)
			// 	continue
			// }
			if err = funcs[key](context, dice); err != nil {
				return fmt.Errorf("%v generation failed: %v", key, err)
			}
			//context.Inject(key, ValueOf(sah, key))
		}

	}
	return nil
}

func methodKeys(genMethod string) []string {
	switch genMethod {
	case method.Basic:
		return basicKeys()
	}
	panic(fmt.Sprintf("method keys for '%v' not implemented"))
	return nil
}

func basicKeys() []string {
	return []string{profile.KEY_Tidal}
}

func basicRollFuncMap(o *Orbit) map[string]func(profile.Profile, *dice.Dicepool) error {
	funcMap := make(map[string]func(profile.Profile, *dice.Dicepool) error)
	funcMap[profile.KEY_Tidal] = o.rollTidalLock

	return funcMap
}

func (o *Orbit) rollTidalLock(ctx profile.Profile, dice *dice.Dicepool) error {
	gm := ctx.Field(profile.GENERATION_METHOD)
	switch gm {
	default:
		return fmt.Errorf("unknown method '%v'")
	case method.Basic:
		r := dice.Sroll("2d6")
		switch r {
		case 12:
			o.TidalLock = true
			if err := ctx.Inject(profile.KEY_Tidal, "Y"); err != nil {
				return fmt.Errorf("basic tidal lock roll failed: %v", err)
			}
		default:
			if err := ctx.Inject(profile.KEY_Tidal, "N"); err != nil {
				return fmt.Errorf("basic tidal lock roll failed: %v", err)
			}
		}

	}
	return nil
}
