package stellarhex

import (
	"errors"

	"github.com/Galdoba/tabletoptools/pkg/dice"
)

// RollPresence - presence roll for hex
func RollPresence(density int, dice *dice.Dicepool) (bool, error) {
	code := ""
	tn := 0
	switch density {
	default:
		return false, errors.New("density unknown or undefined")
	case ExtraGalactic:
		code = "3d6"
		tn = 3
	case Rift:
		code = "2d6"
		tn = 2
	case Sparse:
		code = "1d6"
		tn = 1
	case Scattered:
		code = "1d6"
		tn = 2
	case Standard:
		code = "1d6"
		tn = 3
	case Dense:
		code = "1d6"
		tn = 4
	case Cluster:
		code = "1d6"
		tn = 5
	case Core:
		code = "2d6"
		tn = 11
	}
	return dice.Sroll(code) <= tn, nil
}
