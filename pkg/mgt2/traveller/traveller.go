package traveller

import (
	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/characteristic"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/skill"
)

type Traveller struct {
	Species      string
	CharSet      *characteristic.Set
	SkillSet     *skill.Set
	dice         DiceRoller
	creationMode bool
}

func New(options ...Options) *Traveller {
	tr := Traveller{}
	tr.CharSet = characteristic.NewSet()
	tr.SkillSet = skill.NewSet()
	settings := defaultCreationOptions()
	for _, enrich := range options {
		enrich(&settings)
	}
	tr.Species = settings.race
	tr.dice = settings.dice
	if tr.dice == nil {
		tr.dice = dice.New()
	}
	//tr.CharSet.RollPreset(settings.race)

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
