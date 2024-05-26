package traveller

import (
	"fmt"
	"testing"

	"github.com/Galdoba/tabletoptools/pkg/dice"
)

func TestString(t *testing.T) {
	fmt.Println("test render")
	// tr := Traveller{}
	// tr.Personal.Name = "Trav Name"
	// tr.Personal.Age = 18
	// tr.Personal.Homeworld = "Earth"
	// tr.Personal.Species = "Aslan"
	// tr.Personal.Traits = []Trait{
	// 	{
	// 		Name:        "Dewclaw",
	// 		Description: "Melee (natural): 1D+2 damage",
	// 	},
	// 	{
	// 		Name:        "Heightened Senses",
	// 		Description: "DM+1 to any Recon and Survival checks",
	// 	},
	// }
	// tr.Finance.Cash_On_Hand = 5000
	// tr.Characteristics.STR = "5"
	// tr.Characteristics.DEX = "6"
	// tr.Characteristics.END = "7"
	// tr.Characteristics.INT = "8"
	// tr.Characteristics.EDU = "9"
	// tr.Characteristics.SOC = "A"
	tr, err := New(dice.New())
	if err != nil {
		t.Error(err)
	}
	rn := tr.Render()
	fmt.Println(rn)
	fmt.Println("end test render")
}
