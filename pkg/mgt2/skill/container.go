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

func (cs *Set) Train(name string) error {
	if _, ok := cs.SklByCode[name]; !ok {
		sk, err := New(name, MaxScore(4), EffectiveScore(1))
		if err != nil {
			return err
		}
		cs.SklByCode[name] = sk
		return nil
	}
	sk := cs.SklByCode[name]
	switch len(sk.specialities) {
	case 0:
		if (sk.effectiveScore + 1) > sk.maxScore {
			return fmt.Errorf("can't train %v: maximum level (%v) reached", sk.name, sk.maxScore)
		}
		sk.effectiveScore++
		return nil
	default:
		specVals := []int{}
		for _, specKey := range sk.specialities {
			if specSkill, ok := cs.SklByCode[specKey]; ok {
				specVals = append(specVals, specSkill.effectiveScore)
			}
		}

	}

}

func parseGain(value string) (string, int) {
	sp1 := strings.Split(value, " ")
	key := parseBrakets(value)
	min, err := strconv.Atoi(sp1[len(sp1)-1])
	if err != nil {
		min = 1
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
