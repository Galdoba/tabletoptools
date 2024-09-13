package traveller

import (
	"fmt"
	"testing"

	"github.com/Galdoba/tabletoptools/pkg/dice"
)

func TestDev(t *testing.T) {
	tr := New()
	fmt.Println(tr.RollCharacteristics())
	fmt.Println(tr.CharSet.ByCode["C1"].Encode())
	fmt.Println(tr.CharSet.ByCode["C1"].Mod())
	dice := dice.New()
	fmt.Println(dice.Sroll("1d6+10"))
}
