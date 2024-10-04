package basic

import (
	"fmt"

	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

func getBasicPrequisites(context profile.Profile, datatype string) (map[string]int, error) {
	prequisites := make(map[string]int)
	keys := []string{}
	switch datatype {
	default:
		return nil, fmt.Errorf("prequisites not defined for datatype '%v'", datatype)
	case profile.KEY_Size, profile.KEY_Pops:
		return prequisites, nil
	case profile.KEY_Atmo:
		keys = []string{profile.KEY_Size}
	case profile.KEY_Temperature:
		keys = []string{profile.KEY_Size, profile.KEY_Atmo}
	case profile.KEY_Hydr:
		keys = []string{profile.KEY_Size, profile.KEY_Atmo, profile.KEY_Temperature}
	case profile.KEY_Govr, profile.KEY_Port:
		keys = []string{profile.KEY_Pops}
	case profile.KEY_Laws:
		keys = []string{profile.KEY_Pops, profile.KEY_Govr}
	case profile.KEY_TL:
		keys = []string{profile.KEY_Port, profile.KEY_Size, profile.KEY_Atmo, profile.KEY_Hydr, profile.KEY_Pops, profile.KEY_Govr}

	}
	for _, key := range keys {
		val, err := prequisite(context, key)
		if err != nil {
			return prequisites, fmt.Errorf("failed to get prequisite[%v]: %v", key, err)
			fmt.Println(err.Error())
			val = 0
		}
		prequisites[key] = val
	}
	return prequisites, nil
}

func prequisite(ctx profile.Profile, key string) (int, error) {
	_, val, err := ctx.Ehex(key)
	return val, err
}
