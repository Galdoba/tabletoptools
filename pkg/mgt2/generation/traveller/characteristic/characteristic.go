package characteristic

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/traveller/key"
)

const (
	STR    = key.C1_STR
	DEX    = key.C2_DEX
	END    = key.C3_END
	INT    = key.C4_INT
	EDU    = key.C5_EDU
	SOC    = key.C6_SOC
	STR_DM = "STR DM"
	DEX_DM = "DEX DM"
	END_DM = "END DM"
	INT_DM = "INT DM"
	EDU_DM = "EDU DM"
	SOC_DM = "SOC DM"
)

type CharSet struct {
	Chars map[string]int
}

func (cs *CharSet) Value(key string) int {
	switch key {
	default:
		panic(fmt.Sprintf("invalid key '%v'", key))
	case STR, DEX, END, INT, EDU, SOC:
		return cs.Chars[key]
	case STR_DM:
		return characteristicDM(cs.Chars[STR])
	case DEX_DM:
		return characteristicDM(cs.Chars[DEX])
	case END_DM:
		return characteristicDM(cs.Chars[END])
	case INT_DM:
		return characteristicDM(cs.Chars[INT])
	case EDU_DM:
		return characteristicDM(cs.Chars[EDU])
	case SOC_DM:
		return characteristicDM(cs.Chars[SOC])
	}
}

func (cs *CharSet) Modify(key string, byValue int) error {
	switch key {
	default:
		log.Fatalf("invalid characteristic key '%v'", key)
	case STR, DEX, END, INT, EDU, SOC:
		cs.Chars[key] = boundChar(cs.Chars[key] + byValue)
	}
	return nil
}

func boundChar(i int) int {
	if i < 0 {
		return 0
	}
	if i > 15 {
		return 15
	}
	return i
}

func NewSet() *CharSet {
	cs := CharSet{}
	cs.Chars = make(map[string]int)
	cs.Chars[STR] = 0
	cs.Chars[DEX] = 0
	cs.Chars[END] = 0
	cs.Chars[INT] = 0
	cs.Chars[EDU] = 0
	cs.Chars[SOC] = 0
	return &cs
}

type Characteristic struct {
	Name  string
	Value int
}

func (c *Characteristic) String() string {
	return fmt.Sprintf("%v : %v", c.Name, c.Value)
}

func UnString(s string) *Characteristic {
	c := &Characteristic{}
	parts := strings.Split(s, " : ")
	if len(parts) != 2 {
		return nil
	}
	c.Name = parts[0]
	c.Value, _ = strconv.Atoi(parts[1])
	return c
}

func characteristicDM(i int) int {
	if i <= 0 {
		return -3
	}
	switch i {
	case 1, 2:
		return -2
	case 3, 4, 5:
		return -1
	case 6, 7, 8:
		return 0
	case 9, 10, 11:
		return 1
	case 12, 13, 14:
		return 2
	default:
		return 3
	}
}
