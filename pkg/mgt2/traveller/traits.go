package traveller

type Trait struct {
	Name        string `json:"Name"`
	Description string `json:"Description,omitempty"`
}

func NewTrait(name string, descr ...string) Trait {
	traitMap := make(map[string]string)
	traitMap["PC"] = "Player control"
	if desc, ok := traitMap[name]; ok {
		return Trait{name, desc}
	}
	for _, desc := range descr {
		return Trait{name, desc}
	}
	return Trait{name, "[no description]"}
}
