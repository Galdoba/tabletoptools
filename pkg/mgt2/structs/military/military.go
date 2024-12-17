package military

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/Galdoba/tabletoptools/pkg/ehex"
)

type Military struct {
	Enforcement   int
	Militia       int
	Army          int
	WetNavy       int
	AirForce      int
	SystemDefence int
	Navy          int
	Marines       int
	Budget        float64
}

var anyEhex = "([0123456789ABCDEFGHJKLMNPQRSTUVWXYZ])"

var budget = `([0123456789].[0123456789][0123456789])`

func FromProfile(s string) (*Military, error) {
	reString := fmt.Sprintf(`%v%v%v%v%v-%v%v%v:%v`, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, budget)
	re := regexp.MustCompile(reString)
	subs := re.FindStringSubmatch(s)
	mil := Military{}

	for i, match := range subs {
		switch i {
		case 1:
			mil.Enforcement = ehex.ValueOf(match)
		case 2:
			mil.Militia = ehex.ValueOf(match)
		case 3:
			mil.Army = ehex.ValueOf(match)
		case 4:
			mil.WetNavy = ehex.ValueOf(match)
		case 5:
			mil.AirForce = ehex.ValueOf(match)
		case 6:
			mil.SystemDefence = ehex.ValueOf(match)
		case 7:
			mil.Navy = ehex.ValueOf(match)
		case 8:
			mil.Marines = ehex.ValueOf(match)
		case 9:
			fl, err := strconv.ParseFloat(match, 64)
			if err != nil {
				return nil, err
			}
			mil.Budget = fl
		}
	}

	return &mil, nil
}

func (mil *Military) Profile() string {
	return fmt.Sprintf("%v%v%v%v%v-%v%v%v:%v",
		fmt.Sprintf("%v", ehex.ToCode(mil.Enforcement)),
		fmt.Sprintf("%v", ehex.ToCode(mil.Militia)),
		fmt.Sprintf("%v", ehex.ToCode(mil.Army)),
		fmt.Sprintf("%v", ehex.ToCode(mil.WetNavy)),
		fmt.Sprintf("%v", ehex.ToCode(mil.AirForce)),
		fmt.Sprintf("%v", ehex.ToCode(mil.SystemDefence)),
		fmt.Sprintf("%v", ehex.ToCode(mil.Navy)),
		fmt.Sprintf("%v", ehex.ToCode(mil.Marines)),
		fmt.Sprintf("%0.2f", mil.Budget),
	) + "%"
}

func presense(p bool) string {
	if p {
		return "Y"
	}
	return "N"
}

func importance(i int) string {
	s := fmt.Sprintf("%v", i)
	if i > -1 {
		s = "+" + s
	}
	return s
}
