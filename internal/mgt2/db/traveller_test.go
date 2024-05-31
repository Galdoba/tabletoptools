package db

import (
	"fmt"
	"testing"
)

func TestT(t *testing.T) {
	te := TravellerEntry{}
	te.Assets = NewAssets()
	te.Assets.Add("STR")
	te.Assets.Modify("STR", 5)
	te.Assets.Modify("STR", -2)
	bt, err := te.marshal()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bt))

}
