package world

import (
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/ehex"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/tradecode"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/world/location"
	"github.com/Galdoba/tabletoptools/pkg/mgt2/world/name"
)

type World struct {
	Dice       Roller
	Name       name.WorldName
	Location   location.Location
	Profile    profile.Profile
	Bases      []string
	TradeCodes []string
}

type Roller interface {
	Sroll(string) int
}

func New(opts ...WorldOption) (*World, error) {
	w := World{}
	worldData := defaultOption()

	for _, enrich := range opts {
		enrich(&worldData)
	}
	w.Name = name.New(worldData.name)
	w.Location = location.NewSectorCoordinates(worldData.location)
	knownProfile := make(map[string]string)
	knownProfile = worldData.profileKeys
	prf := profile.New(profile.UWP)
	for k, v := range knownProfile {
		if err := prf.SetValue(k, v); err != nil {
			return nil, fmt.Errorf("profile info injection failed: %v", err)
		}
	}

	for k := range worldData.confirmedPresence {
		switch k {
		case "N", "M", "S", "C", "R", "W", "D":
			w.Bases = append(w.Bases, k)
		case "Fr", "Co", "Te", "Ho", "Bo":
			prf.SetValue(profile.KEY_Temp, k)
		}
	}
	prf.GenerateMissingData()
	w.Profile = prf
	w.GenerateBases()

	w.WipeUnsustainable()

	w.TradeCodes = tradecode.Designate(w.Profile)

	return &w, nil
}

func (w *World) WipeUnsustainable() error {
	prf := w.Profile
	tl_actial := ehex.ValueOf(prf.GetValue(profile.KEY_TL))
	tl_minimum := 0
	atmo := ehex.ValueOf(prf.GetValue(profile.KEY_Atmo))
	switch atmo {
	case 0, 1:
		tl_minimum = 8
	case 2, 3:
		tl_minimum = 5
	case 4, 7, 9:
		tl_minimum = 3
	case 10:
		tl_minimum = 8
	case 11:
		tl_minimum = 9
	case 12:
		tl_minimum = 10
	case 13, 14:
		tl_minimum = 5
	case 15:
		tl_minimum = 8
	}
	if tl_actial < tl_minimum {
		fmt.Println(prf, "Unsustainable!!!")
		prf.SetValue(profile.KEY_Pops, "0")
		prf.SetValue(profile.KEY_Govr, "0")
		prf.SetValue(profile.KEY_Laws, "0")
		prf.SetValue(profile.KEY_Starport, "?")
		if err := prf.GenerateMissingData(); err != nil {
			return err
		}
	}
	switch prf.GetValue(profile.KEY_Starport) {
	case "X":
		prf.SetValue(profile.KEY_TL, "0")
	default:

	}
	w.Profile = prf
	return nil
}

func (w *World) GenerateHighport() {
	hpdm := 0
	switch w.Profile.GetValue(profile.KEY_Pops) {
	case "9", "A", "B", "C":
		hpdm += 1
	case "0", "1", "2", "3", "4", "5", "6":
		hpdm += -1
	}
	switch w.Profile.GetValue(profile.KEY_TL) {
	case "9", "A", "B":
		hpdm += 1
	case "C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P":
		hpdm += 2
	}
	tn := 0
	switch w.Profile.GetValue(profile.KEY_Starport) {
	case "A":
		tn = 6
	case "B":
		tn = 8
	case "C":
		tn = 10
	case "D":
		tn = 12
	}
	r := w.Dice.Sroll("2d6") + hpdm
	if tn > r {
		return
	}
	w.Bases = append(w.Bases, "H")
}
