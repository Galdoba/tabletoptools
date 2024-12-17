package skill

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	generationLimit       = 4
	SK_Admin              = "Admin"
	SK_Advocate           = "Advocate"
	SK_Animals            = "Animals"
	SK_Art                = "Art"
	SK_Astrogation        = "Astrogation"
	SK_Athlethics         = "Athlethics"
	SK_Broker             = "Broker"
	SK_Carouse            = "Carouse"
	SK_Deceiption         = "Deceiption"
	SK_Diplomat           = "Diplomat"
	SK_Drive              = "Drive"
	SK_Electronics        = "Electronics"
	SK_Engineer           = "Engineer"
	SK_Explosives         = "Explosives"
	SK_Flyer              = "Flyer"
	SK_Gambler            = "Gambler"
	SK_Gunner             = "Gunner"
	SK_Gun_Combat         = "Gun_Combat"
	SK_Heavy_Weapons      = "Heavy_Weapons"
	SK_Investigate        = "Investigate"
	SK_Jack_of_All_Trades = "Jack-of-All-Trades"
	SK_Language           = "Language"
	SK_Leadership         = "Leadership"
	SK_Science_Life       = "Life Sciences"
	SK_Mechanic           = "Mechanic"
	SK_Medic              = "Medic"
	SK_Melee              = "Melee"
	SK_Navigation         = "Navigation"
	SK_Persuade           = "Persuade"
	SK_Pilot              = "Pilot"
	SK_Science_Physical   = "Physical Sciences"
	SK_Proffesion         = "Proffesion"
	SK_Recon              = "Recon"
	SK_Science_Robotics   = "Robotics Sciences"
	SK_Seafarer           = "Seafarer"
	SK_Science_Social     = "Social Sciences"
	SK_Science_Space      = "Space Sciences"
	SK_Stealth            = "Stealth"
	SK_Steward            = "Steward"
	SK_Streetwise         = "Streetwise"
	SK_Survival           = "Survival"
	SK_Tactics            = "Tactics"
	SK_Vacc_Suit          = "Vacc_Suit"

	SP_Handling          = "Handling"
	SP_Veterinary        = "Veterinary"
	SP_Training          = "Training"
	SP_Performing        = "Performing"
	SP_Creative          = "Creative"
	SP_Presentation      = "Presentation"
	SP_Strenght          = "Strenght"
	SP_Dexterity         = "Dexterity"
	SP_Endurance         = "Endurance"
	SP_Hovercraft        = "Hovercraft"
	SP_Mole              = "Mole"
	SP_Track             = "Track"
	SP_Walker            = "Walker"
	SP_Wheel             = "Wheel"
	SP_Comms             = "Comms"
	SP_Computers         = "Computers"
	SP_Remote_Ops        = "Remote_Ops"
	SP_Sensors           = "Sensors"
	SP_M_drive           = "M-drive"
	SP_J_drive           = "J-drive"
	SP_Life_Support      = "Life Support"
	SP_Power             = "Power"
	SP_Airship           = "Airship"
	SP_Grav              = "Grav"
	SP_Ornithopter       = "Ornithopter"
	SP_Rotor             = "Rotor"
	SP_Wing              = "Wing"
	SP_Turret            = "Turret"
	SP_Ortilery          = "Ortilery"
	SP_Screen            = "Screen"
	SP_Capital           = "Capital"
	SP_Archaic           = "Archaic"
	SP_Energy            = "Energy"
	SP_Slug              = "Slug"
	SP_Artilery          = "Artilery"
	SP_Portable          = "Portable"
	SP_Vechicle          = "Vechicle"
	SP_Lang1             = "Lang1"
	SP_Lang2             = "Lang2"
	SP_Lang3             = "Lang3"
	SP_Unarmed           = "Unarmed"
	SP_Grapple           = "Grapple"
	SP_Striking          = "Striking"
	SP_Fencing           = "Fencing"
	SP_Ocean_Ships       = "Ocean Ships"
	SP_Personal          = "Personal"
	SP_Sail              = "Sail"
	SP_Submarine         = "Submarine"
	SP_Military          = "Military"
	SP_Naval             = "Naval"
	SP_Biology           = "Biology"
	SP_Genetics          = "Genetics"
	SP_Psionicology      = "Psionicology"
	SP_Xenology          = "Xenology"
	SP_Chemistry         = "Chemistry"
	SP_Physics           = "Physics"
	SP_Jumpspace_Physics = "Jumpspace Physics"
	SP_Cybernetcs        = "Cybernetcs"
	SP_Robotics          = "Robotics"
	SP_Archaeology       = "Archaeology"
	SP_Economics         = "Economics"
	SP_History           = "History"
	SP_Linguistics       = "Linguistics"
	SP_Philosophy        = "Philosophy"
	SP_Psyhology         = "Psyhology"
	SP_Sophontology      = "Sophontology"
	SP_Astronomy         = "Astronomy"
	SP_Cosmology         = "Cosmology"
	SP_Planetology       = "Planetology"
)

type SkillSet struct {
	Skill map[string]int
}

func NewSet() *SkillSet {
	ss := SkillSet{}
	ss.Skill = make(map[string]int)
	return &ss
}

type Skill struct {
	Name       string
	Speciality string
	Level      int
}

func (ss *SkillSet) Value(key string) int {
	return -3
}

func (sk *Skill) Key() string {
	if sk.Speciality == "" {
		return sk.Name
	}
	return fmt.Sprintf("%v (%v)", sk.Name, sk.Speciality)
}

func New(key string) *Skill {
	sk := Skill{}
	skill, spec, lv := Unstring(key)
	sk.Name = skill
	sk.Speciality = spec
	sk.Level = lv
	return &sk
}

func Unstring(str string) (string, string, int) {
	reLevel := regexp.MustCompile(`( [0123456789])$?`)
	lev := reLevel.FindString(str)
	reSpec := regexp.MustCompile(`\((.*?)\)`)
	spec := reSpec.FindString(str)
	spec = strings.TrimPrefix(spec, "(")
	spec = strings.TrimSuffix(spec, ")")
	skName := strings.TrimSuffix(str, lev)
	skName = strings.TrimSuffix(skName, fmt.Sprintf(" (%v)", spec))
	val, err := strconv.Atoi(strings.TrimPrefix(lev, " "))
	if err != nil {
		return skName, spec, -1
	}
	return skName, spec, val
}

func (sk *Skill) increase(val ...int) {
	switch len(val) {
	case 0:
		sk.Level = sk.Level + 1
	default:
		if sk.Level < val[0] {
			sk.Level = val[0]
		}
	}
}
