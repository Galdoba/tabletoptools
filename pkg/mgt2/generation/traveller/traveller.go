package traveller

import (
	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/traveller/benefit"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/traveller/characteristic"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/traveller/key"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/traveller/skill"
)

type Traveller struct {
	Streight  int `json:"Streight"`
	Dexterity int `json:"Dexterity"`
	Endurace  int `json:"Endurace"`
	Intellect int `json:"Intellect"`
	Education int `json:"Education"`
	Social    int `json:"Social"`
}

type TravellerModel struct {
	dice           *dice.Dicepool
	generationStep int
	charSet        *characteristic.CharSet
	skillSet       *skill.SkillSet
	age            int
}

func (trv *TravellerModel) gainBonus(k string) error {
	bonus := benefit.Gain(k)
	for _, take := range bonus.Take {
		switch take.Type {
		case key.RESOURCE_CHARACTERISTIC:
			trv.charSet.Modify(take.Key, take.Value)
		}
	}
	selected := trv.selectFrom(bonus.Options)
	if selected == nil {
		return nil
	}
	switch selected.(type) {
	case benefit.Resource:
		take := selected.(benefit.Resource)
		switch take.Type {
		case key.RESOURCE_CHARACTERISTIC:
			trv.charSet.Modify(take.Key, take.Value)
		}

	}
	return nil
}

func (trv *TravellerModel) applyBonus(resource benefit.Resource) error {
	switch resource.Type {
	case key.RESOURCE_CHARACTERISTIC:
		return trv.charSet.Modify(resource.Key, resource.Value)
	}
	return nil
}

func (trv *TravellerModel) selectFrom(opts ...interface{}) interface{} {
	return opts[0]
}

/*
Steps:
Characteristics
	roll each
	calculate mods
Background
	choose background skills
	choose homeworld
Pre-Career
	choose university/academy
	gain skills
	roll event
	roll graduation
		+gain graduation
StartTerm
	choose career
	roll qualification
		-choose drifter
			-roll draft
	basic training
		-roll skill table
Survival
	roll survival
		-roll mishap
		-end career by default
Event
	roll event
Advancement
	commision if possible
	roll advancement
Continue
	increase age
	choose to continue career
		-choose change branch
			-go to StartTerm
Muster_Out
	roll cash/benefits
	resolve debt/pension
Skill_Package

Adventure

*/
