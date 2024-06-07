package entry

import (
	"encoding/json"
	"fmt"
)

type Entry struct {
	GenerationSeed string         `json:"Seed,omitempty"` //сид генератора по которому все вертится
	Assets         *AssetData     `json:"Assets,omitempty"`
	Attributes     *AttributeData `json:"Attributes,omitempty"` //все то что сводится к описательному значению (оно всегда есть), и имеет перманентный эффект от наличия/отсуствия (трэйты, контакты)
	EventPath      []EventDescr   `json:"Events,omitempty"`     //все то что сводится к событию.

}

func NewEntry() *Entry {
	e := Entry{}
	e.Assets = NewAssets()
	e.Attributes = NewAttributes()
	return &e
}

func (te *Entry) KeyIsBad(k string) bool {
	if _, ok := te.Assets.ByKey[k]; ok {
		return true
	}
	if _, ok := te.Attributes.ByKey[k]; ok {
		return true
	}
	return false
}

// все то что сводится к числовому значению, которое может браться и меняться - не опускается ниже 0 (характеристики, скиллы, предметы, деньги)
type AssetData struct {
	ByKey map[string]uint `json:"By Key"`
}

type AssetHandler interface {
	Asset(string) int
	ModifyAsset(key string, by int) error
	AquireAsset(string) error
	RemoveAsset(string)
}

func NewAssets() *AssetData {
	a := AssetData{}
	a.ByKey = make(map[string]uint)
	return &a
}

func (te *Entry) Asset(k string) int {
	if k == "" {
		return -2
	}
	if _, ok := te.Assets.ByKey[k]; !ok {
		return -1
	}
	return int(te.Assets.ByKey[k])
}

func (te *Entry) ModifyAsset(k string, by int) error {
	if k == "" {
		return fmt.Errorf("can't use asset with null key")
	}
	if _, ok := te.Assets.ByKey[k]; !ok {
		return fmt.Errorf("asset with key '%v' does not exist", k)
	}
	val := int(te.Assets.ByKey[k]) + by
	if val < 0 {
		return fmt.Errorf("can't modify asset below 0")
	}
	te.Assets.ByKey[k] = uint(val)
	return nil
}

func (te *Entry) AquireAsset(k string) error {
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

func (te *Entry) RemoveAsset(k string) {
	delete(te.Assets.ByKey, k)
}

/**/

func NewAttributes() *AttributeData {
	ad := AttributeData{}
	ad.ByKey = make(map[string]string)
	return &ad
}

type AttributeData struct {
	ByKey map[string]string
}

type AttributeHandler interface {
	AquireAttr(string, string) error
	RemoveAttr(string) error
	DescribeAttr(string) string
}

func (te *Entry) AquireAttr(k, descr string) error {
	if te.KeyIsBad(k) {
		return fmt.Errorf("can't add attribute with non-unique key '%s'", k)
	}
	if k == "" {
		return fmt.Errorf("can't add attribute with null key")
	}
	te.Attributes.ByKey[k] = descr
	return nil
}

func (te *Entry) RemoveAttr(k string) {
	delete(te.Attributes.ByKey, k)
}

func (te *Entry) DescribeAttr(k string) string {
	if val, ok := te.Attributes.ByKey[k]; ok {
		return val
	}
	return ""
}

func ToBytes(te Entry) ([]byte, error) {
	return json.MarshalIndent(&te, "", "  ")
}

func FromBytes(bt []byte) (Entry, error) {
	entr := Entry{}
	err := json.Unmarshal(bt, &entr)
	if err != nil {
		return Entry{}, err
	}
	return entr, nil
}
