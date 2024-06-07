package traveller

import (
	"errors"
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/options"
)

type PersonalData struct {
	Name      string  `json:"Name,omitempty"`
	Species   string  `json:"Species,omitempty"`
	Age       int     `json:"Age,omitempty"`
	Homeworld string  `json:"Homeworld,omitempty"`
	Traits    []Trait `json:"Traits,omitempty"`
}

var ErrBadOptionType = errors.New("bad option type")

func newPersonal(opts ...options.Option) (*PersonalData, error) {
	pd := PersonalData{}
	if err := pd.injectoptions(opts...); err != nil {
		return nil, fmt.Errorf("options injections failed: %v", err.Error())
	}
	return &pd, nil
}

func (pd *PersonalData) injectoptions(opts ...options.Option) error {
	for _, opt := range opts {
		switch opt.Key {
		default:
			continue
		case options.PD_NAME:
			val, ok := opt.Val.(string)
			if !ok {
				return fmt.Errorf("bad option type: '%v'", opt.Key)
			}
			pd.Name = val
		case options.PD_AGE:
			val, ok := opt.Val.(int)
			if !ok {
				return fmt.Errorf("bad option type: '%v'", opt.Key)
			}
			pd.Age = val
		case options.PD_SPECIE:
			val, ok := opt.Val.(string)
			if !ok {
				return fmt.Errorf("bad option type: '%v'", opt.Key)
			}
			pd.Species = val
		case options.PD_HOMEWORLD:
			val, ok := opt.Val.(string)
			if !ok {
				return fmt.Errorf("bad option type: '%v'", opt.Key)
			}
			pd.Name = val
		case options.PD_TRAITS:
			val, ok := opt.Val.([]string)
			if !ok {
				return fmt.Errorf("bad option type: '%v'", opt.Key)
			}
			for _, trt := range val {
				pd.Traits = append(pd.Traits, NewTrait(trt))
			}
		}
	}
	return nil
}
