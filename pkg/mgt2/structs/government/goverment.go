package government

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	CENTRALISATION_CONFEDERAL   = "C"
	CENTRALISATION_FEDERAL      = "F"
	CENTRALISATION_UNITARY      = "U"
	AUTHORITY_LEGISLATIVE       = "L"
	AUTHORITY_EXECUTIVE         = "E"
	AUTHORITY_JUDICIAL          = "J"
	AUTHORITY_BALANCED          = "BB"
	STRUCTURE_DEMOS             = "D"
	STRUCTURE_SINGLE_COUNSIL    = "S"
	STRUCTURE_RULER             = "R"
	STRUCTURE_MULTIPLE_COUNSILS = "M"
)

type Government struct {
	Code           string
	Centralisation string
	Authority      string
	//PrimaryStructure     string
	LegislativeStructure string
	ExecutiveStructure   string
	JudicialStructure    string
}

func FromProfile(s string) (*Government, error) {
	re := regexp.MustCompile(`([0123456789ABDCEF]-[CFU]((BB-)?([LEJ])([DSRM])(-[LEJ])([DSRM])?(-[LEJ])([DSRM])?))|([0123456789ABDCEF]-[CFU])([LEJ])([DSRM])`)
	parsed := re.FindString(s)
	if parsed == "" {
		return nil, fmt.Errorf("failed to parse goverment from '%v'", s)
	}
	data := strings.Split(parsed, "-")
	gv := Government{}
	for i, dt := range data {
		switch i {
		case 0:
			gv.Code = dt
		case 1:
			codes := strings.Split(dt, "")
			gv.Centralisation = codes[0]
			gv.Authority = codes[len(codes)-2]
			gv.injectBranch(dt)
		default:
			gv.injectBranch(dt)
		}
	}

	return &gv, nil
}

func (gv *Government) injectBranch(dt string) {
	if strings.Contains(dt, "BB") {
		gv.Authority = AUTHORITY_BALANCED
		return
	}
	codes := strings.Split(dt, "")
	switch codes[len(codes)-2] {
	case AUTHORITY_LEGISLATIVE:
		gv.LegislativeStructure = codes[len(codes)-1]
	case AUTHORITY_EXECUTIVE:
		gv.ExecutiveStructure = codes[len(codes)-1]
	case AUTHORITY_JUDICIAL:
		gv.JudicialStructure = codes[len(codes)-1]
	}
}

func (gv *Government) Profile() string {
	s := fmt.Sprintf("%v-%v", gv.Code, gv.Centralisation)
	switch gv.Authority {
	case AUTHORITY_BALANCED:
		s += AUTHORITY_BALANCED
		s += fmt.Sprintf("-%v%v", AUTHORITY_EXECUTIVE, gv.ExecutiveStructure)
		s += fmt.Sprintf("-%v%v", AUTHORITY_LEGISLATIVE, gv.LegislativeStructure)
		s += fmt.Sprintf("-%v%v", AUTHORITY_JUDICIAL, gv.JudicialStructure)
	case AUTHORITY_EXECUTIVE:
		s += fmt.Sprintf("%v%v", AUTHORITY_EXECUTIVE, gv.ExecutiveStructure)
		s += fmt.Sprintf("-%v%v", AUTHORITY_LEGISLATIVE, gv.LegislativeStructure)
		s += fmt.Sprintf("-%v%v", AUTHORITY_JUDICIAL, gv.JudicialStructure)
	case AUTHORITY_LEGISLATIVE:
		s += fmt.Sprintf("%v%v", AUTHORITY_LEGISLATIVE, gv.LegislativeStructure)
		s += fmt.Sprintf("-%v%v", AUTHORITY_EXECUTIVE, gv.ExecutiveStructure)
		s += fmt.Sprintf("-%v%v", AUTHORITY_JUDICIAL, gv.JudicialStructure)
	case AUTHORITY_JUDICIAL:
		s += fmt.Sprintf("%v%v", AUTHORITY_JUDICIAL, gv.JudicialStructure)
		s += fmt.Sprintf("-%v%v", AUTHORITY_EXECUTIVE, gv.ExecutiveStructure)
		s += fmt.Sprintf("-%v%v", AUTHORITY_LEGISLATIVE, gv.LegislativeStructure)
	}
	return s
}
