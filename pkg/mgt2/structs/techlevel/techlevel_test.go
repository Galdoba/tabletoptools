package techlevel

import (
	"reflect"
	"testing"
)

func TestFromProfile(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *TechLevel
		wantErr bool
	}{
		{
			name: "",
			args: args{"7-6-77765-7777-77-C"},
			want: &TechLevel{
				HighCommon:       "7",
				LowCommon:        "6",
				Energy:           "7",
				Electronics:      "7",
				Manufactoring:    "7",
				Medical:          "6",
				Enviromental:     "5",
				LandTransport:    "7",
				WaterTransport:   "7",
				AirTransport:     "7",
				SpaceTransport:   "7",
				PersonalMilitary: "7",
				HeavyMilitary:    "7",
				Novelty:          "C",
			},
			wantErr: false,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromProfile(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTechLevel_Profile(t *testing.T) {
	tests := []struct {
		name string
		tl   *TechLevel
		want string
	}{
		{
			name: "",
			tl: &TechLevel{
				HighCommon:       "7",
				LowCommon:        "6",
				Energy:           "7",
				Electronics:      "7",
				Manufactoring:    "7",
				Medical:          "6",
				Enviromental:     "5",
				LandTransport:    "7",
				WaterTransport:   "7",
				AirTransport:     "7",
				SpaceTransport:   "7",
				PersonalMilitary: "7",
				HeavyMilitary:    "7",
				Novelty:          "C",
			},
			want: "7-6-77765-7777-77-C",
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tl.Profile(); got != tt.want {
				t.Errorf("TechLevel.Profile() = %v, want %v", got, tt.want)
			}
		})
	}
}
