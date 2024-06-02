package db

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	te := TravellerEntry{}
	te.Assets = NewAssets()
	te.Attributes = NewAttributes()
	if err := te.AquireAsset("STR"); err != nil {
		fmt.Println(err.Error())
	}
	te.ModifyAsset("STR", 5)

	te.AquireAttr("Human", "Have no distinct traits")

	bt, err := te.marshal()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bt))

}
