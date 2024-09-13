package traveller

import (
	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/characteristic"
)

type Traveller struct {
	CharSet *characteristic.Set
	dice    DiceRoller
}

func New() *Traveller {
	tr := Traveller{}
	tr.CharSet = characteristic.NewSet()
	tr.CharSet.Aslan()
	tr.dice = dice.New()
	return &tr
}

type DiceRoller interface {
	Sroll(string) int
}

func (tr *Traveller) RollCharacteristics() error {
	for _, chr := range tr.CharSet.ByCode {
		if err := chr.Roll(tr.dice); err != nil {
			return err
		}
	}
	return nil
}
