package preset

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Galdoba/devtools/stdpath"
	. "github.com/Galdoba/tabletoptools/pkg/mgt2/key"
)

type TravellerPreset struct {
	Craracteristics []*CharacteristicPreset `json:"Characteristics"`
}

type CharacteristicPreset struct {
	Name  string `json:"Characteristic Name"`
	Dices int    `json:"Characteristic Creation Dices"`
	Mod   int    `json:"Characteristic Creation Mod"`
}

func Load(presetName string) (*TravellerPreset, error) {
	app := "mgt2"
	stdpath.SetAppName(app)
	path := stdpath.ProgramDir("assets", "species", "presets")
	file := path + presetName + ".json"
	bt, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("file reading failed: %v", err)
	}
	prst := &TravellerPreset{}
	err = json.Unmarshal(bt, prst)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling failed: %v", err)
	}
	return prst, err
}

func New(name string) error {
	name = "vilani"
	app := "mgt2"
	stdpath.SetAppName(app)
	dir := stdpath.ProgramDir("assets", "species", "presets")
	file := dir + name + ".json"
	os.MkdirAll(dir, 0666)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	prst := TravellerPreset{
		Craracteristics: []*CharacteristicPreset{
			{Name: CHAR_NAME_STR, Dices: 2, Mod: 0},
			{Name: CHAR_NAME_DEX, Dices: 2, Mod: 0},
			{Name: CHAR_NAME_END, Dices: 2, Mod: 0},
			{Name: CHAR_NAME_INT, Dices: 2, Mod: 0},
			{Name: CHAR_NAME_EDU, Dices: 2, Mod: 0},
			{Name: CHAR_NAME_SOC, Dices: 2, Mod: 0},
		},
	}
	bt, err := json.MarshalIndent(&prst, "", "  ")
	if err != nil {
		return err
	}
	f.Write(bt)
	return nil
}
