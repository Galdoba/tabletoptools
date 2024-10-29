package population

import (
	"reflect"
	"testing"
)

func TestPopulation_String(t *testing.T) {
	tests := []struct {
		name string
		p    *Population
		want string
	}{
		{
			name: "String 1",
			p:    &Population{Code: "7", Value: 1.8, PCR: "3", Urbanization: 39, MajorSettlements: 7},
			want: "7-1.8-3-39-7",
		}, // TODO: Add test cases.
		{
			name: "String 2",
			p:    &Population{Code: "7", Value: 1.80, PCR: "3", Urbanization: 39, MajorSettlements: 7},
			want: "7-1.8-3-39-7",
		}, // TODO: Add test cases.
		{
			name: "String 3",
			p:    &Population{Code: "7", Value: 1.80, PCR: "3", Urbanization: 100, MajorSettlements: 7},
			want: "7-1.8-3->99-7",
		}, // TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Profile(); got != tt.want {
				t.Errorf("Population.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *Population
		wantErr bool
	}{
		{
			name:    "FromString 1",
			args:    args{"7-1.8-3-39-7"},
			want:    &Population{Code: "7", Value: 1.8, PCR: "3", Urbanization: 39, MajorSettlements: 7},
			wantErr: false,
		},
		{
			name:    "FromString 2",
			args:    args{"5-1.4-8-9-2"},
			want:    &Population{Code: "5", Value: 1.4, PCR: "8", Urbanization: 9, MajorSettlements: 2},
			wantErr: false,
		}, // TODO: Add test cases.
		{
			name:    "FromString 3",
			args:    args{"5-1.46-8->99-2"},
			want:    &Population{Code: "5", Value: 1.46, PCR: "8", Urbanization: 100, MajorSettlements: 2},
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
