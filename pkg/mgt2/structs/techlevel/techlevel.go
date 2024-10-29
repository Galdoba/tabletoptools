package techlevel

import (
	"fmt"
	"regexp"
	"strings"
)

type TechLevel struct {
	HighCommon       string
	LowCommon        string
	Energy           string
	Electronics      string
	Manufactoring    string
	Medical          string
	Enviromental     string
	LandTransport    string
	WaterTransport   string
	AirTransport     string
	SpaceTransport   string
	PersonalMilitary string
	HeavyMilitary    string
	Novelty          string
}

var anyEhex = "([0123456789ABCDEFGHJKLMNPQRSTUVWXYZ])"

func FromProfile(s string) (*TechLevel, error) {
	reString := fmt.Sprintf(`%v-%v-%v%v%v%v%v-%v%v%v%v-%v%v-%v`, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex, anyEhex)
	re := regexp.MustCompile(reString)
	parsed := re.FindString(s)
	if parsed == "" {
		return nil, fmt.Errorf("failed to parse tech level from '%v'", s)
	}
	tl := TechLevel{}
	data := strings.Split(parsed, "")
	tl.HighCommon = data[0]
	tl.LowCommon = data[2]
	tl.Energy = data[4]
	tl.Electronics = data[5]
	tl.Manufactoring = data[6]
	tl.Medical = data[7]
	tl.Enviromental = data[8]
	tl.LandTransport = data[10]
	tl.WaterTransport = data[11]
	tl.AirTransport = data[12]
	tl.SpaceTransport = data[13]
	tl.PersonalMilitary = data[15]
	tl.HeavyMilitary = data[16]
	tl.Novelty = data[18]
	return &tl, nil
}

func (tl *TechLevel) Profile() string {
	return fmt.Sprintf("%v-%v-%v%v%v%v%v-%v%v%v%v-%v%v-%v",
		tl.HighCommon,
		tl.LowCommon,
		tl.Energy,
		tl.Electronics,
		tl.Manufactoring,
		tl.Medical,
		tl.Enviromental,
		tl.LandTransport,
		tl.WaterTransport,
		tl.AirTransport,
		tl.SpaceTransport,
		tl.PersonalMilitary,
		tl.HeavyMilitary,
		tl.Novelty,
	)
}
