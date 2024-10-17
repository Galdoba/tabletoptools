package social

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/social/goverment"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/social/lawlevel"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/social/population"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/social/starport"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/iiss/social/techlevel"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/tradecode"
)

type Social struct {
	Population *population.Population
	Goverment  *goverment.Goverment
	LawLevel   *lawlevel.LawLevel
	Starport   *starport.Starport
	TechLevel  *techlevel.TechLevel
	TradeCodes *[]string
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
	funcMap := nilRollFuncMap(social)
	switch gm {
	case method.Basic:
		funcMap = basicRollFuncMap(social)
	case method.Continuation:
		funcMap = continuationRollFuncMap(social)
	}
	err := errors.New("no rolls was made")
	for _, key := range methodKeys(gm) {
		if err = funcMap[key](context, dice); err != nil {
			return fmt.Errorf("%v generation failed: %v", key, err)
		}
		context.Inject(key, ValueOf(social, key))
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
	return []string{profile.KEY_Pops, profile.KEY_Govr, profile.KEY_Laws, profile.KEY_Port, profile.KEY_TL}
}

func continuationKeys() []string {
	return []string{profile.KEY_TC}
}

func nilRollFuncMap(social *Social) map[string]func(profile.Profile, *dice.Dicepool) error {
	funcMap := make(map[string]func(profile.Profile, *dice.Dicepool) error)
	return funcMap
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

func continuationRollFuncMap(social *Social) map[string]func(profile.Profile, *dice.Dicepool) error {
	funcMap := make(map[string]func(profile.Profile, *dice.Dicepool) error)
	funcMap[profile.KEY_TC] = social.determineTradeCodes

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
	case profile.KEY_TC:
		return fmt.Sprintf("%v", strings.Join(*social.TradeCodes, " "))

	}
	return ""
}

func (soc *Social) determineTradeCodes(ctx profile.Profile, dice *dice.Dicepool) error {
	if soc.TradeCodes != nil {
		return fmt.Errorf("trade codes were determined")
	}
	tc := tradecode.TradeCodes(ctx)
	soc.TradeCodes = &tc
	return nil
}
