package dynasty

type Archetype struct {
	Type          string
	Requisite     map[string]int
	BaseTraits    []string
	BaseAptitudes map[string]int
	BH_Bonuses    []string
	FG_Bonus      string
}
