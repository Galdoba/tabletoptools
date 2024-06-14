package dice

import (
	"sort"
)

type DiceT5 interface {
	Roll_1D() int
	Roll_2D() int
	Roll_3D() int
	Roll_4D() int
	Roll_5D() int
	Roll_6D() int
	Roll_7D() int
	Roll_8D() int
	Roll_9D() int
	Roll_10D() int
	Flux() int
	Flux_Positive() int
	Flux_Negative() int
	Result() []int
}

func NewT5(seed string) DiceT5 {
	d := New().SetSeed(seed)
	return d
}

func (d *Dicepool) Roll_1D() int {
	return d.Roll("1d6").Sum()
}

func (d *Dicepool) Roll_2D() int {
	return d.Roll("2d6").Sum()
}

func (d *Dicepool) Roll_3D() int {
	return d.Roll("3d6").Sum()
}

func (d *Dicepool) Roll_4D() int {
	return d.Roll("4d6").Sum()
}

func (d *Dicepool) Roll_5D() int {
	return d.Roll("5d6").Sum()
}

func (d *Dicepool) Roll_6D() int {
	return d.Roll("6d6").Sum()
}

func (d *Dicepool) Roll_7D() int {
	return d.Roll("7d6").Sum()
}

func (d *Dicepool) Roll_8D() int {
	return d.Roll("8d6").Sum()
}

func (d *Dicepool) Roll_9D() int {
	return d.Roll("9d6").Sum()
}

func (d *Dicepool) Roll_10D() int {
	return d.Roll("10d6").Sum()
}

func (d *Dicepool) Flux() int {
	r1 := d.Roll_1D()
	r2 := d.Roll_1D()
	return r1 - r2
}

func (d *Dicepool) Flux_Positive() int {
	res := append([]int{}, d.Roll_1D())
	res = append(res, d.Roll_1D())
	sort.Ints(res)
	return res[1] - res[0]
}

func (d *Dicepool) Flux_Negative() int {
	res := append([]int{}, d.Roll_1D())
	res = append(res, d.Roll_1D())
	sort.Ints(res)
	return res[0] - res[1]
}

/*
effect := roll.Two()

*/
