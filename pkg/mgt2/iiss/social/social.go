package social

import (
	"errors"
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/social/goverment"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/social/lawlevel"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/social/population"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/social/starport"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/social/techlevel"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

type Social struct {
	Population *population.Population
	Goverment  *goverment.Goverment
	LawLevel   *lawlevel.LawLevel
	Starport   *starport.Starport
	TechLevel  *techlevel.TechLevel
}

func New() *Social {
	social := Social{}
	social.Population = population.New()
	social.Goverment = goverment.New()
	social.LawLevel = lawlevel.New()
	social.Starport = starport.New()
	social.TechLevel = techlevel.New()
	return &social
}

func (social *Social) GenerateMissing(context profile.Profile, dice *dice.Dicepool) error {
	gm := context.Field(profile.GENERATION_METHOD)
	switch gm {
	case method.Basic:
		err := errors.New("no rolls wes made")
		funcs := basicRollFuncMap(social)
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
			context.Inject(key, ValueOf(social, key))
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
	return []string{profile.KEY_Pops, profile.KEY_Govr, profile.KEY_Laws, profile.KEY_Port, profile.KEY_TL} //, profile.KEY_Atmo, profile.KEY_Temperature, profile.KEY_Hydr}
}

func basicRollFuncMap(social *Social) map[string]func(profile.Profile, *dice.Dicepool) error {
	funcMap := make(map[string]func(profile.Profile, *dice.Dicepool) error)
	funcMap[profile.KEY_Pops] = social.Population.Roll
	funcMap[profile.KEY_Govr] = social.Goverment.Roll
	funcMap[profile.KEY_Laws] = social.LawLevel.Roll
	funcMap[profile.KEY_Port] = social.Starport.Roll
	funcMap[profile.KEY_TL] = social.TechLevel.Roll

	return funcMap
}

func ValueOf(social *Social, key string) string {
	switch key {
	case profile.KEY_Pops:
		return fmt.Sprintf("%v", social.Population.Code)
	case profile.KEY_Govr:
		return fmt.Sprintf("%v", social.Goverment.Code)
	case profile.KEY_Laws:
		return fmt.Sprintf("%v", social.LawLevel.Code)
	case profile.KEY_Port:
		return fmt.Sprintf("%v", social.Starport.Code)
	case profile.KEY_TL:
		return fmt.Sprintf("%v", social.TechLevel.Code)
		// case profile.KEY_Size_Dkm:
		// 	return fmt.Sprintf("%v", social.Size.Diemeter)
		// case profile.KEY_Size_D:
		// 	return fmt.Sprintf("%v", social.Size.Density)
		// case profile.KEY_Size_G:
		// 	return fmt.Sprintf("%v", social.Size.Gravity)
		// case profile.KEY_Size_M:
		// 	return fmt.Sprintf("%v", social.Size.Mass)
		// case profile.KEY_Atmo:
		// 	return fmt.Sprintf("%v", social.Atmosphere.Code)
		// case profile.KEY_Temperature:
		// 	return fmt.Sprintf("%v", social.Temperature)
		// case profile.KEY_Hydr:
		// 	return fmt.Sprintf("%v", social.Hydrosphere.Code)
	}
	return ""
}
