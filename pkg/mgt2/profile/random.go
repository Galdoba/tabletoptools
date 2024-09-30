package profile

import (
	"fmt"
	"strings"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/ehex"
)

type Roller interface {
	Sroll(string) int
}

func newRoller() Roller {
	return dice.New()
}

func (pr *universalProfile) GenerateMissingData(dices ...Roller) error {
	roller := *new(Roller)
	switch len(dices) {
	case 0:
		roller = newRoller()
	}
	generateted := 0
	for _, rollFunc := range []func(*universalProfile, Roller) error{
		//rollTemperature,
		rollSize,
		rollAtmo,
		rollTemperature,
		rollHydr,
		rollPops,
		rollGovr,
		rollLaws,
		rollPort,
		rollTL,
	} {
		err := rollFunc(pr, roller)
		if err != nil {
			continue
		}
		generateted++
	}
	if generateted == 0 {
		return fmt.Errorf("no missing data to generate")
	}
	return nil
}

func assertKeyAbsence(key string, pr *universalProfile) error {
	if val, ok := pr.profilePoints[key]; ok {
		switch val {
		case "?":
		default:
			return fmt.Errorf("%v was generated", key)
		}
	}
	return nil
}

func rollSize(pr *universalProfile, roller Roller) error {
	if err := assertKeyAbsence(KEY_Size, pr); err != nil {
		return err
	}
	size := roller.Sroll("2d6-2")
	if size == 10 {
		for roller.Sroll("1d6") == 6 {
			size++
		}
	}
	size = bound(size, 0, 15)
	pr.SetValue(KEY_Size, ehex.ToCode(size))
	return nil
}

func rollAtmo(pr *universalProfile, roller Roller) error {
	if err := assertKeyAbsence(KEY_Atmo, pr); err != nil {
		return err
	}
	size := ehex.ValueOf(pr.GetValue(KEY_Size))
	atmo := roller.Sroll("2d6-7") + size
	atmo = bound(atmo, 0, 15)
	pr.SetValue(KEY_Atmo, ehex.ToCode(atmo))
	return nil
}

func rollTemperature(pr *universalProfile, roller Roller) error {
	if err := assertKeyAbsence(KEY_Temp, pr); err != nil {
		return err
	}
	switch roller.Sroll("2d6") {
	case 2:
		pr.SetValue(KEY_Temp, "Fr")
	case 3, 4:
		pr.SetValue(KEY_Temp, "Co")
	case 5, 6, 7, 8, 9:
		pr.SetValue(KEY_Temp, "Te")
	case 10, 11:
		pr.SetValue(KEY_Temp, "Ho")
	case 12:
		pr.SetValue(KEY_Temp, "Bo")
	default:
		return fmt.Errorf("roll out of bounds")
	}
	return nil
}

func rollHydr(pr *universalProfile, roller Roller) error {
	if err := assertKeyAbsence(KEY_Hydr, pr); err != nil {
		return err
	}
	size := ehex.ValueOf(pr.GetValue(KEY_Size))
	atmo := ehex.ValueOf(pr.GetValue(KEY_Atmo))
	temp := pr.GetValue(KEY_Temp)
	dm := 0
	switch size {
	case 0, 1:
		dm += -1000
	}
	switch atmo {
	case 0, 1, 10, 11, 12, 13, 14, 15:
		dm += -4
	}
	switch atmo {
	case 13, 15:
	default:
		switch temp {
		case "Ho":
			dm += -2
		case "Bo":
			dm += -6
		}
	}

	hydr := roller.Sroll("2d6-7") + atmo + dm
	hydr = bound(hydr, 0, 10)
	pr.SetValue(KEY_Hydr, ehex.ToCode(hydr))
	return nil
}

func rollPops(pr *universalProfile, roller Roller) error {
	if err := assertKeyAbsence(KEY_Pops, pr); err != nil {
		return err
	}
	pops := roller.Sroll("2d6-2")
	if pops == 10 {
		for roller.Sroll("2d6") == 12 {
			pops++
		}
	}
	pops = bound(pops, 0, 12)
	pr.SetValue(KEY_Pops, ehex.ToCode(pops))
	return nil
}

func rollGovr(pr *universalProfile, roller Roller) error {
	if err := assertKeyAbsence(KEY_Govr, pr); err != nil {
		return err
	}
	pops := ehex.ValueOf(pr.GetValue(KEY_Pops))
	govr := roller.Sroll("2d6-7") + pops
	govr = bound(govr, 0, 15)
	pr.SetValue(KEY_Govr, ehex.ToCode(govr))
	return nil
}

func rollPort(pr *universalProfile, roller Roller) error {
	if err := assertKeyAbsence(KEY_Starport, pr); err != nil {
		return err
	}
	pops := ehex.ValueOf(pr.GetValue(KEY_Pops))
	dm := 0
	switch pops {
	case 0, 1, 2:
		dm += -2
	case 3, 4:
		dm += -1
	case 8, 9:
		dm += 1
	case 10, 11, 12:
		dm += 2
	}
	port := roller.Sroll("2d6") + dm
	portCode := "X"
	switch port {
	case 3, 4:
		portCode = "E"
	case 5, 6:
		portCode = "D"
	case 7, 8:
		portCode = "C"
	case 9, 10:
		portCode = "B"
	case 11, 12, 13, 14:
		portCode = "A"
	}
	if pops == 0 && strings.Contains("ABCD", portCode) {
		portCode = "E"
	}
	pr.SetValue(KEY_Starport, portCode)
	return nil
}

func rollLaws(pr *universalProfile, roller Roller) error {
	if err := assertKeyAbsence(KEY_Laws, pr); err != nil {
		return err
	}
	govr := ehex.ValueOf(pr.GetValue(KEY_Govr))
	laws := roller.Sroll("2d6-7") + govr
	laws = bound(laws, 0, 30)
	pr.SetValue(KEY_Laws, ehex.ToCode(laws))
	return nil
}

func rollTL(pr *universalProfile, roller Roller) error {
	if err := assertKeyAbsence(KEY_TL, pr); err != nil {
		return err
	}
	dm := 0
	switch pr.GetValue(KEY_Starport) {
	case "A":
		dm += 6
	case "B":
		dm += 4
	case "C":
		dm += 2
	case "X":
		dm += -4
	}
	switch pr.GetValue(KEY_Size) {
	case "0", "1":
		dm += 2
	case "2", "3", "4":
		dm += 1
	}
	switch pr.GetValue(KEY_Atmo) {
	case "0", "1", "2", "3", "A", "B", "C", "D", "E", "F":
		dm += 1
	}
	switch pr.GetValue(KEY_Hydr) {
	case "0", "9":
		dm += 1
	case "A":
		dm += 2
	}
	switch pr.GetValue(KEY_Pops) {
	case "1", "2", "3", "4", "5", "8":
		dm += 1
	case "9":
		dm += 2
	case "A", "B", "C":
		dm += 4
	}
	switch pr.GetValue(KEY_Govr) {
	case "0", "5":
		dm += 1
	case "7":
		dm += 2
	case "D", "E", "C":
		dm += -2
	}
	tl := roller.Sroll("1d6") + dm
	tl = bound(tl, 0, 30)
	pr.SetValue(KEY_TL, ehex.ToCode(tl))
	return nil
}

func bound(val, min, max int) int {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
