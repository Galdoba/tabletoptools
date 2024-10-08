package skill

import (
	"reflect"
	"testing"

	"github.com/Galdoba/tabletoptools/pkg/mgt2/key"
)

func TestNew(t *testing.T) {
	type args struct {
		name    string
		options []SkillOption
	}

	tests := []struct {
		name    string
		args    args
		want    *skill
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "simple Admin", args: args{name: key.SKL_Admin}, want: &skill{name: "Admin", parent: "", specialities: []string{}, effectiveScore: 0, maxScore: 15}, wantErr: false},
		{name: "simple Admin with max score limit", args: args{key.SKL_Admin, []SkillOption{MaxScore(4)}}, want: &skill{name: "Admin", parent: "", specialities: []string{}, effectiveScore: 0, maxScore: 4}, wantErr: false},
		{name: "cascade Tactics", args: args{name: key.SKL_Tactics}, want: &skill{name: "Tactics", parent: "", specialities: []string{key.SKL_Military, key.SKL_Naval}, effectiveScore: 0, maxScore: 15}, wantErr: false},
		{name: "cascade Tactics Military", args: args{name: key.SKL_Military}, want: &skill{name: "Tactics (Military)", parent: key.SKL_Tactics, specialities: []string{}, effectiveScore: 0, maxScore: 15}, wantErr: false},
		{name: "bad", args: args{name: "bad key"}, want: nil, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.name, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func list() []string {
// 	return []string{
// 		Admin,
// 		Advocate,
// 		Animals,
// 		Handling,
// 		Veterinary,
// 		Training,
// 		Art,
// 		Performer,
// 		Holography,
// 		Instrument,
// 		VisualMedia,
// 		Write,
// 		Astrogation,
// 		Athletics,
// 		Dexterity,
// 		Endurance,
// 		Strenght,
// 		Broker,
// 		Carouse,
// 		Deception,
// 		Diplomat,
// 		Drive,
// 		Hovercraft,
// 		Mole,
// 		Track,
// 		Walker,
// 		Wheel,
// 		Electronics,
// 		Comms,
// 		Computers,
// 		RemoteOps,
// 		Sensors,
// 		Engineer,
// 		MDrive,
// 		JDrive,
// 		LifeSupport,
// 		Power,
// 		Explosives,
// 		Flyer,
// 		Airship,
// 		Grav,
// 		Ornithopter,
// 		Rotor,
// 		Wing,
// 		Gampler,
// 		Gunner,
// 		Turret,
// 		Ortilery,
// 		Screen,
// 		Capital,
// 		GunCombat,
// 		Archaic,
// 		Energy,
// 		Slug,
// 		HeavyWeapons,
// 		Artilery,
// 		Portable,
// 		Vechicle,
// 		Investigate,
// 		JackOfAllTrades,
// 		Language,
// 		Galinglic,
// 		Vilani,
// 		Zdetl,
// 		Oynprith,
// 		Trokh,
// 		Gvegh,
// 		Leadership,
// 		Mechanic,
// 		Medic,
// 		Melee,
// 		Unarmed,
// 		Blade,
// 		Bludgeon,
// 		Natural,
// 		Navigation,
// 		Persuade,
// 		Pilot,
// 		SmallCraft,
// 		Spacecraft,
// 		CapitalShip,
// 		Profession,
// 		Belter,
// 		Biologicals,
// 		NewProfesion,
// 		Recon,
// 		LifeSciences,
// 		Biology,
// 		Genetics,
// 		Psionicology,
// 		Xenology,
// 		PhysicalSciences,
// 		Chemistry,
// 		Physics,
// 		JumpspacePhysics,
// 		RoboticSciences,
// 		Cybernetics,
// 		Robotics,
// 		SocialSciences,
// 		Archaelogy,
// 		Economics,
// 		History,
// 		Linuistics,
// 		Philosophy,
// 		Psychology,
// 		Sophontology,
// 		SpaceSciences,
// 		Astronomy,
// 		Cosmology,
// 		Planetology,
// 		Seafarer,
// 		OceanShips,
// 		Personal,
// 		Sail,
// 		Submarine,
// 		Stealth,
// 		Steward,
// 		Streetwise,
// 		Survival,
// 		Tactics,
// 		Military,
// 		Naval,
// 		VaccSuit,
// 	}
// }
