package benefit

import "github.com/Galdoba/tabletoptools/pkg/mgt2/generation/traveller/key"

type Bonus struct {
	Take    []Resource
	Options []Resource
}

type Resource struct {
	Type  string
	Key   string
	Value int
}

func Gain(k string) Bonus {

	switch k {
	case "INT +1":
		return Bonus{
			Take: []Resource{
				{key.RESOURCE_CHARACTERISTIC, key.C4_INT, 1},
			},
		}
	}
	return Bonus{}
}
