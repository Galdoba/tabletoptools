package world

import (
	"fmt"
	"testing"

	"github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

func TestWorld(t *testing.T) {
	wrld, err := New(
		WithProfileData(profile.KEY_Temp, "Bo"),
		WithProfileData(profile.KEY_Atmo, "C"),
	)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(wrld.Profile.Profile(), wrld.TradeCodes)
}
