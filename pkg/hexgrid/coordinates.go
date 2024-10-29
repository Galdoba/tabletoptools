package hexgrid

const (
	Axix_Q = "Q"
	Axix_R = "R"
	Axix_S = "S"
	Axix_X = "X"
	Axix_Y = "Y"
	Axix_Row = "Row"
	Axix_Col = "Col"
)

type Coordinates struct {
	axis map[string]*float64
}

func NewCoordinates() *Coordinates {
	crd := make(map[string]*float64)
	return &Coordinates{crd}
}

func (crd *Coordinates) SetAxis(axis string, value *float64) *Coordinates {
	crd.axis[axis] = value
	return crd
}

func Cube(crd *Coordinates) *cube {
	cb := cube{}
	for _, axis := range []string{} {
		
	}

}
