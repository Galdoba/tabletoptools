package characteristic

import (
	"reflect"
	"testing"
)

func Test_characteristic_Encode(t *testing.T) {
	tests := []struct {
		name string
		ch   *characteristic
		want string
	}{
		// TODO: Add test cases.
		{name: "mod positive", ch: &characteristic{name: STR, abb: "STR", code: C1, cType: physical, effectiveScore: 12, maxScore: 15, creationMod: 4, creationDice: 2}, want: "Strenght : 12/15 (2d6+4)"},
		{name: "mod negative", ch: &characteristic{name: STR, abb: "STR", code: C1, cType: physical, effectiveScore: 12, maxScore: 15, creationMod: -4, creationDice: 2}, want: "Strenght : 12/15 (2d6-4)"},
		{name: "max score mot match", ch: &characteristic{name: STR, abb: "STR", code: C1, cType: physical, effectiveScore: 12, maxScore: 15, creationMod: -4, creationDice: 2}, want: "Strenght : 12/15 (2d6-4)"},
		{name: "max score match", ch: &characteristic{name: STR, abb: "STR", code: C1, cType: physical, effectiveScore: 12, maxScore: 12, creationMod: -4, creationDice: 2}, want: "Strenght : 12 (2d6-4)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ch.Encode(); got != tt.want {
				t.Errorf("characteristic.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    *characteristic
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "decode 1", args: args{str: "Strenght : 12/15 (2d6+4)"}, want: &characteristic{name: STR, abb: "STR", code: C1, cType: physical, effectiveScore: 12, maxScore: 15, creationMod: 4, creationDice: 2}, wantErr: false},
		{name: "decode 2", args: args{str: "Strenght : 12/15 (2d6-4)"}, want: &characteristic{name: STR, abb: "STR", code: C1, cType: physical, effectiveScore: 12, maxScore: 15, creationMod: -4, creationDice: 2}, wantErr: false},
		{name: "decode 3", args: args{str: "Strenght : 12/15 (2d6)"}, want: &characteristic{name: STR, abb: "STR", code: C1, cType: physical, effectiveScore: 12, maxScore: 15, creationMod: 0, creationDice: 2}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
