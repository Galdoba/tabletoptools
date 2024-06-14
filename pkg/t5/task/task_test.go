package task

import (
	"fmt"
	"testing"

	"github.com/Galdoba/tabletoptools/pkg/dice"
)

func TestResolution(t *testing.T) {
	dp := dice.New()
	for dice := 1; dice <= 10; dice++ {
		for tn := 1; tn <= 20; tn++ {
			result := dp.Roll(fmt.Sprintf("%vd6", dice)).Result()

			resolve := newResolution(tn, result)
			fmt.Printf("tn:%v result:%v Outcome:%v stupid=%v    \n", tn, result, resolve.Outcome, resolve.SpectacularlyStupid)
		}
		//make result
		//make tn
	}
}
