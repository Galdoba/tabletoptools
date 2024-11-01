package culture

import (
	"fmt"
	"regexp"
	"strings"
)

type Culture struct {
	Diversty        string
	Xenophilia      string
	Uniqueness      string
	Symbology       string
	Cohesion        string
	Progressiveness string
	Expansionism    string
	Militancy       string
}

func FromProfile(s string) (*Culture, error) {
	re := regexp.MustCompile(`([0123456789ABCDEFGHJ])([0123456789ABCDEFGHJ])([0123456789ABCDEFGHJ])([0123456789ABCDEFGHJ])-([0123456789ABCDEFGHJ])([0123456789ABCDEFGHJ])([0123456789ABCDEFGHJ])([0123456789ABCDEFGHJ])`)
	parsed := re.FindString(s)
	if parsed == "" {
		return nil, fmt.Errorf("failed to parse culture from '%v'", s)
	}
	data := strings.Split(parsed, "")
	clt := Culture{}
	clt.Diversty = data[0]
	clt.Xenophilia = data[1]
	clt.Uniqueness = data[2]
	clt.Symbology = data[3]
	clt.Cohesion = data[5]
	clt.Progressiveness = data[6]
	clt.Expansionism = data[7]
	clt.Militancy = data[8]
	return &clt, nil
}

func (clt *Culture) Profile() string {
	return fmt.Sprintf("%v%v%v%v-%v%v%v%v", clt.Diversty, clt.Xenophilia, clt.Uniqueness, clt.Symbology, clt.Cohesion, clt.Progressiveness, clt.Expansionism, clt.Militancy)
}
