package culture

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
		want    *Culture
		wantErr bool
	}{
		{
			name: "from profile 1",
			args: args{"6999-B656"},
			want: &Culture{
				Diversty:        "6",
				Xenophilia:      "9",
				Uniqueness:      "9",
				Symbology:       "9",
				Cohesion:        "B",
				Progressiveness: "6",
				Expansionism:    "5",
				Militancy:       "6",
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

func TestCulture_Profile(t *testing.T) {
	tests := []struct {
		name string
		clt  *Culture
		want string
	}{
		{
			name: "culture 1",
			clt: &Culture{
				Diversty:        "6",
				Xenophilia:      "9",
				Uniqueness:      "9",
				Symbology:       "9",
				Cohesion:        "B",
				Progressiveness: "6",
				Expansionism:    "5",
				Militancy:       "6",
			},
			want: "6999-B656",
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.clt.Profile(); got != tt.want {
				t.Errorf("Culture.Profile() = %v, want %v", got, tt.want)
			}
		})
	}
}
