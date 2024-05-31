package db

import (
	"encoding/json"
	"fmt"
)

type TravellerEntry struct {
	GenerationSeed string            `json:"Seed"` //сид генератора по которому все вертится
	Assets         *AssetData        `json:"Assets"`
	EventPath      []EventDescr      `json:"Events"`          //все то что сводится к событию.
	Possesions     map[string]string `json:"Items,omitempty"` //все то что сводится к описательному значению (оно всегда есть), и имеет перманентный эффект от наличия/отсуствия (вещи, контакты)
}

type udk struct {
	key string
}

func (te *TravellerEntry) KeyIsBad(k string) bool {
	inAsset, inPossesions := false, false
	_, inAsset = te.Assets.ByKey[k]
	_, inPossesions = te.Possesions[k]
	if inAsset || inPossesions {
		return false
	}
	return true
}

//все то что сводится к числовому значению, которое может браться и меняться - не опускается ниже 0 (характеристики, скиллы, таланты, деньги? )
type AssetData struct {
	ByKey map[string]uint `json:"By Key"`
}

type AssetHandler interface {
	Asset(string) int
	Modify(key string, by int) error
}

type PossesionHandler interface {
	Aquire(string) error
	Remove(string) error
	Property(string) string
}

func NewAssets() *AssetData {
	a := AssetData{}
	a.ByKey = make(map[string]uint)
	return &a
}

func (te *TravellerEntry) AddAsset(k string) error {
	if te.KeyIsBad(k) {
		return fmt.Errorf("can't add asset with non-unique key '%s'", k)
	}
	if k == "" {
		return fmt.Errorf("can't add asset with null key")
	}
	if _, ok := te.Assets.ByKey[k]; ok {
		return fmt.Errorf("asset with key '%v' already exist", k)
	}
	te.Assets.ByKey[k] = 0
	return nil
}

func (a *AssetData) Get(k string) int {
	if k == "" {
		return -2
	}
	if _, ok := a.ByKey[k]; !ok {
		return -1
	}
	return int(a.ByKey[k])
}

func (a *AssetData) Modify(k string, by int) error {
	if k == "" {
		return fmt.Errorf("can't use asset with null key")
	}
	if _, ok := a.ByKey[k]; !ok {
		return fmt.Errorf("asset with key '%v' does not exist", k)
	}
	val := int(a.ByKey[k]) + by
	if val < 0 {
		return fmt.Errorf("can't modify asset below 0")
	}
	a.ByKey[k] = uint(val)
	return nil
}

func (te *TravellerEntry) marshal() ([]byte, error) {
	return json.MarshalIndent(te, "", "  ")
}
