package size

import (
	"fmt"
	"regexp"
	"strconv"
)

type Size struct {
	Code       string
	DiameterKm int
	Density    float64
	Gravity    float64
	Mass       float64
}

var anyEhex = "([0123456789ABCDEFGHJKLMNPQRSTUVWXYZ])"

var intVal = `([0123456789]+)`
var floatVal = `([0123456789]+[.]?[0123456789]+)`

func FromProfile(s string) (*Size, error) {

	reString := fmt.Sprintf(`%v-%vkm-%v-%v-%v`, anyEhex, intVal, floatVal, floatVal, floatVal)
	re := regexp.MustCompile(reString)
	subs := re.FindStringSubmatch(s)
	sz := Size{}
	fmt.Println(subs)
	for i, m := range subs {
		fmt.Println(i, m)
	}
	for i, match := range subs {
		switch i {
		case 1:
			sz.Code = match
		case 2:
			v, err := strconv.Atoi(match)
			if err != nil {
				return nil, fmt.Errorf("failed to parse Diameter: %v", err)
			}
			sz.DiameterKm = v
		case 3:
			v, err := strconv.ParseFloat(match, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse Density: %v", err)
			}
			sz.Density = v
		case 4:
			v, err := strconv.ParseFloat(match, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse Gravity: %v", err)
			}
			sz.Gravity = v
		case 5:
			v, err := strconv.ParseFloat(match, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse Mass: %v", err)
			}
			sz.Mass = v
		}
	}

	return &sz, nil
}

// S-Dkm-D-G-M
// 5-8163-1.03-0.66-0.27
func (sz *Size) Profile() string {
	return fmt.Sprintf("%v-%vkm-%v-%v-%v",
		fmt.Sprintf("%v", sz.Code),
		fmt.Sprintf("%v", sz.DiameterKm),
		fmt.Sprintf("%v", sz.Density),
		fmt.Sprintf("%v", sz.Gravity),
		fmt.Sprintf("%v", sz.Mass),
	)
}
