package starsystem

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type StarSystem struct {
	Stars []*Star
	Age   float64
}

type Star struct {
	Type         string
	SubType      string
	Class        string
	Mass         float64
	Diameter     float64
	Luminosity   float64
	Age          float64
	Designation  string
	OrbitN       float64
	Eccentricity float64
}

var starClass = "(Ia|Ib|II|III|IV|V|VI|BD|D)"
var starType = "(O|B|A|F|G|K|M|L|T|Y|)"
var starTSTconnector = "( )?"
var starSubType = "([0123456789])"
var typeFloat = `(-[0123456789](\.)?([0123456789])?([0123456789])?([0123456789])?([0123456789])?([0123456789])?([0123456789])?)?`
var starAge = `([0123456789](\.)?([0123456789])?([0123456789])?([0123456789])?([0123456789])?([0123456789])?([0123456789])?)?`
var starDesignation = `((:)?(A|Aa|Ab|Aab|B|Ba|Bb|Bab|C|Ca|Cb|Cab|D|Da|Db|Dab|))?`

var starProfile = fmt.Sprintf(`%v%v%v%v%v%v%v%v%v%v%v`, starType, starSubType, starTSTconnector, starClass, typeFloat, typeFloat, typeFloat, typeFloat, starDesignation, typeFloat, typeFloat)

var anyEhex = "([0123456789ABCDEFGHJKLMNPQRSTUVWXYZ])"

func FromProfile(s string) (*Star, error) {

	return nil, nil
}

func parseStar(s string) (*Star, error) {
	reString := starProfile
	re := regexp.MustCompile(reString)
	parsed := re.FindString(s)
	if parsed == "" {
		return nil, fmt.Errorf("00failed to parse starsystem from '%v'", s)
	}
	st := Star{}
	for i, mask := range []string{starType, starSubType, starTSTconnector, starClass, typeFloat, typeFloat, typeFloat, typeFloat, starDesignation, typeFloat, typeFloat} {
		switch mask {
		case starTSTconnector:
			_, parsed = parseString(parsed, mask)
		case starType:
			st.Type, parsed = parseString(parsed, mask)
		case starSubType:
			st.SubType, parsed = parseString(parsed, mask)
		case starClass:
			st.Class, parsed = parseString(parsed, mask)
		case starDesignation:
			st.Designation, parsed = parseString(parsed, mask)
		case typeFloat:
			float, rest, err := parseFloat(parsed)
			if err != nil {
				return nil, fmt.Errorf("failed to parse starsystem from '%v': detected %v (%v)", parsed, float, rest)
			}
			parsed = rest
			switch i {
			case 4:
				st.Mass = float
			case 5:
				st.Diameter = float
			case 6:
				st.Luminosity = float
			case 7:
				st.Age = float
			case 9:
				st.OrbitN = float
			case 10:
				st.Eccentricity = float
			}
		}
	}

	// st.Type, parsed = parseString(parsed, starType)
	// st.SubType, parsed = parseString(parsed, starSubType)
	// _, parsed = parseString(parsed, starTSTconnector)
	// st.Class, parsed = parseString(parsed, starClass)
	// for i := 0; i < 7; i++ {
	// 	if parsed == "" {
	// 		continue
	// 	}
	// 	switch i {
	// 	case 4:
	// 		st.Designation, parsed = parseString(parsed, starDesignation)
	// 	default:
	// 		float, rest, err := parseFloat(parsed)
	// 		if err != nil {
	// 			return nil, fmt.Errorf("failed to parse starsystem from '%v': detected %v (%v)", parsed, float, rest)
	// 		}
	// 		parsed = rest
	// 		switch i {
	// 		case 0:
	// 			st.Mass = float
	// 		case 1:
	// 			st.Diameter = float
	// 		case 2:
	// 			st.Luminosity = float
	// 		case 3:
	// 			st.Age = float
	// 		case 5:
	// 			st.OrbitN = float
	// 		case 6:
	// 			st.Eccentricity = float
	// 		}
	// 	}

	// 	//parsed = strings.TrimPrefix(parsed, parsedFloat)
	// }

	return &st, nil
}

func parseFloat(from string) (float64, string, error) {
	if from == "" {
		return 0, "", nil
	}
	reFloat := regexp.MustCompile(typeFloat)
	parsedFloat := reFloat.FindString(from)
	float, err := strconv.ParseFloat(parsedFloat, 64)
	if err != nil {
		return 0, "", fmt.Errorf("failed to parse (%v)==(%v)", from, parsedFloat)
	}
	if float < 0 {
		float = float * -1
	}
	rest := strings.TrimPrefix(from, parsedFloat)
	return float, rest, nil
}

func parseDesignation(from string) (string, string) {
	from = strings.TrimPrefix(from, ":")
	reDesign := regexp.MustCompile(starDesignation)
	des := reDesign.FindString(from)
	rest := strings.TrimPrefix(from, des)
	return des, rest
}

func parseString(from, mask string) (string, string) {
	from = strings.TrimPrefix(from, ":")
	reDesign := regexp.MustCompile(mask)
	des := reDesign.FindString(from)
	rest := strings.TrimPrefix(from, des)
	return des, rest
}

/*
#-T# C-M-D-L[-A]:D-O-E ...
 -T# C-M-D-L[-A]:D-O-E ...


*/
