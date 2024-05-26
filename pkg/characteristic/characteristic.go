package characteristic

import (
	"fmt"
	"strings"
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
	switch name {
	default:
		return ch, fmt.Errorf("undefined name '%v'", name)
	case Strength:
		ch.Name = name
		ch.Position = nameToPos(name)
		ch.Abbreviation = STR
		ch.Type = PHYSICAL
	case Dexterity:
	case Agility:
	case Grace:
	case Endurance:
	case Stamina:
	case Vigor:
	case Intelligence:
	case Education:
	case Training:
	case Instinct:
	case Social_Standing:
	case Charisma:
	case Caste:
	case Psi:
	case Sanity:
	case Wealth:
	case Luck:
	case Morale:
	}
	return ch, nil
}

type Set struct {
	CHR map[Characteristic]*Value
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
