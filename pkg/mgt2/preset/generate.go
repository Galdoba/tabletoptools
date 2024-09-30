package preset

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Galdoba/devtools/stdpath"
	. "github.com/Galdoba/tabletoptools/pkg/mgt2/key"
)

func GenerateAll() error {
	errors := []error{}
	app := "mgt2"
	stdpath.SetAppName(app)
	dir := stdpath.ProgramDir("assets", "species", "presets")
	os.MkdirAll(dir, 0666)

	name := "vilani"
	file := dir + name + ".json"
	if f, err := os.Create(file); err == nil {
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
		if _, err := f.Write(bt); err != nil {
			return fmt.Errorf("write file failed: %v", err)
		}
	} else {
		errors = append(errors, err)
	}

	name = "aslan"
	file = dir + name + ".json"
	if f, err := os.Create(file); err == nil {
		prst := TravellerPreset{
			Craracteristics: []*CharacteristicPreset{
				{Name: CHAR_NAME_STR, Dices: 2, Mod: 2},
				{Name: CHAR_NAME_DEX, Dices: 2, Mod: -2},
				{Name: CHAR_NAME_END, Dices: 2, Mod: 0},
				{Name: CHAR_NAME_INT, Dices: 2, Mod: 0},
				{Name: CHAR_NAME_EDU, Dices: 2, Mod: 0},
				{Name: CHAR_NAME_TER, Dices: 2, Mod: 0},
			},
		}
		bt, err := json.MarshalIndent(&prst, "", "  ")
		if err != nil {
			return err
		}
		if _, err := f.Write(bt); err != nil {
			return fmt.Errorf("write file failed: %v", err)
		}
	} else {
		errors = append(errors, err)
	}

	name = "aslan_imperial"
	file = dir + name + ".json"
	if f, err := os.Create(file); err == nil {
		prst := TravellerPreset{
			Craracteristics: []*CharacteristicPreset{
				{Name: CHAR_NAME_STR, Dices: 2, Mod: 2},
				{Name: CHAR_NAME_DEX, Dices: 2, Mod: -2},
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
		if _, err := f.Write(bt); err != nil {
			return fmt.Errorf("write file failed: %v", err)
		}
	} else {
		errors = append(errors, err)
	}

	name = "vargr_imperial"
	file = dir + name + ".json"
	if f, err := os.Create(file); err == nil {
		prst := TravellerPreset{
			Craracteristics: []*CharacteristicPreset{
				{Name: CHAR_NAME_STR, Dices: 2, Mod: -2},
				{Name: CHAR_NAME_DEX, Dices: 2, Mod: 1},
				{Name: CHAR_NAME_END, Dices: 2, Mod: -1},
				{Name: CHAR_NAME_INT, Dices: 2, Mod: 0},
				{Name: CHAR_NAME_EDU, Dices: 2, Mod: 0},
				{Name: CHAR_NAME_SOC, Dices: 2, Mod: 0},
			},
		}
		bt, err := json.MarshalIndent(&prst, "", "  ")
		if err != nil {
			return err
		}
		if _, err := f.Write(bt); err != nil {
			return fmt.Errorf("write file failed: %v", err)
		}
	} else {
		errors = append(errors, err)
	}

	name = "vargr"
	file = dir + name + ".json"
	if f, err := os.Create(file); err == nil {
		prst := TravellerPreset{
			Craracteristics: []*CharacteristicPreset{
				{Name: CHAR_NAME_STR, Dices: 2, Mod: -2},
				{Name: CHAR_NAME_DEX, Dices: 2, Mod: 1},
				{Name: CHAR_NAME_END, Dices: 2, Mod: -1},
				{Name: CHAR_NAME_INT, Dices: 2, Mod: 0},
				{Name: CHAR_NAME_EDU, Dices: 2, Mod: 0},
				{Name: CHAR_NAME_CHA, Dices: 2, Mod: 0},
			},
		}
		bt, err := json.MarshalIndent(&prst, "", "  ")
		if err != nil {
			return err
		}
		if _, err := f.Write(bt); err != nil {
			return fmt.Errorf("write file failed: %v", err)
		}
	} else {
		errors = append(errors, err)
	}
	return combinedErr(errors...)
}

func combinedErr(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}
	text := "Errors:\n"
	for _, err := range errs {
		if err == nil {
			continue
		}
		text += err.Error() + "\n"
	}
	return fmt.Errorf(text)
}
