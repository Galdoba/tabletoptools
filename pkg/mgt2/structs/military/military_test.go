package military

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
		want    *Military
		wantErr bool
	}{
		{
			name: "",
			args: args{"F2C96-000:4.03%"},
			want: &Military{
				Enforcement:   15,
				Militia:       2,
				Army:          12,
				WetNavy:       9,
				AirForce:      6,
				SystemDefence: 0,
				Navy:          0,
				Marines:       0,
				Budget:        4.03,
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

func TestMilitary_Profile(t *testing.T) {
	tests := []struct {
		name string
		mil  *Military
		want string
	}{
		{
			name: "",
			mil: &Military{
				Enforcement:   15,
				Militia:       2,
				Army:          12,
				WetNavy:       9,
				AirForce:      6,
				SystemDefence: 0,
				Navy:          0,
				Marines:       0,
				Budget:        4.03,
			},
			want: "F2C96-000:4.03%",
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mil.Profile(); got != tt.want {
				t.Errorf("Military.Profile() = %v, want %v", got, tt.want)
			}
		})
	}
}
