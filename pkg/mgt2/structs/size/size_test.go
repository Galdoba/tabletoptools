package size

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
		want    *Size
		wantErr bool
	}{
		{
			name: "",
			args: args{"5-8163km-1.03-0.66-0.27"},
			want: &Size{
				Code:       "5",
				DiameterKm: 8163,
				Density:    1.03,
				Gravity:    0.66,
				Mass:       0.27,
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
