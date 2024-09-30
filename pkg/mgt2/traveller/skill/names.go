package skill

import . "github.com/Galdoba/tabletoptools/pkg/mgt2/key"

func specialitiesAndParent(name string) ([]string, string) {
	switch name {
	case SKL_Admin:
		return []string{}, ""
	case SKL_Advocate:
		return []string{}, ""
	case SKL_Animals:
		return []string{SKL_Handling, SKL_Veterinary, SKL_Training}, ""
	case SKL_Handling:
		return []string{}, SKL_Animals
	case SKL_Veterinary:
		return []string{}, SKL_Animals
	case SKL_Training:
		return []string{}, SKL_Animals
	case SKL_Art:
		return []string{SKL_Performer, SKL_Holography, SKL_Instrument, SKL_VisualMedia, SKL_Write}, ""
	case SKL_Performer:
		return []string{}, SKL_Art
	case SKL_Holography:
		return []string{}, SKL_Art
	case SKL_Instrument:
		return []string{}, SKL_Art
	case SKL_VisualMedia:
		return []string{}, SKL_Art
	case SKL_Write:
		return []string{}, SKL_Art
	case SKL_Astrogation:
		return []string{}, ""
	case SKL_Athletics:
		return []string{SKL_Dexterity, SKL_Endurance, SKL_Strenght}, ""
	case SKL_Dexterity:
		return []string{}, SKL_Athletics
	case SKL_Endurance:
		return []string{}, SKL_Athletics
	case SKL_Strenght:
		return []string{}, SKL_Athletics
	case SKL_Broker:
		return []string{}, ""
	case SKL_Carouse:
		return []string{}, ""
	case SKL_Deception:
		return []string{}, ""
	case SKL_Diplomat:
		return []string{}, ""
	case SKL_Drive:
		return []string{SKL_Hovercraft, SKL_Mole, SKL_Track, SKL_Walker, SKL_Wheel}, ""
	case SKL_Hovercraft:
		return []string{}, SKL_Drive
	case SKL_Mole:
		return []string{}, SKL_Drive
	case SKL_Track:
		return []string{}, SKL_Drive
	case SKL_Walker:
		return []string{}, SKL_Drive
	case SKL_Wheel:
		return []string{}, SKL_Drive
	case SKL_Electronics:
		return []string{SKL_Comms, SKL_Computers, SKL_RemoteOps, SKL_Sensors}, ""
	case SKL_Comms:
		return []string{}, SKL_Electronics
	case SKL_Computers:
		return []string{}, SKL_Electronics
	case SKL_RemoteOps:
		return []string{}, SKL_Electronics
	case SKL_Sensors:
		return []string{}, SKL_Electronics
	case SKL_Engineer:
		return []string{SKL_MDrive, SKL_JDrive, SKL_LifeSupport, SKL_Power}, ""
	case SKL_MDrive:
		return []string{}, SKL_Engineer
	case SKL_JDrive:
		return []string{}, SKL_Engineer
	case SKL_LifeSupport:
		return []string{}, SKL_Engineer
	case SKL_Power:
		return []string{}, SKL_Engineer
	case SKL_Explosives:
		return []string{}, ""
	case SKL_Flyer:
		return []string{SKL_Airship, SKL_Grav, SKL_Ornithopter, SKL_Rotor, SKL_Wing}, ""
	case SKL_Airship:
		return []string{}, SKL_Flyer
	case SKL_Grav:
		return []string{}, SKL_Flyer
	case SKL_Ornithopter:
		return []string{}, SKL_Flyer
	case SKL_Rotor:
		return []string{}, SKL_Flyer
	case SKL_Wing:
		return []string{}, SKL_Flyer
	case SKL_Gampler:
		return []string{}, ""
	case SKL_Gunner:
		return []string{SKL_Turret, SKL_Ortilery, SKL_Screen, SKL_Capital}, ""
	case SKL_Turret:
		return []string{}, SKL_Gunner
	case SKL_Ortilery:
		return []string{}, SKL_Gunner
	case SKL_Screen:
		return []string{}, SKL_Gunner
	case SKL_Capital:
		return []string{}, SKL_Gunner
	case SKL_GunCombat:
		return []string{SKL_Archaic, SKL_Energy, SKL_Slug}, ""
	case SKL_Archaic:
		return []string{}, SKL_GunCombat
	case SKL_Energy:
		return []string{}, SKL_GunCombat
	case SKL_Slug:
		return []string{}, SKL_GunCombat
	case SKL_HeavyWeapons:
		return []string{SKL_Artilery, SKL_Portable, SKL_Vechicle}, ""
	case SKL_Artilery:
		return []string{}, SKL_HeavyWeapons
	case SKL_Portable:
		return []string{}, SKL_HeavyWeapons
	case SKL_Vechicle:
		return []string{}, SKL_HeavyWeapons
	case SKL_Investigate:
		return []string{}, ""
	case SKL_JackOfAllTrades:
		return []string{}, ""
	case SKL_Language:
		return []string{SKL_Galinglic, SKL_Vilani, SKL_Zdetl, SKL_Oynprith, SKL_Trokh, SKL_Gvegh}, ""
	case SKL_Galinglic:
		return []string{}, SKL_Language
	case SKL_Vilani:
		return []string{}, SKL_Language
	case SKL_Zdetl:
		return []string{}, SKL_Language
	case SKL_Oynprith:
		return []string{}, SKL_Language
	case SKL_Trokh:
		return []string{}, SKL_Language
	case SKL_Gvegh:
		return []string{}, SKL_Language
	case SKL_Leadership:
		return []string{}, ""
	case SKL_Mechanic:
		return []string{}, ""
	case SKL_Medic:
		return []string{}, ""
	case SKL_Melee:
		return []string{SKL_Unarmed, SKL_Blade, SKL_Bludgeon, SKL_Natural}, ""
	case SKL_Unarmed:
		return []string{}, SKL_Melee
	case SKL_Blade:
		return []string{}, SKL_Melee
	case SKL_Bludgeon:
		return []string{}, SKL_Melee
	case SKL_Natural:
		return []string{}, SKL_Melee
	case SKL_Navigation:
		return []string{}, ""
	case SKL_Persuade:
		return []string{}, ""
	case SKL_Pilot:
		return []string{SKL_SmallCraft, SKL_Spacecraft, SKL_CapitalShip}, ""
	case SKL_SmallCraft:
		return []string{}, SKL_Pilot
	case SKL_Spacecraft:
		return []string{}, SKL_Pilot
	case SKL_CapitalShip:
		return []string{}, SKL_Pilot
	case SKL_Profession:
		return []string{SKL_Belter, SKL_Biologicals, SKL_NewProfesion}, ""
	case SKL_Belter:
		return []string{}, SKL_Profession
	case SKL_Biologicals:
		return []string{}, SKL_Profession
	case SKL_NewProfesion:
		return []string{}, SKL_Profession
	case SKL_Recon:
		return []string{}, ""
	case SKL_LifeSciences:
		return []string{SKL_Biology, SKL_Genetics, SKL_Psionicology, SKL_Xenology}, ""
	case SKL_Biology:
		return []string{}, SKL_LifeSciences
	case SKL_Genetics:
		return []string{}, SKL_LifeSciences
	case SKL_Psionicology:
		return []string{}, SKL_LifeSciences
	case SKL_Xenology:
		return []string{}, SKL_LifeSciences
	case SKL_PhysicalSciences:
		return []string{SKL_Chemistry, SKL_Physics, SKL_JumpspacePhysics}, ""
	case SKL_Chemistry:
		return []string{}, SKL_PhysicalSciences
	case SKL_Physics:
		return []string{}, SKL_PhysicalSciences
	case SKL_JumpspacePhysics:
		return []string{}, SKL_PhysicalSciences
	case SKL_RoboticSciences:
		return []string{SKL_Cybernetics, SKL_Robotics}, ""
	case SKL_Cybernetics:
		return []string{}, SKL_RoboticSciences
	case SKL_Robotics:
		return []string{}, SKL_RoboticSciences
	case SKL_SocialSciences:
		return []string{SKL_Archaelogy, SKL_Economics, SKL_History, SKL_Linuistics, SKL_Philosophy, SKL_Psychology, SKL_Sophontology}, ""
	case SKL_Archaelogy:
		return []string{}, SKL_SocialSciences
	case SKL_Economics:
		return []string{}, SKL_SocialSciences
	case SKL_History:
		return []string{}, SKL_SocialSciences
	case SKL_Linuistics:
		return []string{}, SKL_SocialSciences
	case SKL_Philosophy:
		return []string{}, SKL_SocialSciences
	case SKL_Psychology:
		return []string{}, SKL_SocialSciences
	case SKL_Sophontology:
		return []string{}, SKL_SocialSciences
	case SKL_SpaceSciences:
		return []string{SKL_Astronomy, SKL_Cosmology, SKL_Planetology}, ""
	case SKL_Astronomy:
		return []string{}, SKL_SpaceSciences
	case SKL_Cosmology:
		return []string{}, SKL_SpaceSciences
	case SKL_Planetology:
		return []string{}, SKL_SpaceSciences
	case SKL_Seafarer:
		return []string{SKL_OceanShips, SKL_Personal, SKL_Sail, SKL_Submarine}, ""
	case SKL_OceanShips:
		return []string{}, SKL_Seafarer
	case SKL_Personal:
		return []string{}, SKL_Seafarer
	case SKL_Sail:
		return []string{}, SKL_Seafarer
	case SKL_Submarine:
		return []string{}, SKL_Seafarer
	case SKL_Stealth:
		return []string{}, ""
	case SKL_Steward:
		return []string{}, ""
	case SKL_Streetwise:
		return []string{}, ""
	case SKL_Survival:
		return []string{}, ""
	case SKL_Tactics:
		return []string{SKL_Military, SKL_Naval}, ""
	case SKL_Military:
		return []string{}, SKL_Tactics
	case SKL_Naval:
		return []string{}, SKL_Tactics
	case SKL_VaccSuit:
		return []string{}, ""
	default:
		return []string{""}, "bad name"
	}
}
