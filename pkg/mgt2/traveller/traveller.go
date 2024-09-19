package traveller

import (
	"encoding/json"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/traveller/characteristic"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/traveller/skill"
)

type Traveller struct {
	Name         string              `json:"Name"`
	Species      string              `json:"Species"`
	CharSet      *characteristic.Set `json:"-"`
	CharMap      map[string]string   `json:"Characteristics"`
	SkillSet     *skill.Set          `json:"-"`
	SkillMap     map[string]int      `json:"Skills"`
	dice         DiceRoller
	creationMode bool
}

func New(options ...Options) (*Traveller, error) {
	tr := Traveller{}
	tr.CharMap = make(map[string]string)
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
	err := tr.CharSet.ImportPreset(settings.race)
	if err != nil {
		return nil, err
	}

	return &tr, nil
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

func (tr *Traveller) Marshal() ([]byte, error) {
	tr.CharMap = tr.CharSet.Map()
	tr.SkillMap = tr.SkillSet.Map()
	bt, err := json.MarshalIndent(tr, "", "  ")
	return bt, err
}
