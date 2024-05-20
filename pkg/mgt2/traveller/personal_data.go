package traveller

type PersonalData struct {
	Name      string  `json:"Name,omitempty"`
	Species   string  `json:"Species,omitempty"`
	Age       int     `json:"Age,omitempty"`
	Homeworld string  `json:"Homeworld,omitempty"`
	Traits    []Trait `json:"Traits,omitempty"`
}
