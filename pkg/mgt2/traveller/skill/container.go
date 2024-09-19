package skill

import (
	"fmt"
	"strconv"
	"strings"
)

type Set struct {
	SklByCode map[string]*skill
}

func NewSet() *Set {
	sks := Set{}
	sks.SklByCode = make(map[string]*skill)
	return &sks
}

func (ss *Set) Train(name string) error {
	sk, ok := ss.SklByCode[name]
	switch ok {
	case false:
		skl, err := New(name, MaxScore(4))
		if err != nil {
			return fmt.Errorf("skill creation failed: %v")
		}
		if len(skl.specialities) == 0 {
			skl.effectiveScore++
		}
		ss.SklByCode[name] = skl
	case true:
		if len(sk.specialities) != 0 {
			return fmt.Errorf("speciality must be chosen")
		}
		if sk.effectiveScore+1 > sk.maxScore {
			return fmt.Errorf("maximum score reached")
		}
		ss.SklByCode[name].effectiveScore++
	}
	return nil
}

func (ss *Set) Ensure(name string, value int) error {
	sk, ok := ss.SklByCode[name]
	switch ok {
	case false:
		skl, err := New(name, EffectiveScore(4))
		if err != nil {
			return fmt.Errorf("skill creation failed: %v")
		}
		switch len(skl.specialities) {
		default:
		case 0:
			skl.effectiveScore = value
		}
		ss.SklByCode[name] = skl
	case true:
		if len(sk.specialities) != 0 && value != 0 {
			return fmt.Errorf("speciality must be chosen")
		}
		if value > sk.maxScore {
			return fmt.Errorf("new value is higher than maximum score")
		}
		if sk.effectiveScore < value {
			sk.effectiveScore = value
		}
		ss.SklByCode[name] = sk
	}
	return nil
}

func parseGain(value string) (string, int) {
	sp1 := strings.Split(value, " ")
	key := parseBrakets(value)
	min, err := strconv.Atoi(sp1[len(sp1)-1])
	if err != nil {
		min = -1
	}
	minStr := fmt.Sprintf(" %v", min)
	switch key {
	case "no brackets":
		key = strings.TrimSuffix(value, minStr)
	}
	return key, min
}

func parseBrakets(text string) string {
	filtered := ""
	read := false
	for _, letter := range strings.Split(text, "") {
		switch read {
		case false:
			if letter == "(" {
				read = true
			}
		case true:
			if letter == ")" {
				return filtered
			}
			filtered += letter
		}
	}
	if read {
		return "error"
	}
	return "no brackets"
}

func (ss *Set) Map() map[string]int {
	ssmap := make(map[string]int)
	for k, s := range ss.SklByCode {
		ssmap[k] = s.effectiveScore
	}
	return ssmap
}

func UnMap(ssmap map[string]int) (*Set, error) {
	ss := NewSet()
	for k, score := range ssmap {
		skl, err := New(k, EffectiveScore(score))
		if err != nil {
			return nil, err
		}
		ss.SklByCode[k] = skl
	}
	return ss, nil
}

/*
type StatChanger interface {
Max(key string) (int, error)
NewMax(key string, max int) error

}
*/
