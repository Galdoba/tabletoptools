package traveller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func (tr *Traveller) Marshal() ([]byte, error) {
	return json.MarshalIndent(tr, "", "  ")
}

func UnmarshalJSON(filename string) (*Traveller, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("can't open json file: %v", err.Error())
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("can't read json file: %v", err.Error())
	}
	tr := Traveller{}
	if err := json.Unmarshal(jsonData, &tr); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
		return nil, fmt.Errorf("can't unmarshal json file: %v", err.Error())
	}
	return &tr, nil
}
