package population

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Population struct {
	Code             string
	Value            float64
	PCR              string
	Urbanization     int
	MajorSettlements int
}

func (p *Population) Profile() string {
	s := fmt.Sprintf("%v", p.Code)
	f := strconv.FormatFloat(p.Value, 'f', 3, 64)
	for f != strings.TrimSuffix(f, "0") {
		f = strings.TrimSuffix(f, "0")
	}
	s += fmt.Sprintf("-%v", f)
	s += fmt.Sprintf("-%v", p.PCR)
	urb := fmt.Sprintf("%v", p.Urbanization)
	if p.Urbanization > 99 {
		urb = ">99"
	}
	s += fmt.Sprintf("-%v", urb)
	s += fmt.Sprintf("-%d", p.MajorSettlements)
	return s
}

func FromProfile(s string) (*Population, error) {
	re := regexp.MustCompile(`[0123456789ABDCEF]-\d(\.)?(\d)?(\d)?(\d)?-[0123456789A]-(\>)?\d(\d)?-\d(\d)?`)
	parsed := re.FindString(s)
	if parsed == "" {
		return nil, fmt.Errorf("failed to parse '%v'", s)
	}
	p := Population{}
	data := strings.Split(parsed, "-")
	p.Code = data[0]
	f, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Population Value '%v'", data[1])
	}
	p.Value = f
	p.PCR = data[2]
	switch data[3] {
	case ">99":
		p.Urbanization = 100
	default:
		i1, err := strconv.Atoi(data[3])
		if err != nil {
			return nil, fmt.Errorf("failed to parse Population Urbanisation '%v'", data[3])
		}
		p.Urbanization = i1
	}

	i2, err := strconv.Atoi(data[4])
	if err != nil {
		return nil, fmt.Errorf("failed to parse Population Major Settlements '%v'", data[4])
	}
	p.MajorSettlements = i2
	return &p, nil
}
