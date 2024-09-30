package preset

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	err := GenerateAll()
	if err != nil {
		fmt.Println(err.Error())
	}

	pres, err := Load("vargr")
	fmt.Println(pres.Craracteristics[1])
	fmt.Println(err)
}
