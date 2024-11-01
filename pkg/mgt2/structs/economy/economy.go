package economy

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Galdoba/tabletoptools/pkg/ehex"
)

type Economy struct {
	TradeCodes           []string
	Importance           int
	ResourceFactor       int
	LaborFactor          int
	InfrastructureFactor int
	EfficiencyFactor     int
	RU                   float64
	PCGWP                float64 //per capita Gross World Product
	WTN                  int     //World Trade Number
	InequalityRating     int
	DevelopmentScore     float64 //
}

//"Ag Ri, +1, F6D-2, -2.340.0, 5940.0, 8, 58, 2.49"

func FromProfile(s string) (*Economy, error) {
	parts := strings.Split(s, ", ")
	if len(parts) != 8 {
		return nil, fmt.Errorf("failed to parse economy from '%v'", s)
	}
	eco := Economy{}
	tc, err := parseTradeCodes(parts[0])
	if err != nil {
		return nil, err
	}
	eco.TradeCodes = tc
	imp, err := parseInt(parts[1])
	if err != nil {
		return nil, err
	}
	eco.Importance = imp
	ecoStats, err := parseEcoStats(parts[2])
	if err != nil {
		return nil, err
	}
	eco.ResourceFactor = ecoStats[0]
	eco.LaborFactor = ecoStats[1]
	eco.InfrastructureFactor = ecoStats[2]
	eco.EfficiencyFactor = ecoStats[3]

	ru, err := parseFloat(parts[3])
	if err != nil {
		return nil, err
	}
	eco.RU = ru
	pcGWP, err := parseFloat(parts[4])
	if err != nil {
		return nil, err
	}
	eco.PCGWP = pcGWP

	wtn, err := parseInt(parts[5])
	if err != nil {
		return nil, err
	}
	eco.WTN = wtn

	ir, err := parseInt(parts[6])
	if err != nil {
		return nil, err
	}
	eco.InequalityRating = ir

	ds, err := parseFloat(parts[7])
	if err != nil {
		return nil, err
	}
	eco.DevelopmentScore = ds

	return &eco, nil
}

func (eco *Economy) Profile() string {
	tcStr := ""
	for _, tc := range eco.TradeCodes {
		tcStr += tc + " "
	}
	tcStr = strings.TrimSuffix(tcStr, " ")

	impStr := fmt.Sprintf("%v", eco.Importance)
	if eco.Importance > -1 {
		impStr = "+" + impStr
	}

	effStr := fmt.Sprintf("%v", eco.EfficiencyFactor)
	if eco.EfficiencyFactor > -1 {
		effStr = "+" + effStr
	}

	return fmt.Sprintf("%v, %v, %v%v%v%v, %v, %v, %v, %v, %v",
		tcStr,
		impStr,
		fmt.Sprintf("%v", ehex.ToCode(eco.ResourceFactor)),
		fmt.Sprintf("%v", ehex.ToCode(eco.LaborFactor)),
		fmt.Sprintf("%v", ehex.ToCode(eco.InfrastructureFactor)),
		effStr,
		fmt.Sprintf("%.1f", eco.RU),
		fmt.Sprintf("%.1f", eco.PCGWP),
		eco.WTN,
		eco.InequalityRating,
		eco.DevelopmentScore,
	)
}

func parseTradeCodes(str string) ([]string, error) {
	if str == "" {
		return []string{}, nil
	}
	re := regexp.MustCompile(`([A-Z][a-z] )+([A-Z][a-z])`)
	parsed := re.FindString(str)
	if parsed == "" {
		return nil, fmt.Errorf("failed to parse trade codes from '%v'", str)
	}
	tc := []string{}
	for _, field := range strings.Fields(parsed) {
		tc = append(tc, field)
	}
	return tc, nil
}

func parseInt(str string) (int, error) {
	imp, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("failed to parse int field '%v'", str)
	}

	return imp, nil
}

func parseFloat(str string) (float64, error) {
	fl, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse float field '%v'", str)
	}

	return fl, nil
}

func parseEcoStats(str string) ([]int, error) {
	re := regexp.MustCompile(`([0123456789ABCDEFGHJ][0123456789ABCDEFGHJ][0123456789ABCDEFGHJ])([+-])([012345])`)
	parsed := re.FindString(str)
	if parsed == "" {
		return nil, fmt.Errorf("failed to parse economy block from '%v'", str)
	}
	econ := []int{}
	data := strings.Split(parsed, "")
	if len(data) != 5 {
		return nil, fmt.Errorf("failed to parse economy block from '%v'", str)
	}
	econ = append(econ, ehex.ValueOf(data[0]))
	econ = append(econ, ehex.ValueOf(data[1]))
	econ = append(econ, ehex.ValueOf(data[2]))
	val, err := strconv.Atoi(data[3] + data[4])
	if err != nil {
		return nil, fmt.Errorf("failed to parse economy block from '%v'", str)
	}
	econ = append(econ, val)
	return econ, nil
}
