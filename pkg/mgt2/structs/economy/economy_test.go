package economy

import (
	"reflect"
	"testing"
)

func TestEconomy_Profile(t *testing.T) {
	tests := []struct {
		name string
		eco  *Economy
		want string
	}{
		{
			name: "string 1",
			eco: &Economy{
				TradeCodes:           []string{"Ag", "Ri"},
				Importance:           1,
				ResourceFactor:       15,
				LaborFactor:          6,
				InfrastructureFactor: 13,
				EfficiencyFactor:     -2,
				RU:                   -2340,
				PCGWP:                5940,
				WTN:                  8,
				InequalityRating:     58,
				DevelopmentScore:     2.49,
			},
			want: "Ag Ri, +1, F6D-2, -2340.0, 5940.0, 8, 58, 2.49",
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.eco.Profile(); got != tt.want {
				t.Errorf("Economy.Profile() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestFromProfile(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *Economy
		wantErr bool
	}{
		{
			name: "from 1",
			args: args{"Ag Ri, +1, F6D-2, -2340.0, 5940.0, 8, 58, 2.49"},
			want: &Economy{
				TradeCodes:           []string{"Ag", "Ri"},
				Importance:           1,
				ResourceFactor:       15,
				LaborFactor:          6,
				InfrastructureFactor: 13,
				EfficiencyFactor:     -2,
				RU:                   -2340,
				PCGWP:                5940,
				WTN:                  8,
				InequalityRating:     58,
				DevelopmentScore:     2.49,
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
