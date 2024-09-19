package characteristic

import (
	"fmt"

	. "github.com/Galdoba/tabletoptools/pkg/mgt2/key"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/preset"
)

type Set struct {
	ByCode map[string]*characteristic
}

func NewSet() *Set {
	cs := Set{}
	cs.ByCode = make(map[string]*characteristic)
	return &cs
}

func (cs *Set) ImportPreset(presetName string) error {
	if len(cs.ByCode) != 0 {
		return fmt.Errorf("non empty set received")
	}
	presetValues, err := preset.Load(presetName)
	if err != nil {
		return fmt.Errorf("preset loading failed: %v", err)
	}
	for _, imported := range presetValues.Craracteristics {
		chr, err := New(imported.Name, CreationDice(imported.Dices), CreationMod(imported.Mod))
		if err != nil {
			return fmt.Errorf("characteristic '%v' creation failed: %v", imported.Name, err)
		}
		cs.ByCode[chr.code] = chr

	}
	if len(cs.ByCode) != 6 {
		return fmt.Errorf("failed to create new set")
	}
	return nil
}

func (cs *Set) Human() (*Set, error) {
	if len(cs.ByCode) != 0 {
		return nil, fmt.Errorf("non empty set received")
	}
	for _, name := range []string{CHAR_NAME_STR, CHAR_NAME_DEX, CHAR_NAME_END, CHAR_NAME_INT, CHAR_NAME_EDU, CHAR_NAME_SOC} {
		chr, err := New(name)
		switch name {
		default:
		}
		if err != nil {
			return nil, fmt.Errorf("characteristic '%v' creation failed: %v", name, err)
		}
		cs.ByCode[chr.code] = chr
	}
	if len(cs.ByCode) != 6 {
		return nil, fmt.Errorf("failed to create new set")
	}
	return cs, nil
}

func (cs *Set) Aslan() (*Set, error) {
	if len(cs.ByCode) != 0 {
		return nil, fmt.Errorf("non empty set received")
	}
	for _, name := range []string{CHAR_NAME_STR, CHAR_NAME_DEX, CHAR_NAME_END, CHAR_NAME_INT, CHAR_NAME_EDU, CHAR_NAME_TER} {
		chr, err := New(name)
		switch name {
		case CHAR_NAME_STR:
			chr, err = New(name, CreationMod(2))
		case CHAR_NAME_DEX:
			chr, err = New(name, CreationMod(-2))
		default:
		}
		if err != nil {
			return nil, fmt.Errorf("characteristic '%v' creation failed: %v", name, err)
		}
		cs.ByCode[chr.code] = chr
	}
	if len(cs.ByCode) != 6 {
		return nil, fmt.Errorf("failed to create new set")
	}
	return cs, nil
}

func (cs *Set) Vargr() (*Set, error) {
	if len(cs.ByCode) != 0 {
		return nil, fmt.Errorf("non empty set received")
	}
	for _, name := range []string{CHAR_NAME_STR, CHAR_NAME_DEX, CHAR_NAME_END, CHAR_NAME_INT, CHAR_NAME_EDU, CHAR_NAME_CHA} {
		chr, err := New(name)
		switch name {
		case CHAR_NAME_STR:
			chr, err = New(name, CreationMod(-1))
		case CHAR_NAME_DEX:
			chr, err = New(name, CreationMod(1))
		case CHAR_NAME_END:
			chr, err = New(name, CreationMod(-1))
		default:
		}
		if err != nil {
			return nil, fmt.Errorf("characteristic '%v' creation failed: %v", name, err)
		}
		cs.ByCode[chr.code] = chr
	}
	if len(cs.ByCode) != 6 {
		return nil, fmt.Errorf("failed to create new set")
	}
	return cs, nil
}

func (cs *Set) Map() map[string]string {
	chrMap := make(map[string]string)
	for k, chr := range cs.ByCode {
		chrMap[k] = chr.Encode()
	}
	return chrMap
}

func UnMap(chrMap map[string]string) *Set {
	set := NewSet()
	for _, v := range chrMap {
		chr, err := Decode(v)
		if err != nil {
			panic(err)
		}
		set.ByCode[chr.code] = chr
	}
	return set
}

/*
character.Characteristic("C1") *characteristic

character.CharSet, err := characteristic.NewSet().Human()
*/
