package continuation

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

func getContinuationPrequisites(context profile.Profile, datatype string) (map[string]int, error) {
	prequisites := make(map[string]int)
	keys := []string{}
	switch datatype {
	default:
		return nil, fmt.Errorf("prequisites not defined for datatype '%v'", datatype)
	case profile.KEY_Pops_PCR:
		keys = []string{profile.KEY_Pops}
	}
	for _, key := range keys {
		val, err := prequisite(context, key)
		if err != nil {
			return prequisites, fmt.Errorf("failed to get prequisite[%v]: %v", key, err)
		}
		prequisites[key] = val
	}
	return prequisites, nil
}

func prequisite(ctx profile.Profile, key string) (int, error) {
	_, val, err := ctx.Ehex(key)
	return val, err
}
