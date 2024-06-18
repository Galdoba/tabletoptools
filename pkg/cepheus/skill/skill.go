package skill

import "fmt"

const (
	ADMIN              = "Admin"
	AIRCRAFT           = "Aircraft"
	ANIMALS            = "Animals"
	ATHLETHICS         = "Athlethics"
	CAROUSING          = "Carousing"
	COMPUTER           = "Computer"
	DECEPTION          = "Deception"
	DEMOLITION         = "Demolition"
	DRIVING            = "Driving"
	ENGINEERING        = "Engineering"
	GRAV_VECHICLES     = "Grav Vehicles"
	GUN_COMBAT         = "Gun Combat"
	GUNNERY            = "Gunnery"
	HEAVY_WEAPONS      = "Heavy Weapons"
	INVESTIGATION      = "Investigation"
	JACK_OF_ALL_TRADES = "Jack Of All Trades"
	LEADERSHIP         = "Leadership"
	LIASON             = "Liason"
	MEDICINE           = "Medicine"
	MELEE_COMBAT       = "Melee Combat"
	PILOTING           = "Pilot"
	RECON              = "Recon"
	REPAIR             = "Repair"
	SCIENCE            = "Science"
	STEALTH            = "Stealth"
	STEWARD            = "Steward"
	STREETWISE         = "Streetwise"
	SURVIVAL           = "Survival"
	TACTICS            = "Tactics"
	WATERCRAFT         = "Watercraft"
	ZERO_G             = "Zero-G"
	maximumSkillValue  = 5
)

func List() []string {
	return []string{
		ADMIN,
		AIRCRAFT,
		ANIMALS,
		ATHLETHICS,
		CAROUSING,
		COMPUTER,
		DECEPTION,
		DEMOLITION,
		DRIVING,
		ENGINEERING,
		GRAV_VECHICLES,
		GUN_COMBAT,
		GUNNERY,
		HEAVY_WEAPONS,
		INVESTIGATION,
		JACK_OF_ALL_TRADES,
		LEADERSHIP,
		LIASON,
		MEDICINE,
		MELEE_COMBAT,
		PILOTING,
		RECON,
		REPAIR,
		SCIENCE,
		STEALTH,
		STEWARD,
		STREETWISE,
		SURVIVAL,
		TACTICS,
		WATERCRAFT,
		ZERO_G,
	}
}

func InList(skl string) bool {
	for _, s := range List() {
		if s == skl {
			return true
		}
	}
	return false
}

type SkillSet struct {
	Skill map[string]Skill
}

func NewSkillSet() *SkillSet {
	ss := SkillSet{}
	ss.Skill = map[string]Skill{}
	return &ss
}

func (ss *SkillSet) Inject(data map[string]int) error {
	for k, v := range data {
		if err := ss.CreateSkill(k, v); err != nil {
			//LOG HERE
			fmt.Println("injection failed:", err.Error())
			continue
		}
		fmt.Println("injected:", k, v)
	}
	return nil
}

// CRUD
func (ss *SkillSet) CreateSkill(skl string, val ...int) error {
	initVal := 0
	if !InList(skl) {
		return fmt.Errorf("unknown skill name provided: %v", skl)
	}
	for _, v := range val {
		if v < 0 {
			return fmt.Errorf("initial value must not be negative")
		}
		initVal = v
	}
	ss.Skill[skl] = newSkill(initVal)
	return nil
}

func (ss *SkillSet) ReadSkill(skl string) Skill {
	return ss.Skill[skl]
}

func (ss *SkillSet) Train(skl string) error {
	if ss.ReadSkill(skl) == nil {
		return fmt.Errorf("skill not found: %v", skl)
	}
	return ss.Skill[skl].Train()
}

func (ss *SkillSet) AddRelevance(skl, relevance string) error {
	oldRelevance := ss.Skill[skl].Relevance()
	if _, ok := oldRelevance[relevance]; ok {
		return fmt.Errorf("token already exists: %v", relevance)
	}
	oldRelevance[relevance] = true
	ss.Skill[skl].Relevance(oldRelevance)
	return nil
}

func (ss *SkillSet) RemoveRelevance(skl, relevance string) error {
	oldRelevance := ss.Skill[skl].Relevance()
	if _, ok := oldRelevance[relevance]; !ok {
		return fmt.Errorf("token not exists: %v", relevance)
	}
	delete(oldRelevance, relevance)
	ss.Skill[skl].Relevance(oldRelevance)
	return nil
}

func (ss *SkillSet) DeleteSkill(skl string) error {
	delete(ss.Skill, skl)
	return nil
}

type skill struct {
	actualVal     int
	relevantItems map[string]bool
}

func newSkill(initialValue int) *skill {
	sk := skill{}
	sk.actualVal = initialValue
	sk.relevantItems = make(map[string]bool)
	return &sk
}

type Skill interface {
	Value() int
	Train() error
	Relevance(...map[string]bool) map[string]bool
}

func (s *skill) Value() int {
	return s.actualVal
}

func (s *skill) Train() error {
	if s.actualVal >= maximumSkillValue {
		fmt.Errorf("maximum level reached")
	}
	s.actualVal++
	return nil
}

func (s *skill) Relevance(newRel ...map[string]bool) map[string]bool {
	if len(newRel) == 0 {
		return s.relevantItems
	}
	for _, nr := range newRel {
		s.relevantItems = nr
		break
	}
	return s.relevantItems
}

func TrainingCost(sk Skill) int {
	switch sk.Value() {
	case 0:
		return 10
	case 1, 2:
		return 20
	case 3:
		return 30
	case 4:
		return 40
	default:
		return 9999999999
	}
}
