package starport

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Starport struct {
	Class              string
	HighPortPresense   bool
	DownPortPresense   bool
	AdjustedImportance int
}

var anyEhex = "([0123456789ABCDEFGHJKLMNPQRSTUVWXYZ])"

var portClass = "([ABCDEFGHYX])"
var portPresent = "([YN])"
var impModified = "([+-][0123456])"

func FromProfile(s string) (*Starport, error) {
	reString := fmt.Sprintf(`%v-H%v:D%v:%v`, portClass, portPresent, portPresent, impModified)
	re := regexp.MustCompile(reString)
	parsed := re.FindString(s)
	if parsed == "" {
		return nil, fmt.Errorf("failed to parse starport from '%v'", s)
	}
	sp := Starport{}
	data := strings.Split(parsed, "")
	sp.Class = data[0]
	if data[3] == "Y" {
		sp.HighPortPresense = true
	}
	if data[6] == "Y" {
		sp.DownPortPresense = true
	}
	imp, err := strconv.Atoi(data[8] + data[9])
	if err != nil {
		return nil, fmt.Errorf("failed to parse starport from '%v'", s)
	}
	sp.AdjustedImportance = imp

	return &sp, nil
}

func (sp *Starport) Profile() string {
	return fmt.Sprintf("%v-H%v:D%v:%v",
		sp.Class,
		presense(sp.HighPortPresense),
		presense(sp.DownPortPresense),
		importance(sp.AdjustedImportance),
	)
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
