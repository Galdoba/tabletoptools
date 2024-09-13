package characteristic

import (
	"fmt"
	"testing"

	"github.com/Galdoba/tabletoptools/pkg/dice"
)

func TestSet(t *testing.T) {
	set, err := NewSet().Vargr()
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(set)
	dice := dice.New()
	for k, v := range set.byCode {
		v.Roll(dice)
		fmt.Println(k, v)
		fmt.Println(v.Encode())
		for i := 5; i > 0; i-- {
			fmt.Printf("train %v (curent max is %v)\n", v.name, v.maxScore)
			if err := v.Train(); err != nil {
				fmt.Println(err)
			}
		}
	}
}
