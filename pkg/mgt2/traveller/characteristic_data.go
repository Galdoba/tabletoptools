package traveller

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/characteristic"
	"github.com/Galdoba/tabletoptools/pkg/definition/ruleset"
)

const (
	STR = "STREIGHT"
	DEX = "DEXTERITY"
	END = "ENDURANCE"
	INT = "INTELIGENCE"
	EDU = "EDUCATION"
	SOC = "SOCIAL"
	PSI = "PSIONICS"
	SAN = "SANITY"
)

type CharacteristicsData struct {
	Header string `json:"Characteristics,omitempty"`
	Vals   string `json:"Values,omitempty"`
	UPP    string `json:"Universal Profile,omitempty"`
}

/*
0123456789
|   STR   |  DEX B  |  END C  |   INT   |  STR A  |  STR A  |  STR A  |  STR A  |  STR A  |  STR A  |  STR A  |  STR A  |
| 10 (+1) | 11 (+1) | 12 (+2) |  8 (+0) | 10 (+1) | 10 (+1) | 10 (+1) | 10 (+1) | 10 (+1) | 10 (+1) | 10 (+1) | 10 (+1) |
*/

func newCharacteristicsData(cs *characteristic.Set) CharacteristicsData {
	cd := CharacteristicsData{}
	for _, ch := range listOrder() {
		for chr, val := range cs.CHR {
			if chr.Abbreviation == ch {
				cd.Header += "   " + ch + "   "
				block := fmt.Sprintf("%v", val.Max)
				for len(block) < 2 {
					block = " " + block
				}
				block = " " + block + " ("
				mod := val.DM(ruleset.MGT2)
				mds := fmt.Sprintf("%v", mod)
				if mod >= 0 {
					mds = "+" + mds
				}
				block += mds + ") "
				cd.Vals += block
			}
		}
	}
	return cd
}

func listOrder() []string {
	return []string{
		characteristic.STR,
		characteristic.DEX,
		characteristic.AGI,
		characteristic.GRA,
		characteristic.END,
		characteristic.STA,
		characteristic.VIG,
		characteristic.INT,
		characteristic.EDU,
		characteristic.TRA,
		characteristic.INS,
		characteristic.SOC,
		characteristic.CHA,
		characteristic.CAS,
		characteristic.TER,
		characteristic.PSI,
		characteristic.SAN,
		characteristic.WLT,
		characteristic.LCK,
		characteristic.MOR,
	}
}
