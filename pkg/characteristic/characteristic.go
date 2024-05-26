package characteristic

import (
	"fmt"
	"strings"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/options"
)

type Characteristic struct {
	//Value          int //пойдет в значение
	//GeneticValue   string//пойдет в значение
	Name         string `json:"Name"`
	Position     string `json:"Position"`
	Abbreviation string `json:"Abbreviation"`
	Type         string `json:"Type"`
	//GeneticCode  string `json:"Genetic Code,omitempty"`

}

func defaultChar(name string) (Characteristic, error) {
	ch := Characteristic{}
	ch.Name = name
	ch.Position = nameToPos(name)
	switch name {
	default:
		return ch, fmt.Errorf("undefined name '%v'", name)
	case Strength:
		ch.Abbreviation = STR
		ch.Type = PHYSICAL
	case Dexterity:
		ch.Abbreviation = DEX
		ch.Type = PHYSICAL
	case Agility:
		ch.Abbreviation = AGI
		ch.Type = PHYSICAL
	case Grace:
		ch.Abbreviation = GRA
		ch.Type = PHYSICAL
	case Endurance:
		ch.Abbreviation = END
		ch.Type = PHYSICAL
	case Stamina:
		ch.Abbreviation = STA
		ch.Type = PHYSICAL
	case Vigor:
		ch.Abbreviation = VIG
		ch.Type = PHYSICAL
	case Intelligence:
		ch.Abbreviation = INT
		ch.Type = MENTAL
	case Education:
		ch.Abbreviation = EDU
		ch.Type = MENTAL
	case Training:
		ch.Abbreviation = TRA
		ch.Type = MENTAL
	case Instinct:
		ch.Abbreviation = INS
		ch.Type = MENTAL
	case Social_Standing:
		ch.Abbreviation = SOC
		ch.Type = SOCIAL
	case Charisma:
		ch.Abbreviation = CHA
		ch.Type = SOCIAL
	case Caste:
		ch.Abbreviation = CAS
		ch.Type = SOCIAL
	case Psi:
		ch.Abbreviation = PSI
		ch.Type = OBSCURE
	case Sanity:
		ch.Abbreviation = SAN
		ch.Type = OBSCURE
	case Wealth:
		ch.Abbreviation = WLT
		ch.Type = OBSCURE
	case Luck:
		ch.Abbreviation = LCK
		ch.Type = OBSCURE
	case Morale:
		ch.Abbreviation = MOR
		ch.Type = OBSCURE
	}

	return ch, nil
}

type Set struct {
	CHR map[Characteristic]*Value
}

func NewCharSet(codes ...string) (*Set, error) {
	cs := Set{}
	cs.CHR = map[Characteristic]*Value{}
	for _, code := range codes {
		ch, err := defaultChar(code)
		if err != nil {
			return nil, fmt.Errorf("default char creation: %v", code)
		}
		cs.CHR[ch] = &Value{}
	}
	posMap := make(map[string]int)
	for ch := range cs.CHR {
		posMap[ch.Position]++
	}
	for k, v := range posMap {
		if v > 1 {
			return nil, fmt.Errorf("characteristic position '%v' met %v times", k, v)
		}
	}
	return &cs, nil
}

func (cs *Set) ByCode(code string) (Characteristic, *Value, error) {
	for k, v := range cs.CHR {
		if k.Position == code {
			return k, v, nil
		}
		if k.Name == code {
			return k, v, nil
		}
		if k.Abbreviation == code {
			return k, v, nil
		}
		if k.Position == nameToPos(code) {
			return k, v, nil
		}
	}
	return Characteristic{}, nil, fmt.Errorf("no charactiristic by code '%v'", code)
}

func nameToPos(name string) string {
	switch strings.ToUpper(name) {
	case strings.ToUpper(Strength), STR:
		return "C1"
	case strings.ToUpper(Dexterity), DEX:
		return "C2"
	case strings.ToUpper(Agility), AGI:
		return "C2"
	case strings.ToUpper(Grace), GRA:
		return "C2"
	case strings.ToUpper(Endurance), END:
		return "C3"
	case strings.ToUpper(Stamina), STA:
		return "C3"
	case strings.ToUpper(Vigor), VIG:
		return "C3"
	case strings.ToUpper(Intelligence), INT:
		return "C4"
	case strings.ToUpper(Education), EDU:
		return "C5"
	case strings.ToUpper(Training), TRA:
		return "C5"
	case strings.ToUpper(Instinct), INS:
		return "C5"
	case strings.ToUpper(Social_Standing), SOC:
		return "C6"
	case strings.ToUpper(Charisma), CHA:
		return "C6"
	case strings.ToUpper(Caste), CAS:
		return "C6"
	case strings.ToUpper(Territory), TER:
		return "R1"
	case strings.ToUpper(Psi), PSI:
		return "CP"
	case strings.ToUpper(Sanity), SAN:
		return "CS"
	// case strings.ToUpper(Wealth), WLT:
	// 	return "CW"
	// case strings.ToUpper(Luck), LCK:
	// 	return "CL"
	// case strings.ToUpper(Morale), MOR:
	// 	return "CM"
	default:
		return "obscure"
	}

}

func AbbByRace(r string) []string {
	switch r {
	default:
		return []string{"NO SET"}
	case "Human":
		return []string{STR, DEX, END, INT, EDU, SOC}
	}
}

func (cs *Set) Roll(dice *dice.Dicepool, opts ...options.Option) error {
	for k := range cs.CHR {
		setAs := 0
		for _, o := range opts {
			if k.Name == o.Key || k.Abbreviation == o.Key {
				val, ok := o.Val.(int)
				if !ok {
					return fmt.Errorf("roll characteristics: '%v' is not int", o.Key)
				}
				setAs = val
			}
		}
		if setAs == 0 {
			setAs = dice.Sroll("2d6")
		}
		chVal := &Value{}
		chVal.Current = setAs
		chVal.Max = setAs
		cs.CHR[k] = chVal
	}
	return nil
}
