package skill

const (
	Admin            = "Admin"
	Advocate         = "Advocate"
	Animals          = "Animals"
	Handling         = "Handling"
	Veterinary       = "Veterinary"
	Training         = "Training"
	Art              = "Art"
	Performer        = "Performer"
	Holography       = "Holography"
	Instrument       = "Instrument"
	VisualMedia      = "Visual Media"
	Write            = "Write"
	Astrogation      = "Astrogation"
	Athletics        = "Athletics"
	Dexterity        = "Dexterity"
	Endurance        = "Endurance"
	Strenght         = "Strenght"
	Broker           = "Broker"
	Carouse          = "Carouse"
	Deception        = "Deception"
	Diplomat         = "Diplomat"
	Drive            = "Drive"
	Hovercraft       = "Hovercraft"
	Mole             = "Mole"
	Track            = "Track"
	Walker           = "Walker"
	Wheel            = "Wheel"
	Electronics      = "Electronics"
	Comms            = "Comms"
	Computers        = "Computers"
	RemoteOps        = "Remote Ops"
	Sensors          = "Sensors"
	Engineer         = "Engineer"
	MDrive           = "M-drive"
	JDrive           = "J-drive"
	LifeSupport      = "Life Support"
	Power            = "Power"
	Explosives       = "Explosives"
	Flyer            = "Flyer"
	Airship          = "Airship"
	Grav             = "Grav"
	Ornithopter      = "Ornithopter"
	Rotor            = "Rotor"
	Wing             = "Wing"
	Gampler          = "Gampler"
	Gunner           = "Gunner"
	Turret           = "Turret"
	Ortilery         = "Ortilery"
	Screen           = "Screen"
	Capital          = "Capital"
	GunCombat        = "Gun Combat"
	Archaic          = "Archaic"
	Energy           = "Energy"
	Slug             = "Slug"
	HeavyWeapons     = "Heavy Weapons"
	Artilery         = "Artilery"
	Portable         = "Portable"
	Vechicle         = "Vechicle"
	Investigate      = "Investigate"
	JackOfAllTrades  = "Jack-of-All-Trades"
	Language         = "Language"
	Galinglic        = "Galinglic"
	Vilani           = "Vilani"
	Zdetl            = "Zdetl"
	Oynprith         = "Oynprith"
	Trokh            = "Trokh"
	Gvegh            = "Gvegh"
	Leadership       = "Leadership"
	Mechanic         = "Mechanic"
	Medic            = "Medic"
	Melee            = "Melee"
	Unarmed          = "Unarmed"
	Blade            = "Blade"
	Bludgeon         = "Bludgeon"
	Natural          = "Natural"
	Navigation       = "Navigation"
	Persuade         = "Persuade"
	Pilot            = "Pilot"
	SmallCraft       = "Small Craft"
	Spacecraft       = "Spacecraft"
	CapitalShip      = "Capital Ship"
	Profession       = "Profession"
	Belter           = "Belter"
	Biologicals      = "Biologicals"
	NewProfesion     = "NewProfesion"
	Recon            = "Recon"
	LifeSciences     = "Life Sciences"
	Biology          = "Biology"
	Genetics         = "Genetics"
	Psionicology     = "Psionicology"
	Xenology         = "Xenology"
	PhysicalSciences = "Physical Sciences"
	Chemistry        = "Chemistry"
	Physics          = "Physics"
	JumpspacePhysics = "Jumpspace Physics"
	RoboticSciences  = "RoboticSciences"
	Cybernetics      = "Cybernetics"
	Robotics         = "Robotics"
	SocialSciences   = "Social Sciences"
	Archaelogy       = "Archaelogy"
	Economics        = "Economics"
	History          = "History"
	Linuistics       = "Linuistics"
	Philosophy       = "Philosophy"
	Psychology       = "Psychology"
	Sophontology     = "Sophontology"
	SpaceSciences    = "Space Sciences"
	Astronomy        = "Astronomy"
	Cosmology        = "Cosmology"
	Planetology      = "Planetology"
	Seafarer         = "Seafarer"
	OceanShips       = "Ocean Ships"
	Personal         = "Personal"
	Sail             = "Sail"
	Submarine        = "Submarine"
	Stealth          = "Stealth"
	Steward          = "Steward"
	Streetwise       = "Streetwise"
	Survival         = "Survival"
	Tactics          = "Tactics"
	Military         = "Military"
	Naval            = "Naval"
	VaccSuit         = "Vacc Suit"
)

