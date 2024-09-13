package skill

import (
	"fmt"
	"strconv"
	"strings"
)

type SkillSet struct {
	SklByCode map[string]*skill
}

func NewSet() *SkillSet {
	sks := SkillSet{}
	sks.SklByCode = make(map[string]*skill)
	return &sks
}

func (sks *SkillSet) Gain(value string) error {
	key, minVal := parseGain(value)
	fmt.Println(key, minVal)
	if _, ok := sks.SklByCode[key]; !ok {
		skNew, err := New(key, MaxScore(4), EffectiveScore(minVal))
		if err != nil {
			return err
		}
		sks.SklByCode[key] = skNew
		fmt.Println("New", minVal, skNew)
		return nil
	}
	if sks.SklByCode[key].effectiveScore < minVal {
		sks.SklByCode[key].effectiveScore = minVal
		return nil
	}
	sks.SklByCode[key].effectiveScore++
	return nil
}

func (sks *SkillSet) validate() error {
	for _, sk := range sks.SklByCode {
		if sk.parent == "" && len(sk.specialities) == 0 {
			continue
		}
		for _, spec := range sk.specialities {
			if _, ok := sks.SklByCode[spec]; !ok {
				newSpec, err := New(spec)
				if err != nil {
					return err
				}
				sks.SklByCode[spec] = newSpec
			}
			specSk := sks.SklByCode[spec]
			if sk.effectiveScore > specSk.effectiveScore {
				return fmt.Errorf("speciality score is more than parent score: %v=%v %v=%v", specSk.name, specSk.effectiveScore, sk.name, sk.effectiveScore)
			}

		}
	}
	return nil
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
