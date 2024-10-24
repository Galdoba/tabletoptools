package stellarhex

import "github.com/Galdoba/tabletoptools/pkg/dice"

const (
	Undefined = iota
	ExtraGalactic
	Rift
	Sparse
	Scattered
	Standard
	Dense
	Cluster
	Core
)

type Hex struct {
	Density int
}

func NewHex(dice *dice.Dicepool) *Hex {
	hx := Hex{}
	return &hx
}
