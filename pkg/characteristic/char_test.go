package characteristic

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestChars(t *testing.T) {
	ch, err := defaultChar(Strength)
	if err != nil {
		t.Error(err)
	}
	bt, err := json.MarshalIndent(&ch, "", "  ")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(bt))
}
