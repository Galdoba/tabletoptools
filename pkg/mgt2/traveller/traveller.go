package traveller

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/characteristic"
	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/options"
)

type Traveller struct {
	Characteristics *characteristic.Set `json:"Characteristics"`
	Personal        *PersonalData       `json:"Personal Info"`
	Finance         FinanceData         `json:"Finance Info"`
}

type option struct {
	key string
	val interface{}
}

func Option(key string, val interface{}) option {
	return option{key, val}
}

func New(dice *dice.Dicepool, opts ...options.Option) (*Traveller, error) {
	tr := Traveller{}
	err := fmt.Errorf("no created")
	tr.Characteristics, err = characteristic.NewCharSet()
	if err != nil {
		return nil, fmt.Errorf("can't create charset block: %v", err.Error())
	}
	tr.Personal, err = newPersonal(opts...)
	if err != nil {
		return nil, fmt.Errorf("can't create personal block: %v", err.Error())
	}
	return &tr, err
}

func (tr *Traveller) Roll(dice Roller, opts ...options.Option) error {
	err := fmt.Errorf("no generation comenced")
	if err = tr.Characteristics.Roll(dice, opts...); err != nil {
		return fmt.Errorf("can't roll characteristic block: %v", err.Error())
	}
	return nil
}

type Roller interface {
	Sroll(string) int
}
