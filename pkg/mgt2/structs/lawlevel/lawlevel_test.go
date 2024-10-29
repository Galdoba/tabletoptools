package lawlevel

import (
	"reflect"
	"testing"
)

func TestFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *LawLevel
		wantErr bool
	}{
		{
			name: "",
			args: args{"8-8AA68"},
			want: &LawLevel{
				Code:     "8",
				Weapons:  "8",
				Economic: "A",
				Criminal: "A",
				Private:  "6",
				Personal: "8",
			},
			wantErr: false,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromProfile(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLawLevel_String(t *testing.T) {
	tests := []struct {
		name string
		ll   *LawLevel
		want string
	}{
		{
			name: "",
			ll: &LawLevel{
				Code:     "8",
				Weapons:  "8",
				Economic: "A",
				Criminal: "A",
				Private:  "6",
				Personal: "8",
			},
			want: "8-8AA68",
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ll.Profile(); got != tt.want {
				t.Errorf("LawLevel.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
