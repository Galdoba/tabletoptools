package characteristic

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (ch *characteristic) Encode() string {
	return fmt.Sprintf("%v : %v%v (%vd6%v)", ch.name, ch.effectiveScore, maxScoreToStr(ch.effectiveScore, ch.maxScore), ch.creationDice, modToStr(ch.creationMod))
}

func Decode(str string) (*characteristic, error) {
	//C1 : "Strenght : 12/15 (2d6+4)"
	err := errors.New("no decoding started")
	sp1 := strings.Split(str, " : ")
	if len(sp1) != 2 {
		return nil, fmt.Errorf("decoding failed: first split '%v'", str)
	}
	name := sp1[0]
	sp2 := strings.Split(sp1[1], " ")
	if len(sp2) != 2 {
		return nil, fmt.Errorf("decoding failed: second split '%v'", sp1[1])
	}
	sp3 := strings.Split(sp2[0], "/")
	if len(sp3) > 2 {
		return nil, fmt.Errorf("decoding failed: third split '%v'", sp2[0])
	}
	score := 0
	maxScore := 0
	score, err = strconv.Atoi(sp3[0])
	if err != nil {
		return nil, fmt.Errorf("decoding failed: score '%v': %v", sp3[0], err)
	}
	if len(sp3) == 2 {
		maxScore, err = strconv.Atoi(sp3[1])
		if err != nil {
			return nil, fmt.Errorf("decoding failed: maxscore '%v': %v", sp3[1], err)
		}
	}

	dice := strings.TrimPrefix(sp2[1], "(")
	dice = strings.TrimSuffix(dice, ")")
	crDice := ""
	crMod := 0
	diceNum := 0
	switch {
	case strings.Contains(dice, "+"):
		sp4 := strings.Split(dice, "+")
		if len(sp4) != 2 {
			return nil, fmt.Errorf("decoding failed: fourth split '%v'", dice)
		}
		crDice = strings.TrimSuffix(sp4[0], "d6")
		crMod, err = strconv.Atoi(sp4[1])
		if err != nil {
			return nil, fmt.Errorf("decoding failed: creationMod '%v': %v", sp4[1], err)
		}
	case strings.Contains(dice, "-"):
		sp4 := strings.Split(dice, "-")
		if len(sp4) != 2 {
			return nil, fmt.Errorf("decoding failed: fourth split '%v'", dice)
		}
		crDice = strings.TrimSuffix(sp4[0], "d6")
		crMod, err = strconv.Atoi(sp4[1])
		crMod = crMod * -1
		if err != nil {
			return nil, fmt.Errorf("decoding failed: creationMod '%v': %v", sp4[1], err)
		}
	default:
		crDice = strings.TrimSuffix(dice, "d6")
	}
	diceNum, err = strconv.Atoi(crDice)
	if err != nil {
		return nil, fmt.Errorf("decoding failed: diceNum '%v': %v", diceNum, err)
	}
	return New(name, EffectiveScore(score), MaxScore(maxScore), CreationDice(diceNum), CreationMod(crMod))
}

func modToStr(m int) string {
	if m == 0 {
		return ""
	}
	if m > 0 {
		return fmt.Sprintf("+%v", m)
	}
	return fmt.Sprintf("%v", m)
}

func maxScoreToStr(cs, ms int) string {
	if cs == ms {
		return ""
	}
	return fmt.Sprintf("/%v", ms)
}
