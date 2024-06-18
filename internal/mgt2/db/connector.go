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
	chrSet := tr.Characteristics
	for key, chr := range chrSet.CHR {
		if err := te.AquireAsset(key.Abbreviation); err != nil {
			return nil, err
		}
		if err := te.ModifyAsset(key.Abbreviation, chr.Current); err != nil {
			return nil, err
		}
	}
	return te, nil
}

// func ToStruct(e *entry.Entry) (*traveller.Traveller, error) {
// 	tr := traveller.Traveller{}
// }

/*
Load(seed string) (*mgt2.Traveller, error)
Save(*mgt2.Traveller) error

*/
