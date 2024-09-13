package characteristic

import "fmt"

type Set struct {
	ByCode map[string]*characteristic
}

func NewSet() *Set {
	cs := Set{}
	cs.ByCode = make(map[string]*characteristic)
	return &cs
}

func (cs *Set) Human() (*Set, error) {
	if len(cs.ByCode) != 0 {
		return nil, fmt.Errorf("non empty set received")
	}
	for _, name := range []string{STR, DEX, END, INT, EDU, SOC} {
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
	for _, name := range []string{STR, DEX, END, INT, EDU, TER} {
		chr, err := New(name)
		switch name {
		case STR:
			chr, err = New(name, CreationMod(2))
		case DEX:
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
	for _, name := range []string{STR, DEX, END, INT, EDU, CHA} {
		chr, err := New(name)
		switch name {
		case STR:
			chr, err = New(name, CreationMod(-1))
		case DEX:
			chr, err = New(name, CreationMod(1))
		case END:
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

/*
character.Characteristic("C1") *characteristic

character.CharSet, err := characteristic.NewSet().Human()
*/
