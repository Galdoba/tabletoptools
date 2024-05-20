package characteristic

type Value struct {
	Current       int //текущее значение
	Max           int //максимальное значение для сета
	Mod           int //модификатор: метод => val.Mod(ruleset string) int (ruleset = mgt2/HOSTILE/BARBARIC и т.д.)
	InheritedGene int //1D/2D.../6D
}
