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

var starClass = "(Ia|Ib|II|III|IV|V|VI|BD|D)?"
var starType = "(O|B|A|F|G|K|M|L|T|Y|)?"
var starTSTconnector = "( )?"
var starSubType = "([0123456789])?"
var typeFloat = `(-[0123456789]([0123456789])?([0123456789])?(\.)?([0123456789])?([0123456789])?([0123456789])?([0123456789])?([0123456789])?([0123456789])?)?`
var starDesignation = `((:)?(Aab|Aa|Ab|A|Bab|Ba|Bb|B|Cab|Ca|Cb|C|Dab|Da|Db|D|))?`

var starProfile = fmt.Sprintf(`%v%v%v%v%v%v%v%v%v%v%v`, starType, starSubType, starTSTconnector, starClass, typeFloat, typeFloat, typeFloat, typeFloat, starDesignation, typeFloat, typeFloat)

var starNumber = `([2345678]-)?` + starProfile

var starSystem = fmt.Sprintf(`%v%v`, starNumber, starProfile)

var anyEhex = "([0123456789ABCDEFGHJKLMNPQRSTUVWXYZ])"

func (st *Star) Profile() string {
	s := ""
	s += fmt.Sprintf("%v", st.Type)
	s += fmt.Sprintf("%v", st.SubType)
	if st.SubType != "" {
		s += " "
	}
	s += fmt.Sprintf("%v", st.Class)
	if st.Mass > 0.0 {
		s += fmt.Sprintf("-%v", strconv.FormatFloat(st.Mass, 'f', -1, 64))
	}
	if st.Diameter > 0.0 {
		s += fmt.Sprintf("-%v", strconv.FormatFloat(st.Diameter, 'f', -1, 64))
	}
	if st.Luminosity > 0.0 {
		s += fmt.Sprintf("-%v", strconv.FormatFloat(st.Luminosity, 'f', -1, 64))
	}
	if st.Age > 0.0 {
		s += fmt.Sprintf("-%v", strconv.FormatFloat(st.Age, 'f', -1, 64))
	}
	if st.Designation != "" {
		s += fmt.Sprintf(":%v", st.Designation)
	}
	if st.OrbitN > 0.0 {
		s += fmt.Sprintf("-%v", strconv.FormatFloat(st.OrbitN, 'f', -1, 64))
	}
	if st.Eccentricity > 0.0 {
		s += fmt.Sprintf("-%v", strconv.FormatFloat(st.Eccentricity, 'f', -1, 64))
	}

	return s
}

func (ss *StarSystem) Profile() string {
	s := ""
	switch len(ss.Stars) {
	case 1:
		s = fmt.Sprintf("%v", ss.Stars[0].Profile())
	default:
		age := ss.Age
		for i := range ss.Stars {
			ss.Stars[i].Age = 0
		}
		ss.Stars[0].Age = age
		s += fmt.Sprintf("%v", len(ss.Stars))
		for _, star := range ss.Stars {
			s += fmt.Sprintf("-%v", star.Profile())
		}
	}
	return s
}

func FromProfile(from string) (*StarSystem, error) {
	ss := StarSystem{}
	// mask := starSystem
	// re := regexp.MustCompile(mask)
	// parsed := re.FindString(from)
	num, next := grepStarNumber(from)
	switch num {
	case 1:
		st, err := StarFromProfile(next)
		if err != nil {
			return nil, fmt.Errorf("failed to parse star: %v", err)
		}
		ss.Stars = append(ss.Stars, st)
	default:
		rest := strings.TrimPrefix(from, fmt.Sprintf("%v-", num))
		for i := 0; i < num; i++ {
			reString := starProfile
			re := regexp.MustCompile(reString)
			next := re.FindString(rest)
			st, err := StarFromProfile(next)
			if err != nil {
				return nil, fmt.Errorf("failed to parse star %v: %v\n%v", num, err, rest)
			}
			ss.Stars = append(ss.Stars, st)
			rest = strings.TrimPrefix(rest, next)
			rest = strings.TrimPrefix(rest, "-")

		}
	}
	ss.Age = ss.Stars[0].Age
	for i := range ss.Stars {
		if i == 0 {
			continue
		}
		ss.Stars[i].Age = 0
	}
	return &ss, nil
}

func grepStarNumber(s string) (int, string) {
	for i := 0; i < 9; i++ {
		pref := fmt.Sprintf("%v-", i)
		//fmt.Println(pref)
		if strings.HasPrefix(s, pref) {
			return i, strings.TrimPrefix(s, pref)
		}
	}
	return 1, s
}

func StarFromProfile(s string) (*Star, error) {
	reString := starProfile
	re := regexp.MustCompile(reString)
	parsed := re.FindString(s)
	if parsed == "" {
		return nil, fmt.Errorf("00failed to parse stars from '%v'", s)
	}
	st := Star{}
	for i, mask := range []string{starType, starSubType, starTSTconnector, starClass, typeFloat, typeFloat, typeFloat, typeFloat, starDesignation, typeFloat, typeFloat} {
		switch mask {
		case starTSTconnector:
			_, parsed = parseString(parsed, mask)
			//			fmt.Println("case starTSTconnector:", parsed)
		case starType:
			st.Type, parsed = parseString(parsed, mask)
			//			fmt.Println("case starType:", st.Type)
		case starSubType:
			st.SubType, parsed = parseString(parsed, mask)
			//			fmt.Println("case starSubType:", st.SubType)
		case starClass:
			st.Class, parsed = parseString(parsed, mask)
			//			fmt.Println("case Class:", st.Class)
		case starDesignation:
			st.Designation, parsed = parseString(parsed, mask)
			//			fmt.Println("case starDesignation:", st.Designation)
		case typeFloat:
			float, rest, err := parseFloat(parsed)
			if err != nil {
				continue
				//return nil, fmt.Errorf("failed to parse star from '%v': detected %v (%v)", parsed, float, rest)
			}
			parsed = rest
			//			fmt.Println("case float:", float, err)
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
