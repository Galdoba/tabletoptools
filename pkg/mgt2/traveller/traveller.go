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
	err := fmt.Errorf("no generation comenced")
	tr.Characteristics, err = characteristic.NewCharSet(characteristic.CoreChars()...)
	if err != nil {
		return nil, fmt.Errorf("can't create characteristic block: %v", err.Error())
	}
	tr.Characteristics.Roll(dice, opts...)
	tr.Personal, err = newPersonal(dice, opts...)
	if err != nil {
		return nil, fmt.Errorf("can't create personal block: %v", err.Error())
	}
	return &tr, nil
}
