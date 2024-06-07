package check

import "fmt"

const (
	UNDEFINED           = "Undefined"
	TYPE_CHARACTERISTIC = "Characteristic"
	TYPE_SKILL          = "Skill"

	DIF_Easy           = 4
	DIF_Routine        = 6
	DIF_Average        = 8
	DIF_Difficult      = 10
	DIF_Very_Difficult = 12
	DIF_Formidable     = 14
	DIF_Impossible     = 16
)

type Check struct {
	Type       string //   skill/char
	DiceCode   string //   2D/3D/...
	Ruleset    int
	Difficulty int
	Mods       map[string]Modifier
	rolled     bool
	result     int
}

type Modifier interface {
	DM(int) int
}

type Roller interface {
	Sroll(string) int
}

func newCheck(ruleset int) *Check {
	c := Check{}
	c.Ruleset = ruleset
	c.Difficulty = DIF_Average
	c.DiceCode = UNDEFINED
	c.Type = UNDEFINED
	c.Mods = map[string]Modifier{}
	return &c
}

func (c *Check) WithMod(key string, mod Modifier) *Check {
	c.Mods[key] = mod
	return c
}

func (c *Check) WithDiceCode(code string) *Check {
	c.DiceCode = code
	return c
}

type customDM struct {
	key string
	val int
}

func CustomDM(key string, val int) *customDM {
	return &customDM{key: key, val: val}
}

func (cdm *customDM) Mod(ruleset int) int {
	return cdm.val
}

func (c *Check) Roll(roller Roller) error {
	if c.Type == UNDEFINED {
		return fmt.Errorf("can't roll undefined type check")
	}
	if c.DiceCode == UNDEFINED {
		return fmt.Errorf("can't roll with undefined dice code")
	}
	r := roller.Sroll(c.DiceCode)
	for _, dm := range c.Mods {
		r += dm.DM(c.Ruleset)
	}
	c.result = r
	c.rolled = true
	return nil
}

func (c *Check) Passed(roller Roller) bool {
	if !c.rolled {
		c.Roll(roller)
	}
	return c.result >= c.Difficulty
}

func (c *Check) Effect(roller Roller) int {
	if !c.rolled {
		c.Roll(roller)
	}
	return c.result - c.Difficulty
}
