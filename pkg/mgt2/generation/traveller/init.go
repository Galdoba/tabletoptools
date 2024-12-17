package traveller

import (
	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/traveller/characteristic"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/traveller/skill"
)

func StartModel(seed int64) (*TravellerModel, error) {
	trv := TravellerModel{}
	trv.dice = dice.New(seed)
	trv.charSet = characteristic.NewSet()
	trv.skillSet = skill.NewSet()
	trv.generationStep = 1

	return &trv, nil
}
