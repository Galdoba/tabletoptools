package characteristic

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name    string
		options []CharacteristicOption
	}
	tests := []struct {
		name    string
		args    args
		want    *characteristic
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Effective Score Valid", args: args{STR, []CharacteristicOption{EffectiveScore(5), MaxScore(5)}}, want: &characteristic{name: STR, abb: "STR", code: "C1", cType: physical, effectiveScore: 5, maxScore: 5, creationMod: 0, creationDice: 2}, wantErr: false},
		{name: "Effective Score Invalid 1", args: args{STR, []CharacteristicOption{EffectiveScore(5), MaxScore(4)}}, want: nil, wantErr: true},
		{name: "Effective Score Invalid 2", args: args{STR, []CharacteristicOption{MaxScore(-4)}}, want: nil, wantErr: true},
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
