package continuation

import (
	"errors"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

func Roll_PCR(ctx profile.Profile, dice *dice.Dicepool) (int, error) {
	pcr := 0
	err := errors.New("roll was not made")
	switch ctx.Field(profile.GENERATION_METHOD) {
	case method.Continuation:
		pcr, err = continuation_PCR(ctx, dice)
	}
	return pcr, err
}

func continuation_PCR(ctx profile.Profile, dice *dice.Dicepool) (int, error) {
	// prequisites, err := getContinuationPrequisites(ctx, profile.KEY_Pops_PCR)
	// if err != nil {
	// 	return 0, fmt.Errorf("prequisites aquisition failed: %v", err)
	// }
	// pops := sequisites[profile.KEY_Pops]

	return 9, nil
}
