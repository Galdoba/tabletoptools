package skill

import "fmt"

type skill struct {
	name           string
	parent         string
	specialities   []string
	effectiveScore int
	maxScore       int
}

func New(name string, options ...SkillOption) (*skill, error) {
	sk := skill{}
	sk.specialities, sk.parent = specialitiesAndParent(name)
	if sk.parent == "bad name" {
		return nil, fmt.Errorf("skill '%v' has bad key or unimplemented", name)
	}
	sk.name = longName(name, sk.parent)

	settings := defaultOptions()
	for _, set := range options {
		set(&settings)
	}
	sk.maxScore = settings.maxScore
	sk.effectiveScore = settings.effectiveScore
	if len(sk.specialities) > 0 {
		sk.maxScore = 0
		sk.effectiveScore = 0
	}
	return &sk, nil
}

func longName(name, parent string) string {
	switch parent {
	case "":
		return name
	default:
		return fmt.Sprintf("%v (%v)", parent, name)
	}
}
