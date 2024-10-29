package dynasty

type Dynasty struct {
	Name                 string     `json:"Name"`
	Archetype            string     `json:"Archetype"`
	PowerBase            PowerBase  `json:"Power Base"`
	Managment            string     `json:"Managment"`
	Stats                *Stats     `json:"Stats"`
	Traits               *Traits    `json:"Traits"`
	Aptitudes            *Aptitudes `json:"Aptitudes"`
	FirstGenerationBonus string     `json:"First Generation Bonus"`
	BackgroundEvents     []string   `json:"Background Events"`
}

type Stats struct {
	Cleverness int `json:"Cleverness"`
	Greed      int `json:"Greed"`
	Loyalty    int `json:"Loyalty"`
	Militarism int `json:"Militarism"`
	Popularity int `json:"Popularity"`
	Scheming   int `json:"Scheming"`
	Tenacity   int `json:"Tenacity"`
	Tradition  int `json:"Tradition"`
}

type Traits struct {
	Culture            int `json:"Culture"`
	FiscalDefence      int `json:"Fiscal Defence"`
	Fleet              int `json:"Fleet"`
	Technology         int `json:"Technology"`
	TerritorialDefence int `json:"Territorial Defence"`
}

type Aptitudes struct {
	Acquisition     *int `json:"Acquisition"`
	Bureaucracy     *int `json:"Bureaucracy"`
	Conquest        *int `json:"Conquest"`
	Economics       *int `json:"Economics"`
	Entertain       *int `json:"Entertain"`
	Expression      *int `json:"Expression"`
	Hostility       *int `json:"Hostility"`
	Illicit         *int `json:"Illicit"`
	Intel           *int `json:"Intel"`
	Maintenance     *int `json:"Maintenance"`
	Polotics        *int `json:"Polotics"`
	Posturing       *int `json:"Posturing"`
	Propaganda      *int `json:"Propaganda"`
	PublicRelations *int `json:"Public Relations"`
	Recruit         *int `json:"Recruit"`
	Research        *int `json:"Research"`
	Sabotage        *int `json:"Sabotage"`
	Security        *int `json:"Security"`
	Tactical        *int `json:"Tactical"`
	Tutelage        *int `json:"Tutelage"`
}
