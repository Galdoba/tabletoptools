package characteristic

import "github.com/Galdoba/tabletoptools/pkg/definition/ruleset"

type Value struct {
	Current int `json:"Current"` //текущее значение
	Max     int `json:"Maximum"` //максимальное значение для сущности
}

//Val - return maximum for THIS entity
func (v *Value) Val() int {
	return v.Current
}

//MaxVal - return maximum for THIS entity
func (v *Value) MaxVal() int {
	return v.Max
}

//for interface DM
//this.DM(rules)

//Mod - return value based on charactiristic current value and ruleset
func (v *Value) DM(rules int) int {
	switch rules {
	default:
		return 0
	case ruleset.HOSTILE:
		return modHOSTILE(v.Current)
	case ruleset.MGT2:
		return modMGT2(v.Current)
	}
}

func modMGT2(c int) int {
	switch c {
	case 1, 2:
		return -2
	case 3, 4, 5:
		return -1
	case 6, 7, 8:
		return 0
	case 9, 10, 11:
		return 1
	case 12, 13, 14:
		return 2
	default:
		if c <= 0 {
			return -3
		}
		return 3
	}
}

func modHOSTILE(c int) int {
	switch c {
	case 0, 1, 2:
		return -2
	case 3, 4, 5:
		return -1
	case 6, 7, 8:
		return 0
	case 9, 10, 11:
		return 1
	case 12, 13, 14:
		return 2
	default:
		if c < 0 {
			return -2
		}
		return 3
	}
}
