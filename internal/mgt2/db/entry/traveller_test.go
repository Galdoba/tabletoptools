package entry

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	te := Entry{}
	te.Assets = NewAssets()
	te.Attributes = NewAttributes()
	if err := te.AquireAsset("STR"); err != nil {
		fmt.Println(err.Error())
	}
	te.ModifyAsset("STR", 5)

	te.AquireAttr("Human", "Have no distinct traits")

	bt, err := ToBytes(te)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bt))

}
