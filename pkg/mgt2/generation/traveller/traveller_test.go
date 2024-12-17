package traveller

import (
	"fmt"
	"testing"
)

func TestTrav(t *testing.T) {
	trv, err := StartModel(16)
	fmt.Println(err, trv)
}
