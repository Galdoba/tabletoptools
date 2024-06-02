package event

import (
	"fmt"
	"testing"
)

func TestAction(t *testing.T) {
	sub := &subjStr{"World", "---"}
	// sub2 := &subjStr{"Dog", "Bark"}

	fun := speak

	act := NewAction(sub, fun, sub)
	fmt.Println(sub)
	act.Exec()
	fmt.Println(sub)
	act2 := NewAction(sub, bark, sub)
	act2.Exec()
	fmt.Println(sub)

}
