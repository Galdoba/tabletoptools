package dynasty

type PowerBase struct {
	Base        string `json:"Type"`
	Description string `json:"Description"`
	modifiers   map[string]int
}

var ColonySettlement = PowerBase{
	Base:        BASE_ColonySettlement,
	Description: "The Dynasty rises in a frontier town, small colony or rural community. Easily becoming the ‘big fish in a little pond’, the Dynasty can count on a moderate degree of control over the local population – especially if it has something beneficial to offer them.",
	modifiers: map[string]int{
		TRAIT_Culture:            1,
		TRAIT_TerritorialDefence: -1,
		APT_Expression:           1,
		APT_Recruitment:          1,
		APT_Maintenance:          1,
		APT_Propaganda:           1,
		APT_Tutelage:             1,
	},
}

var ConflictZone = PowerBase{
	Base:        BASE_ConflictZone,
	Description: "The Dynasty’s first generation appears within the confines of an area ravaged by war, battle or other longstanding conflict. This is the perfect breeding ground for highly faithful or militaryminded Dynasties; no surprise considering the odds they must have overcome to survive and prosper.",
	modifiers: map[string]int{
		TRAIT_TerritorialDefence: 2,
		TRAIT_FiscalDefence:      -1,
		TRAIT_Fleet:              -1,
		APT_Hostility:            2,
		APT_Posturing:            1,
		APT_Security:             1,
		APT_Tactical:             1,
	},
}

var Megalopolis = PowerBase{
	Base:        BASE_Megalopolis,
	Description: "The Dynasty grew among the enormous, continent or even planetary-scale cityscapes. They present massive populations to draw from, higher average Technology Levels and access o excellent communication, economic trends and a variety f other ‘modern’ benefits to growing amidst a sea of ultraurbanites.",
	modifiers: map[string]int{
		TRAIT_FiscalDefence: 1,
		TRAIT_Technology:    1,
		TRAIT_Culture:       -2,
		APT_Bureaucracy:     2,
		APT_Economics:       1,
		APT_PublicRelations: 1,
		APT_Research:        1,
	},
}

var MilitaryCompound = PowerBase{
	Base:        BASE_MilitaryCompound,
	Description: "The Dynasty formed within the boundaries of a militarily controlled territory – the perfect grounds to create a draconian tradition of growth, mastery and conflict-management. Peaceful dynasties are not likely to appear here but more martially inclined ones will thrive.",
	modifiers: map[string]int{
		TRAIT_TerritorialDefence: 2,
		TRAIT_Fleet:              1,
		TRAIT_FiscalDefence:      -2,
		APT_Conquest:             2,
		APT_Tactical:             2,
		APT_Politics:             1,
		APT_Posturing:            1,
		APT_Security:             1,
	},
}

var NobleEstate = PowerBase{
	Base:        BASE_NobleEstate,
	Description: "The Dynasty had a wealthy, noble family’s territories and resources to call upon. Money, status and power was always nearby; they wanted for nothing and used this access to grow and prosper with relative protection and shelter from outside threats.",
	modifiers: map[string]int{
		TRAIT_Culture:            1,
		TRAIT_FiscalDefence:      1,
		TRAIT_TerritorialDefence: -2,
		TRAIT_Fleet:              -1,
		APT_Bureaucracy:          2,
		APT_Politics:             2,
		APT_Expression:           1,
		APT_Posturing:            1,
		APT_Security:             1,
	},
}

var StarshipFlotilla = PowerBase{
	Base:        BASE_StarshipFlotilla,
	Description: "The Dynasty evolved from the enclosed population of a large starship, orbital flotilla or space station. This resulted in a powerful respect for the open vacuum of space as well as all of the things required to call it home. Life in space has altered their perception of outsiders but it has given them a superior connection to the spacefaring culture at large.",
	modifiers: map[string]int{
		TRAIT_Fleet:              2,
		TRAIT_Technology:         1,
		TRAIT_TerritorialDefence: -2,
		APT_Intel:                2,
		APT_Conquest:             1,
		APT_Economics:            1,
		APT_Maintenance:          1,
		APT_Posturing:            1,
		APT_Research:             1,
		APT_Tactical:             1,
	},
}

var TempleHolyLand = PowerBase{
	Base:        BASE_TempleHolyLand,
	Description: "The Dynasty grew in the protection of a religion’s heartland; the core of the faith and the centre of its influence. This strengthened its resolve and offered the Dynasty strength that arms and assets cannot equal – strength of belief, even if it may be misguided.",
	modifiers: map[string]int{
		TRAIT_Culture:       2,
		TRAIT_Technology:    -2,
		APT_Expression:      2,
		APT_Recruitment:     2,
		APT_Maintenance:     1,
		APT_Propaganda:      1,
		APT_PublicRelations: 1,
		APT_Tutelage:        1,
	},
}

var UnchartedWilderness = PowerBase{
	Base:        BASE_UnchartedWilderness,
	Description: "The Dynasty defeated the odds, its first generation springing up from the wild and untamed frontier where they were allowed to be as secluded and secretive as they wanted – lending a hand toward unseen goals or practices well enough.",
	modifiers: map[string]int{
		TRAIT_TerritorialDefence: 1,
		TRAIT_Technology:         -1,
		APT_Security:             3,
		APT_Entertain:            1,
		APT_Illicit:              1,
	},
}

var UnderworldSlum = PowerBase{
	Base:        BASE_UnderworldSlum,
	Description: "The Dynasty rose up from the depths of the criminal underworld. Surrounded by illegal activity, amoral choices and ample resources that were surely acquired in a morally grey fashion, the Dynasty has a dark foundation with which to grow a potentially even darker future.",
	modifiers: map[string]int{
		TRAIT_FiscalDefence:      1,
		TRAIT_TerritorialDefence: 1,
		TRAIT_Culture:            -2,
		APT_Illicit:              2,
		APT_Sabotage:             2,
		APT_Entertain:            1,
		APT_Intel:                1,
		APT_Posturing:            1,
		APT_Security:             1,
	},
}

var UrbanOffices = PowerBase{
	Base:        BASE_UrbanOffices,
	Description: "The Dynasty was drawn and drafted within the urban sprawl of a downtown business centre or its equivalent. Tall buildings filled with possible allies and enemies of a white-collar nature surround them, making their growth in the commercial world second nature as they gathered power and resources.",
	modifiers: map[string]int{
		TRAIT_Culture:       1,
		TRAIT_FiscalDefence: 1,
		TRAIT_Fleet:         -1,
		APT_Acquisition:     2,
		APT_Economics:       2,
		APT_Bureaucracy:     1,
		APT_Intel:           1,
		APT_PublicRelations: 1,
		APT_Tutelage:        1,
	},
}
