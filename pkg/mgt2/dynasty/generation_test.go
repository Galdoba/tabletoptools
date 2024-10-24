package dynasty

import (
	"fmt"
	"testing"

	"github.com/Galdoba/tabletoptools/pkg/dice"
)

func TestGeneration(t *testing.T) {
	for i := 0; i < 5; i++ {
		dice := dice.New().SetSeed(fmt.Sprintf("%v", i))
		gen := NewDynastyGenerator(dice, method_DiceRolling, true)
		d, err := gen.Generate()
		if err != nil {
			fmt.Println(err)
		}
		if d != nil {
			d.Name = fmt.Sprintf("Test %v", i)
			fmt.Println(d.String())
		}
	}

}
