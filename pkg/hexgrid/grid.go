package hexgrid

import "errors"

const (
	Q_odd  = 1
	Q_even = 2
	R_odd  = 3
	R_even = 4
)

type Grid struct {
	Orientation int
	Coordinates map[Coordinates]bool
}

func NewGrid(orientation int) (*Grid, error) {
	gr := Grid{}
	switch orientation {
	case Q_odd, Q_even, R_odd, R_even:
		gr.Orientation = orientation
	default:
		return nil, errors.New("invalid orientation")
	}
	gr.Coordinates = make(map[Coordinates]bool)
	return &gr, nil
}

type Coordinates struct {
	Orientation int
	X           int
	Y           int
	Q           int
	R           int
	S           int
}

type offset struct {
	r int
	c int
}

type hex struct {
	q int
	r int
	s int
}

/*
function cube_to_oddq(hex):
    var col = hex.q
    var row = hex.r + (hex.q - (hex.q&1)) / 2
    return OffsetCoord(col, row)

function oddq_to_cube(hex):
    var q = hex.col
    var r = hex.row - (hex.col - (hex.col&1)) / 2
    return Cube(q, r, -q-r)
*/
