package traveller

import (
	"fmt"
	"testing"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/key"
)

func TestDev(t *testing.T) {
	dice := dice.New().SetSeed("15")

	tr, err := New(WithRace("vargr"), WithDice(dice))
	if err != nil {
		fmt.Println(err.Error())
	}
	tr.Name = "Trav 1"
	if err := tr.SkillSet.Train(key.SKL_Advocate); err != nil {
		fmt.Println("-", err.Error())
	}
	if err := tr.SkillSet.Train(key.SKL_Astrogation); err != nil {
		fmt.Println("-", err.Error())
	}

	// fmt.Println(tr.CharSet.ByCode["C1"].Encode())
	// fmt.Println(tr.CharSet.ByCode["C1"].Mod())
	fmt.Println(tr.RollCharacteristics())
	bt, err := tr.Marshal()
	fmt.Println(string(bt))
}
