package tradecode

import (
	"slices"

	profile "github.com/Galdoba/tabletoptools/pkg/mgt2/profile"
)

type TradeCode struct {
	code           string
	classification string
	description    string
	effect         string
	validCodes     map[string][]string
}

func Designate(pr profile.Profile) []string {
	codes := []string{}
	for _, tc := range []TradeCode{Ag, As, Ba, De, Fl, Ga, Hi, Ht, Ic, In, Lo, Lt, Na, Ni, Po, Ri, Va, Wa, Fr, Co, Te, Ho, Bo} {
		validated := 0
		for key, valid := range tc.validCodes {
			for _, target := range valid {
				if pr.Field(key) == target {
					validated++
				}
			}

		}
		if validated == len(tc.validCodes) {
			codes = append(codes, tc.code)
		}
	}
	slices.Sort(codes)
	return codes
}

var Ag = TradeCode{
	code:           "Ag",
	classification: "Agricultural",
	description:    "Dedicated to farming and food production. Often devided into vast semifeudal estates.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Atmo: {"4", "5", "6", "7", "8", "9"},
		profile.KEY_Hydr: {"4", "5", "6", "7", "8"},
		profile.KEY_Pops: {"5", "6", "7"},
	},
}

var As = TradeCode{
	code:           "As",
	classification: "Asteroid",
	description:    "Usualy mining colony, but can be orbital factory or commerce.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Size: {"0"},
		profile.KEY_Atmo: {"0"},
		profile.KEY_Hydr: {"0"},
	},
}

var Ba = TradeCode{
	code:           "Ba",
	classification: "Barren",
	description:    "Uncolonised and empty.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Pops: {"0"},
		profile.KEY_Govr: {"0"},
		profile.KEY_Laws: {"0"},
		profile.KEY_Port: {"X"},
	},
}

var De = TradeCode{
	code:           "De",
	classification: "Desert",
	description:    "Dry and barely habitable.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Atmo: {"2", "3", "4", "5", "6", "7", "8", "9"},
		profile.KEY_Hydr: {"0"},
	},
}

var Fl = TradeCode{
	code:           "Fl",
	classification: "Fluid Oceans",
	description:    "Surface liquid is something other than water. Incompatible with Earch-derived life.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Atmo: {"A", "B", "C", "D", "E", "F"},
		profile.KEY_Hydr: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "A"},
	},
}

var Ga = TradeCode{
	code:           "Ga",
	classification: "Garden",
	description:    "Earth-like world.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Size: {"6", "7", "8"},
		profile.KEY_Atmo: {"5", "6", "8"},
		profile.KEY_Hydr: {"5", "6", "7"},
	},
}

var Hi = TradeCode{
	code:           "Hi",
	classification: "High Population",
	description:    "Population in a billions.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Pops: {"9", "A", "B", "C", "D", "E", "F"},
	},
}

var Ht = TradeCode{
	code:           "Ht",
	classification: "High Tech",
	description:    "Among the most technologicaly advanced in Charted Space.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_TL: {"C", "D", "E", "F", "G", "H", "J", "K", "L", "M"},
	},
}

var Ic = TradeCode{
	code:           "Ic",
	classification: "Ice-Capped",
	description:    "Most of the surface liquid frozen on polar caps. Cold and dry.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Atmo: {"0", "1"},
		profile.KEY_Hydr: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "A"},
	},
}

var In = TradeCode{
	code:           "In",
	classification: "Industrial",
	description:    "Dominated by factories and cities.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Atmo: {"0", "1", "2", "4", "7", "9", "A", "B", "C"},
		profile.KEY_Pops: {"9", "A", "B", "C", "D", "E", "F"},
	},
}

var Lo = TradeCode{
	code:           "Lo",
	classification: "Low Population",
	description:    "Population is only a few thousand or less.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Pops: {"1", "2", "3"},
	},
}

var Lt = TradeCode{
	code:           "Lt",
	classification: "Low Tech",
	description:    "Pre-industrial and cannot produce advanced goods.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Pops: {"1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"},
		profile.KEY_TL:   {"0", "1", "2", "3", "4", "5"},
	},
}

var Na = TradeCode{
	code:           "Na",
	classification: "Non-Agricultural",
	description:    "Too dry or barren to support their population using conventional food production.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Atmo: {"0", "1", "2", "3"},
		profile.KEY_Hydr: {"0", "1", "2", "3"},
		profile.KEY_Pops: {"6", "7", "8", "9", "A", "B", "C", "D", "E", "F"},
	},
}

var Ni = TradeCode{
	code:           "Ni",
	classification: "Non-Industial",
	description:    "Too low in population to mantain an extensive industrial base.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Pops: {"4", "5", "6"},
	},
}

var Po = TradeCode{
	code:           "Po",
	classification: "Poor",
	description:    "Lacking resources or viable land to be other than marginal colony.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Atmo: {"2", "3", "4", "5"},
		profile.KEY_Hydr: {"0", "1", "2", "3"},
	},
}

var Ri = TradeCode{
	code:           "Ri",
	classification: "Rich",
	description:    "Blessed with stable goverment and viable biosphere, making it economic powerhouse.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Atmo: {"6", "8"},
		profile.KEY_Pops: {"6", "7", "8"},
		profile.KEY_Govr: {"4", "5", "6", "7", "8", "9"},
	},
}

var Va = TradeCode{
	code:           "Va",
	classification: "Vacuum",
	description:    "No atmosphere.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Atmo: {"0"},
	},
}

var Wa = TradeCode{
	code:           "Wa",
	classification: "Water World",
	description:    "Almost entirely ocean across their surface.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Atmo: {"3", "4", "5", "6", "7", "8", "9", "D", "E", "F"},
		profile.KEY_Hydr: {"A"},
	},
}

var Fr = TradeCode{
	code:           "Fr",
	classification: "Frozen",
	description:    "No liquid water, very dry atmosphere.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Temperature: {"0", "1", "2"},
	},
}

var Co = TradeCode{
	code:           "Co",
	classification: "Cold",
	description:    "Little liquid water, extencive ice caps, few clouds.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Temperature: {"3", "4"},
	},
}

var Te = TradeCode{
	code:           "Te",
	classification: "Temperate",
	description:    "Earth-like. Liquid and vaporized water are common, moderate ice caps.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Temperature: {"5", "6", "7", "8", "9"},
	},
}

var Ho = TradeCode{
	code:           "Ho",
	classification: "Hot",
	description:    "Small or no ice caps, little liquid water. Most water in the form of clouds",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Temperature: {"A", "B"},
	},
}

var Bo = TradeCode{
	code:           "Bo",
	classification: "Boiling",
	description:    "No ice caps, little liquid water.",
	effect:         "",
	validCodes: map[string][]string{
		profile.KEY_Temperature: {"C", "D", "E", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q"},
	},
}
