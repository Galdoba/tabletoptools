package characteristic

const (
	Strength        = "Strength"
	Dexterity       = "Dexterity"
	Agility         = "Agility"
	Grace           = "Grace"
	Endurance       = "Endurance"
	Stamina         = "Stamina"
	Vigor           = "Vigor"
	Intelligence    = "Intelligence"
	Education       = "Education"
	Training        = "Training"
	Instinct        = "Instinct"
	Social_Standing = "Social Standing"
	Charisma        = "Charisma"
	Caste           = "Caste"
	Territory       = "Territory"
	Psi             = "Psi"
	Sanity          = "Sanity"
	Wealth          = "Wealth"
	Luck            = "Luck"
	Morale          = "Morale"
	STR             = "STR"
	DEX             = "DEX"
	AGI             = "AGI"
	GRA             = "GRA"
	END             = "END"
	STA             = "STA"
	VIG             = "VIG"
	INT             = "INT"
	EDU             = "EDU"
	TRA             = "TRA"
	INS             = "INS"
	SOC             = "SOC"
	CHA             = "CHA"
	CAS             = "CAS"
	TER             = "TER"
	PSI             = "PSI"
	SAN             = "SAN"
	WLT             = "WLT"
	LCK             = "LCK"
	MOR             = "MOR"
	C1              = "C1"
	C2              = "C2"
	C3              = "C3"
	C4              = "C4"
	C5              = "C5"
	C6              = "C6"
	CS              = "CS"
	CP              = "CP"
	CM              = "CM"
	CL              = "CL"
	PHYSICAL        = "PHYSICAL"
	MENTAL          = "MENTAL"
	SOCIAL          = "SOCIAL"
	RACIAL          = "RACIAL"
	OBSCURE         = "OBSCURE"
)

func CoreChars() []string {
	return []string{
		Strength,
		Dexterity,
		Endurance,
		Intelligence,
		Education,
		Social_Standing,
	}
}

type Roller interface {
	Sroll(string) int
}
