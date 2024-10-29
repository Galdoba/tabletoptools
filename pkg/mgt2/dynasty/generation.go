package dynasty

import (
	"encoding/json"
	"fmt"

	"github.com/Galdoba/tabletoptools/pkg/dice"
)

func New() *Dynasty {
	d := Dynasty{}
	return &d
}

type dynastyGenerator struct {
	dice   *dice.Dicepool
	method int
	auto   bool
}

func NewDynastyGenerator(dice *dice.Dicepool, method int, auto_decidion bool) *dynastyGenerator {
	return &dynastyGenerator{dice, method, auto_decidion}
}

func (gen *dynastyGenerator) Generate() (*Dynasty, error) {
	d := New()
	d.Stats = gen.generateStats()
	d.PowerBase = gen.choosePowerBase()
	return d, nil
}

func (gen *dynastyGenerator) generateStats() *Stats {
	switch gen.method {
	default:
		panic(fmt.Sprintf("characteristic generation: method %v not implemented", gen.method))
	case method_DiceRolling:
		return randomStats(gen.dice)
	}
}

func randomStats(dice *dice.Dicepool) *Stats {
	st := Stats{}
	st.Cleverness = dice.Roll_2D()
	st.Greed = dice.Roll_2D()
	st.Loyalty = dice.Roll_2D()
	st.Militarism = dice.Roll_2D()
	st.Popularity = dice.Roll_2D()
	st.Scheming = dice.Roll_2D()
	st.Tenacity = dice.Roll_2D()
	st.Tradition = dice.Roll_2D()
	return &st
}

func (gen *dynastyGenerator) choosePowerBase() PowerBase {
	switch gen.auto {
	case true:
		return rollPowerBase(gen.dice)
	case false:
		panic("power base: manual selection failed")
	}
	return PowerBase{}
}

func rollPowerBase(dice *dice.Dicepool) PowerBase {
	bases := []PowerBase{
		ColonySettlement,
		ConflictZone,
		Megalopolis,
		MilitaryCompound,
		NobleEstate,
		StarshipFlotilla,
		TempleHolyLand,
		UnchartedWilderness,
		UnderworldSlum,
		UrbanOffices,
	}
	code := fmt.Sprintf("1d%v-1", len(bases))
	return bases[dice.Sroll(code)]
}

////////////////////
//
//
//
///////////////////

func (d *Dynasty) String() string {
	bt, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(bt)
}
