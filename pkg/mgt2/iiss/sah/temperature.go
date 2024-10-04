package sah

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method/basic"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

func (sah *SAH) RollTemperature(ctx profile.Profile, dice *dice.Dicepool) error {
	temp, err := basic.RollTemperature(ctx, dice)
	if err != nil {
		return fmt.Errorf("temperature roll failed: %v", err)
	}
	sah.Temperature = temp
	return ctx.Inject(profile.KEY_Temperature, temp)
}
