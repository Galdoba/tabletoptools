package career

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Career struct {
	Career         string        `yaml:"Career"`
	Assignments    []Assignment  `yaml:"Assignments"`
	Qualification  Qualification `yaml:"QUALIFICATION"`
	Progress       Progress      `yaml:"CAREER PROGRESS"`
	Commission     Qualification `yaml:"COMMISSION,omitempty"`
	MusteringOut   MusteringOut  `yaml:"MUSTERING OUT BENEFITS"`
	TrainingTables SkillTables   `yaml:"TRAINING AND SKILLS"`
	Ranks          RankTables    `yaml:"RANKS AND BONUSES"`
	Mishaps        MishapInfo    `yaml:"MISHAPS"`
	Events         EventInfo     `yaml:"EVENTS"`
}

type Assignment struct {
	Assignment  string `yaml:"branch"`
	Description string `yaml:"description,omitempty"`
}

type Qualification struct {
	Roll string            `yaml:"roll"`
	Mods map[string]string `yaml:"modifiers,omitempty"`
}

type Progress struct {
	Survival    map[string]string `yaml:"survival"`
	Advancement map[string]string `yaml:"advancement"`
}

type MusteringOut struct {
	Cash     map[int]string `yaml:"cash"`
	Benefits map[int]string `yaml:"benefits"`
}

type SkillTables struct {
	Tables map[string][]Bonus `yaml:"training tables"`
}

type Bonus struct {
	Roll        int      `yaml:"1D,omitempty"`
	BonusAll    []string `yaml:"take,omitempty"`
	BonusChoose []string `yaml:"choose,omitempty"`
}

type RankTables struct {
	ByAssignment map[string][]RankInfo `yaml:"rank tables"`
}

type RankInfo struct {
	Rank     int    `yaml:"rank"`
	Position string `yaml:"position"`
	Bonus    string `yaml:"skill or bonus,omitempty"`
}

type MishapInfo struct {
	Mishap map[int]string `yaml:"mishap tables,omitempty"`
}

type EventInfo struct {
	Events map[int]Event `yaml:"event tables,omitempty"`
}

type Event struct {
	Description string `yaml:"description"`
}

func DummyCareer() Career {
	crr := Career{}
	crr.Career = "ARMY"
	crr.Assignments = []Assignment{
		{
			Assignment:  "Support",
			Description: "You are engineer...",
		},
		{
			Assignment:  "Infantry",
			Description: "You are solder...",
		},
		{
			Assignment:  "Cavalry",
			Description: "You are tankist...",
		},
	}
	crr.Qualification = Qualification{
		Roll: "END 5+",
		Mods: map[string]string{
			"for every previous career": "DM-1",
			"you are aged 30+":          "DM-2",
		},
	}
	crr.Progress = Progress{
		Survival: map[string]string{
			crr.Assignments[0].Assignment: "END 5+",
			crr.Assignments[1].Assignment: "STR 6+",
			crr.Assignments[2].Assignment: "DEX 7+",
		},
		Advancement: map[string]string{
			crr.Assignments[0].Assignment: "EDU 7+",
			crr.Assignments[1].Assignment: "EDU 6+",
			crr.Assignments[2].Assignment: "INT 5+",
		},
	}
	crr.Commission = Qualification{
		Roll: "SOC 8+",
		Mods: map[string]string{
			"graduation with honors: military academy": "AUTO",
		},
	}
	crr.Ranks = RankTables{
		ByAssignment: map[string][]RankInfo{
			crr.Assignments[0].Assignment: []RankInfo{
				{0, "Private", "Gun Combat 1"},
				{1, "Lance Copral", "Recon 1"},
				{2, "Copral", "-"},
			},
		},
	}
	crr.TrainingTables.Tables = make(map[string][]Bonus)
	crr.TrainingTables.Tables[crr.Assignments[0].Assignment] = []Bonus{
		{
			Roll:        1,
			BonusAll:    []string{"Mechanic"},
			BonusChoose: []string{},
		},
		{
			Roll:        2,
			BonusAll:    []string{},
			BonusChoose: []string{"Drive", "Flyer"},
		},
	}

	bt, err := yaml.Marshal(&crr)
	fmt.Println(err)
	fmt.Println(string(bt))
	return crr
}
