package entitytest

import (
	"fmt"
	"strings"

	"github.com/Galdoba/tabletoptools/pkg/entity"
)

type Weapon struct {
	W_Type         string
	W_ID           int64
	Name           string
	TL             int
	Range          string
	Damage         string
	KG             float64
	CostCr         int
	Magazine       int
	MagazineCostCr int
	Traits         []string
	Skill          string
}

const (

	//static_Type              = "Type"
	static_Name              = "Name"
	static_Range             = "Range"
	static_Damage            = "Damage"
	static_Traits            = "Traits"
	static_Skill             = "Skill"
	countable_TL             = "TL"
	countable_KG             = "KG"
	countable_CostCr         = "CostCr"
	countable_Magazine       = "Magazine"
	countable_MagazineCostCr = "MagazineCostCr"
)

func NewPistol() *Weapon {
	w := Weapon{}
	w.W_Type = "Pistol"
	w.Name = "Gauss Pistol"
	w.TL = 13
	w.Range = "20m"
	w.Damage = "3D"
	w.KG = 0.75
	w.CostCr = 500
	w.Magazine = 40
	w.MagazineCostCr = 20
	w.Traits = []string{
		"AP 3",
		"Auto 2",
	}
	return &w
}

func NewBlade() *Weapon {
	w := Weapon{}
	w.W_Type = "One-handed"
	w.Name = "Blade"
	w.TL = 2
	w.Range = "Melee"
	w.Damage = "2D"
	w.KG = 1
	w.CostCr = 100
	w.Magazine = 0
	w.MagazineCostCr = 0
	w.Traits = []string{}
	return &w
}

/*Entity implementation*/
func (w *Weapon) StaticData() map[string]string {
	static := make(map[string]string)
	static[static_Name] = w.Name
	static[static_Range] = w.Range
	static[static_Damage] = w.Damage
	static[static_Traits] = strings.Join(w.Traits, " && ")
	static[static_Skill] = w.Skill
	return static
}

func (w *Weapon) CountableData() map[string]float64 {
	countable := make(map[string]float64)
	countable[countable_TL] = float64(w.TL)
	countable[countable_KG] = w.KG
	countable[countable_CostCr] = float64(w.CostCr)
	countable[countable_Magazine] = float64(w.Magazine)
	countable[countable_MagazineCostCr] = float64(w.MagazineCostCr)
	return countable
}

func (w *Weapon) Type() string {
	return w.W_Type
}

func (w *Weapon) ID() int64 {
	return w.W_ID
}

func (w *Weapon) Classification() []string {
	return []string{"weapon", "personal", w.W_Type}
}

func (w *Weapon) Events() []string {
	return nil
}

func (w *Weapon) Validation() error {
	return nil
}

func Load(id string, classification []string) (*Weapon, error) {
	e, err := entity.Load(id, classification)
	if err != nil {
		return nil, fmt.Errorf("loading failed: %v")
	}
	w, err := fromEntity(e)
	if err != nil {
		return nil, fmt.Errorf("conversion failed: %v")
	}
	return w, nil
}

func fromEntity(e *entity.EntityFile) (*Weapon, error) {
	return &Weapon{}, nil
}
