package profile

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	pr := New(UWP)
	pr.SetValue(KEY_Starport, "C")
	pr.SetValue(KEY_Size, "A")
	pr.SetValue(KEY_Atmo, "6")
	pr.SetValue(KEY_Hydr, "A")
	//pr.SetValue(KEY_Pops, "6")
	pr.SetValue(KEY_Govr, "4")
	pr.SetValue(KEY_Laws, "3")
	pr.SetValue(KEY_TL, "9")
	fmt.Println(pr.Profile())
}
