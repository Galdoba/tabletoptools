package entity

import (
	"fmt"
	"testing"
)

func TestCreation(t *testing.T) {
	e := New("NPC", true, "Sophont", "Human", "Criminal")
	fmt.Println(e.path())
	e2 := New("NPC", false, "Sophont", "Human", "Criminal")
	fmt.Println(e2.path())

}