func specialitiesAndParent(name string) ([]string, string) {
	switch name {
	case Admin:
		return []string{}, ""
	case Advocate:
		return []string{}, ""
	case Animals:
		return []string{Handling, Veterinary, Training}, ""
	case Handling:
		return []string{}, Animals
	case Veterinary:
		return []string{}, Animals
	case Training:
		return []string{}, Animals
	case Art:
		return []string{Performer, Holography, Instrument, VisualMedia, Write}, ""
	case Performer:
		return []string{}, Art
	case Holography:
		return []string{}, Art
	case Instrument:
		return []string{}, Art
	case VisualMedia:
		return []string{}, Art
	case Write:
		return []string{}, Art
	case Astrogation:
		return []string{}, ""
	case Athletics:
		return []string{Dexterity, Endurance, Strenght}, ""
	case Dexterity:
		return []string{}, Athletics
	case Endurance:
		return []string{}, Athletics
	case Strenght:
		return []string{}, Athletics
	case Broker:
		return []string{}, ""
	case Carouse:
		return []string{}, ""
	case Deception:
		return []string{}, ""
	case Diplomat:
		return []string{}, ""
	case Drive:
		return []string{Hovercraft, Mole, Track, Walker, Wheel}, ""
	case Hovercraft:
		return []string{}, Drive
	case Mole:
		return []string{}, Drive
	case Track:
		return []string{}, Drive
	case Walker:
		return []string{}, Drive
	case Wheel:
		return []string{}, Drive
	case Electronics:
		return []string{Comms, Computers, RemoteOps, Sensors}, ""
	case Comms:
		return []string{}, Electronics
	case Computers:
		return []string{}, Electronics
	case RemoteOps:
		return []string{}, Electronics
	case Sensors:
		return []string{}, Electronics
	case Engineer:
		return []string{MDrive, JDrive, LifeSupport, Power}, ""
	case MDrive:
		return []string{}, Engineer
	case JDrive:
		return []string{}, Engineer
	case LifeSupport:
		return []string{}, Engineer
	case Power:
		return []string{}, Engineer
	case Explosives:
		return []string{}, ""
	case Flyer:
		return []string{Airship, Grav, Ornithopter, Rotor, Wing}, ""
	case Airship:
		return []string{}, Flyer
	case Grav:
		return []string{}, Flyer
	case Ornithopter:
		return []string{}, Flyer
	case Rotor:
		return []string{}, Flyer
	case Wing:
		return []string{}, Flyer
	case Gampler:
		return []string{}, ""
	case Gunner:
		return []string{Turret, Ortilery, Screen, Capital}, ""
	case Turret:
		return []string{}, Gunner
	case Ortilery:
		return []string{}, Gunner
	case Screen:
		return []string{}, Gunner
	case Capital:
		return []string{}, Gunner
	case GunCombat:
		return []string{Archaic, Energy, Slug}, ""
	case Archaic:
		return []string{}, GunCombat
	case Energy:
		return []string{}, GunCombat
	case Slug:
		return []string{}, GunCombat
	case HeavyWeapons:
		return []string{Artilery, Portable, Vechicle}, ""
	case Artilery:
		return []string{}, HeavyWeapons
	case Portable:
		return []string{}, HeavyWeapons
	case Vechicle:
		return []string{}, HeavyWeapons
	case Investigate:
		return []string{}, ""
	case JackOfAllTrades:
		return []string{}, ""
	case Language:
		return []string{Galinglic, Vilani, Zdetl, Oynprith, Trokh, Gvegh}, ""
	case Galinglic:
		return []string{}, Language
	case Vilani:
		return []string{}, Language
	case Zdetl:
		return []string{}, Language
	case Oynprith:
		return []string{}, Language
	case Trokh:
		return []string{}, Language
	case Gvegh:
		return []string{}, Language
	case Leadership:
		return []string{}, ""
	case Mechanic:
		return []string{}, ""
	case Medic:
		return []string{}, ""
	case Melee:
		return []string{Unarmed, Blade, Bludgeon, Natural}, ""
	case Unarmed:
		return []string{}, Melee
	case Blade:
		return []string{}, Melee
	case Bludgeon:
		return []string{}, Melee
	case Natural:
		return []string{}, Melee
	case Navigation:
		return []string{}, ""
	case Persuade:
		return []string{}, ""
	case Pilot:
		return []string{SmallCraft, Spacecraft, CapitalShip}, ""
	case SmallCraft:
		return []string{}, Pilot
	case Spacecraft:
		return []string{}, Pilot
	case CapitalShip:
		return []string{}, Pilot
	case Profession:
		return []string{Belter, Biologicals, NewProfesion}, ""
	case Belter:
		return []string{}, Profession
	case Biologicals:
		return []string{}, Profession
	case NewProfesion:
		return []string{}, Profession
	case Recon:
		return []string{}, ""
	case LifeSciences:
		return []string{Biology, Genetics, Psionicology, Xenology}, ""
	case Biology:
		return []string{}, LifeSciences
	case Genetics:
		return []string{}, LifeSciences
	case Psionicology:
		return []string{}, LifeSciences
	case Xenology:
		return []string{}, LifeSciences
	case PhysicalSciences:
		return []string{Chemistry, Physics, JumpspacePhysics}, ""
	case Chemistry:
		return []string{}, PhysicalSciences
	case Physics:
		return []string{}, PhysicalSciences
	case JumpspacePhysics:
		return []string{}, PhysicalSciences
	case RoboticSciences:
		return []string{Cybernetics, Robotics}, ""
	case Cybernetics:
		return []string{}, RoboticSciences
	case Robotics:
		return []string{}, RoboticSciences
	case SocialSciences:
		return []string{Archaelogy, Economics, History, Linuistics, Philosophy, Psychology, Sophontology}, ""
	case Archaelogy:
		return []string{}, SocialSciences
	case Economics:
		return []string{}, SocialSciences
	case History:
		return []string{}, SocialSciences
	case Linuistics:
		return []string{}, SocialSciences
	case Philosophy:
		return []string{}, SocialSciences
	case Psychology:
		return []string{}, SocialSciences
	case Sophontology:
		return []string{}, SocialSciences
	case SpaceSciences:
		return []string{Astronomy, Cosmology, Planetology}, ""
	case Astronomy:
		return []string{}, SpaceSciences
	case Cosmology:
		return []string{}, SpaceSciences
	case Planetology:
		return []string{}, SpaceSciences
	case Seafarer:
		return []string{OceanShips, Personal, Sail, Submarine}, ""
	case OceanShips:
		return []string{}, Seafarer
	case Personal:
		return []string{}, Seafarer
	case Sail:
		return []string{}, Seafarer
	case Submarine:
		return []string{}, Seafarer
	case Stealth:
		return []string{}, ""
	case Steward:
		return []string{}, ""
	case Streetwise:
		return []string{}, ""
	case Survival:
		return []string{}, ""
	case Tactics:
		return []string{Military, Naval}, ""
	case Military:
		return []string{}, Tactics
	case Naval:
		return []string{}, Tactics
	case VaccSuit:
		return []string{}, ""
	default:
		return []string{""}, "bad name"
	}
}
