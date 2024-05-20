package characteristic

type Characteristic struct {
	//Value          int //пойдет в значение
	//GeneticValue   string//пойдет в значение
	Name         string `json:"Name"`
	Position     string `json:"Position"`
	Abbreviation string `json:"Abbreviation"`
	GeneticCode  string `json:"GeneticCode"`
	Type         string `json:"Type"`
}
