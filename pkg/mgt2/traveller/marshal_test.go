package traveller

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	tr := &Traveller{}
	tr.Characteristics.UPP = "789ABC"
	tr.Characteristics.STR = "7"
	tr.Characteristics.SOC = "C"
	tr.Personal.Name = "John Wick"
	tr.Personal.Species = "Human"
	tr.Personal.Homeworld = "Earth"

	tr.Personal.Age = 42
	tr.Personal.Traits = []Trait{{Name: "PC", Description: "Played by human"}}
	bt, err := tr.Marshal()
	if err != nil {
		t.Errorf("marshal err := %v", err.Error())
	}
	fmt.Println(string(bt))
}
