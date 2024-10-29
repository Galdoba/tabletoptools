package dynasty

const (
	ALL_STATs = "all_stats"
	ALL_APTs  = "all_apts"

	APT_Acquisition          = "Acquisition"
	APT_Bureaucracy          = "Bureaucracy"
	APT_Conquest             = "Conquest"
	APT_Economics            = "Economics"
	APT_Entertain            = "Entertain"
	APT_Expression           = "Expression"
	APT_Hostility            = "Hostility"
	APT_Illicit              = "Illicit"
	APT_Intel                = "Intel"
	APT_Maintenance          = "Maintenance"
	APT_Politics             = "Politics"
	APT_Posturing            = "Posturing"
	APT_Propaganda           = "Propaganda"
	APT_PublicRelations      = "Public Relations"
	APT_Recruitment          = "Recruitment"
	APT_Research             = "Research"
	APT_Sabotage             = "Sabotage"
	APT_Security             = "Security"
	APT_Tactical             = "Tactical"
	APT_Tutelage             = "Tutelage"
	BASE_ColonySettlement    = "Colony Settlement"
	BASE_ConflictZone        = "Conflict Zone"
	BASE_Megalopolis         = "Megalopolis"
	BASE_MilitaryCompound    = "Military Compound"
	BASE_NobleEstate         = "Noble Estate"
	BASE_StarshipFlotilla    = "Starship/Flotilla"
	BASE_TempleHolyLand      = "Temple/Holy Land"
	BASE_UnchartedWilderness = "Uncharted Wilderness"
	BASE_UnderworldSlum      = "Underworld Slum"
	BASE_UrbanOffices        = "Urban Offices"
	STAT_Cleverness          = "Cleverness"
	STAT_Greed               = "Greed"
	STAT_Loyalty             = "Loyalty"
	STAT_Militarism          = "Militarism"
	STAT_Popularity          = "Popularity"
	STAT_Scheming            = "Scheming"
	STAT_Tenacity            = "Tenacity"
	STAT_Tradition           = "Tradition"
	TRAIT_Culture            = "Culture"
	TRAIT_FiscalDefence      = "FiscalDefence"
	TRAIT_Fleet              = "Fleet"
	TRAIT_Technology         = "Technology"
	TRAIT_TerritorialDefence = "TerritorialDefence"
	VALUE_Morale             = "Morale"
	VALUE_Populance          = "Populance"
	VALUE_Wealth             = "Wealth"

	CONDITION_ALWAYS             = "Always"
	CONDITION_BetweenGenerations = "Between Generations"
	CONDITION_Value_Loss_Wealth  = "On Wealth Loss"

	method_DiceRolling   = 1
	method_PointBased    = 2
	method_CharaterBased = 3
)

func CombineKeys(ks ...string) []string {
	keysCombined := []string{}
	for _, k := range ks {
		switch k {
		case ALL_APTs:
			keysCombined = append(keysCombined, AptitudeKeys()...)
		case ALL_STATs:
			keysCombined = append(keysCombined, CharacteristicKeys()...)
		default:
			keysCombined = append(keysCombined, k)
		}
	}
	return keysCombined
}

func AptitudeKeys() []string {
	return []string{
		APT_Acquisition,
		APT_Bureaucracy,
		APT_Conquest,
		APT_Economics,
		APT_Entertain,
		APT_Expression,
		APT_Hostility,
		APT_Illicit,
		APT_Intel,
		APT_Maintenance,
		APT_Politics,
		APT_Posturing,
		APT_Propaganda,
		APT_PublicRelations,
		APT_Recruitment,
		APT_Research,
		APT_Sabotage,
		APT_Security,
		APT_Tactical,
		APT_Tutelage,
	}
}

func CharacteristicKeys() []string {
	return []string{
		STAT_Cleverness,
		STAT_Greed,
		STAT_Loyalty,
		STAT_Militarism,
		STAT_Popularity,
		STAT_Scheming,
		STAT_Tenacity,
		STAT_Tradition,
	}
}
