package world

import (
	"fmt"
	"testing"

	"github.com/Galdoba/tabletoptools/pkg/dice"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/generation/method"
	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

func TestWorld(t *testing.T) {
	w := New("test world")
	w.SetGenerationMethod(method.Basic)
	err := w.GenerateMissingData(dice.New())
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(w)

	fmt.Println(w.Profile.Format(profile.UWP))
	w.SetGenerationMethod(method.Continuation)
	err = w.GenerateMissingData(dice.New())
	if err != nil {
		fmt.Println(err.Error())
	}
	w.Profile.List()
}
