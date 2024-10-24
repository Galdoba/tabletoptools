package dynasty

type Bonus struct {
	Name             string
	StatCost         map[string]int
	PointCost        int
	Description      string
	CheckBonuses     map[string]int
	DefenceBonuses   map[string]int
	CheckEffectLimit map[string]int
	RerollAtempts    int
	RerollKeys       []string
	Condition        string
	ActionSTG        *action_SpendToGain
}

var Commercial_Psions = Bonus{
	Name: "Commercial Psions",
	StatCost: map[string]int{
		STAT_Popularity: -1,
	},
	PointCost:   -10,
	Description: "The Dynasty has a tradition of employing skilled telepathic psions for use in their business dealings; eliminating a lot of dishonesty while giving a special edge of their own. The Conglomerate can, between generations, spend one point of its Populace Value (signifying focussed breeding) to gain 1d6 – 3 (minimum of 0) Wealth Value.",
	Condition:   CONDITION_BetweenGenerations,
	ActionSTG:   &action_Commercial_Psions,
}

var Endless_Funds = Bonus{
	Name: "Endless Funds",
	StatCost: map[string]int{
		TRAIT_FiscalDefence: -2,
	},
	PointCost:     -15,
	Description:   "The Dynasty has countless financial connections and investors to call upon, spread across the galaxy. Even when the coffers seem to be thinning, a new influx of capital presents itself. The Conglomerate can re-roll any failed Aptitude or Characteristic check that would result in a loss to the Wealth Value.",
	RerollAtempts: 1,
	RerollKeys:    CombineKeys(ALL_APTs, ALL_STATs),
	Condition:     CONDITION_Value_Loss_Wealth,
}

var Govermental_Backing = Bonus{
	Name: "Govermental Backing",
	StatCost: map[string]int{
		STAT_Tradition: -1,
	},
	PointCost:   -5,
	Description: "Dynasty has acquired the ironclad backing of several local governments or perhaps one superpower, giving them an edge with certain issues. The Conglomerate can always count its Fiscal Defence and Territorial Defence as one point higher for testing purposes (but not the actual Value itself).",
	CheckBonuses: map[string]int{
		TRAIT_FiscalDefence:      1,
		TRAIT_TerritorialDefence: 1,
	},
	DefenceBonuses: map[string]int{
		TRAIT_FiscalDefence:      1,
		TRAIT_TerritorialDefence: 1,
	},
	Condition: CONDITION_ALWAYS,
}

/////
/////
/////
/////
/////
/////

type action_SpendToGain struct {
	spendKey    string
	spendValue  string
	gainKey     string
	gainValue   string
	gainMinimum int
}

var action_Commercial_Psions = action_SpendToGain{
	spendKey:    VALUE_Populance,
	spendValue:  "1d1",
	gainKey:     VALUE_Wealth,
	gainValue:   "1d6-3",
	gainMinimum: 0,
}
