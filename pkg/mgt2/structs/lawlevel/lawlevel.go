package lawlevel

import (
	"fmt"
	"regexp"
	"strings"
)

type LawLevel struct {
	Code     string
	Weapons  string
	Economic string
	Criminal string
	Private  string
	Personal string
}

func FromProfile(s string) (*LawLevel, error) {
	re := regexp.MustCompile(`([0123456789ABCDEFGHJ])-([0123456789ABCDEFGHJ])([0123456789ABCDEFGHJ])([0123456789ABCDEFGHJ])([0123456789ABCDEFGHJ])([0123456789ABCDEFGHJ])`)
	parsed := re.FindString(s)
	if parsed == "" {
		return nil, fmt.Errorf("failed to parse law level from '%v'", s)
	}
	data := strings.Split(parsed, "")
	ll := LawLevel{}
	ll.Code = data[0]
	ll.Weapons = data[2]
	ll.Economic = data[3]
	ll.Criminal = data[4]
	ll.Private = data[5]
	ll.Personal = data[6]
	return &ll, nil
}

func (ll *LawLevel) Profile() string {
	return fmt.Sprintf("%v-%v%v%v%v%v", ll.Code, ll.Weapons, ll.Economic, ll.Criminal, ll.Private, ll.Personal)
}
