package traveller

const (
	STR = "STREIGHT"
	DEX = "DEXTERITY"
	END = "ENDURANCE"
	INT = "INTELIGENCE"
	EDU = "EDUCATION"
	SOC = "SOCIAL"
	PSI = "PSIONICS"
	SAN = "SANITY"
)

type CharacteristicsData struct {
	STR string `json:"STREIGHT,omitempty"`
	DEX string `json:"DEXTERITY,omitempty"`
	END string `json:"ENDURANCE,omitempty"`
	INT string `json:"INTELIGENCE,omitempty"`
	EDU string `json:"EDUCATION,omitempty"`
	SOC string `json:"SOCIAL,omitempty"`
	PSI string `json:"PSIONICS,omitempty"`
	SAN string `json:"SANITY,omitempty"`
	UPP string `json:"Universal Profile,omitempty"`
}
