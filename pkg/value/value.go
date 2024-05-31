package value

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/options"
)

type val struct {
	ofType     string
	gameSystem int
	current    int
	maximum    int
	isPresent  bool
}

func New(opts ...options.Option) (*val, error) {
	v := val{}
	v.ofType = "undefined"
	v.gameSystem = -1
	for _, o := range opts {
		switch o.Key {
		default:
			continue
		case options.RULESET:
			if ruleset, ok := o.Val.(int); ok {
				v.gameSystem = ruleset
				continue
			}
		case options.VALUE_TYPE:
			if valType, ok := o.Val.(string); ok {
				v.ofType = valType
				continue
			}
		}
		return nil, fmt.Errorf("value creation: bad option: %v", o)
	}
	v.isPresent = true
	return &v, nil
}

type Value interface {
	DM() int
	ChangeBy(int) error
	Set(int) error
	Get() int
}
