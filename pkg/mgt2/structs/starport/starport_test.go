package starport

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
		want    *Starport
		wantErr bool
	}{
		{
			name: "",
			args: args{"C-HN:DY:-1"},
			want: &Starport{
				Class:              "C",
				HighPortPresense:   false,
				DownPortPresense:   true,
				AdjustedImportance: -1,
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

func TestStarport_Profile(t *testing.T) {
	tests := []struct {
		name string
		sp   *Starport
		want string
	}{
		{
			name: "",
			sp: &Starport{
				Class:              "C",
				HighPortPresense:   false,
				DownPortPresense:   true,
				AdjustedImportance: -1,
			},
			want: "C-HN:DY:-1",
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sp.Profile(); got != tt.want {
				t.Errorf("Starport.Profile() = %v, want %v", got, tt.want)
			}
		})
	}
}
