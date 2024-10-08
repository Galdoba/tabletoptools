package db

import (
	"github.com/Galdoba/tabletoptools/internal/mgt2/db/entry"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/traveller"
)

type DB struct {
	TravByKey map[string]entry.Entry
}

func FromStruct(tr *traveller.Traveller) (*entry.Entry, error) {
	te := entry.NewEntry()
	chrSet := tr.CharSet
	for key, chr := range chrSet.ByCode {
		if err := te.AquireAsset(key); err != nil {
			return nil, err
		}
		if err := te.ModifyAsset(key, chr.Score()); err != nil {
			return nil, err
		}
	}
	return te, nil
}

func ToStruct(e *entry.Entry) (*traveller.Traveller, error) {
	tr := traveller.Traveller{}
	return &tr, nil
}

/*
Load(seed string) (*mgt2.Traveller, error)
Save(*mgt2.Traveller) error

*/
