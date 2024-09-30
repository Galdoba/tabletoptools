package characteristic

import "fmt"

const (
	// STR      = "Strenght"
	// DEX      = "Dexterity"
	// END      = "Endurance"
	// INT      = "Intelligence"
	// EDU      = "Education"
	// SOC      = "Social Standing"
	// WLT      = "Wealth"
	// LCK      = "Luck"
	// MRL      = "Morale"
	// STY      = "Sanity"
	// PSY      = "Psionics"
	// INS      = "Instinct"
	// PCK      = "Pack"
	// TER      = "Territory"
	// CHA      = "Charisma"
	physical = "physical"
	mental   = "mental"
	obscure  = "obscure"
	// C1       = "C1"
	// C2       = "C2"
	// C3       = "C3"
	// C4       = "C4"
	// C5       = "C5"
	// C6       = "C6"
	// CS       = "CS"
	// CP       = "CP"
	// CM       = "CM"
	// CW       = "CW"
	// CL       = "CL"
	// CO       = "CO"
)

type characteristic struct {
	name           string
	abb            string
	code           string
	cType          string
	effectiveScore int
	maxScore       int
	creationMod    int
	creationDice   int
}

func New(name string, options ...CharacteristicOption) (*characteristic, error) {
	ch := characteristic{}
	ch.name = name
	ch.abb, ch.code, ch.cType = abbCodeType(name)
	if ch.abb == "invalid" {
		return nil, fmt.Errorf("invalid input: name '%v' is not implemented", name)
	}

	stats := defaultOptions()
	for _, change := range options {
		change(&stats)
	}
	ch.effectiveScore = stats.effectiveScore
	ch.maxScore = stats.maxScore
	ch.creationDice = stats.creationDice
	ch.creationMod = stats.creationMod

	if ch.maxScore < 0 {
		return nil, fmt.Errorf("invalid input: max score is less than zero")
	}
	if ch.maxScore > unmodifiedLimit(ch.creationDice) {
		return nil, fmt.Errorf("invalid input: max score is more than unmodified limit")
	}
	if ch.effectiveScore > ch.maxScore {
		return nil, fmt.Errorf("invalid input: max score is less than effective score")
	}

	return &ch, nil
}

func unmodifiedLimit(diceNum int) int {
	limitAddition := 0
	for i := 1; i <= diceNum; i++ {
		switch i {
		case 1:
			limitAddition = 1
		default:
			limitAddition = limitAddition + (limitAddition * i)
		}
	}
	return diceNum*6 + limitAddition
}

type DiceRoller interface {
	Sroll(string) int
}

func IsActive(ch *characteristic) bool {
	if ch.maxScore != 0 {
		return true
	}
	return false
}

func (ch *characteristic) Roll(dice DiceRoller) error {
	if ch.maxScore != 0 {
		return fmt.Errorf("cannot roll active characteristic")
	}
	code := fmt.Sprintf("%vd6", ch.creationDice)
	r := dice.Sroll(code)
	r += ch.creationMod
	if r < 1 {
		r = 1
	}

	limit := unmodifiedLimit(ch.creationDice)
	if r > limit {
		r = limit
	}
	ch.maxScore = r
	ch.effectiveScore = r
	return nil
}

func (ch *characteristic) Score() int {
	return ch.effectiveScore
}
