package government

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
		want    *Government
		wantErr bool
	}{
		{
			name: "FromSting 1",
			args: args{"4-FES-LM-JS"},
			want: &Government{
				Code:                 "4",
				Centralisation:       CENTRALISATION_FEDERAL,
				Authority:            AUTHORITY_EXECUTIVE,
				LegislativeStructure: STRUCTURE_MULTIPLE_COUNSILS,
				ExecutiveStructure:   STRUCTURE_SINGLE_COUNSIL,
				JudicialStructure:    STRUCTURE_SINGLE_COUNSIL,
			},
			wantErr: false,
		}, // TODO: Add test cases.
		{
			name: "FromSting 2",
			args: args{"4-FBB-ES-LM-JS"},
			want: &Government{
				Code:                 "4",
				Centralisation:       CENTRALISATION_FEDERAL,
				Authority:            AUTHORITY_BALANCED,
				LegislativeStructure: STRUCTURE_MULTIPLE_COUNSILS,
				ExecutiveStructure:   STRUCTURE_SINGLE_COUNSIL,
				JudicialStructure:    STRUCTURE_SINGLE_COUNSIL,
			},
			wantErr: false,
		}, // TODO: Add test cases.
		{
			name: "FromSting 3",
			args: args{"4-FES"},
			want: &Government{
				Code:               "4",
				Centralisation:     CENTRALISATION_FEDERAL,
				Authority:          AUTHORITY_EXECUTIVE,
				ExecutiveStructure: STRUCTURE_SINGLE_COUNSIL,
			},
			wantErr: false,
		}, // TODO: Add test cases.
		{
			name: "FromSting 4",
			args: args{"4-FBB-LM-JS-ES"},
			want: &Government{
				Code:                 "4",
				Centralisation:       CENTRALISATION_FEDERAL,
				Authority:            AUTHORITY_BALANCED,
				LegislativeStructure: STRUCTURE_MULTIPLE_COUNSILS,
				ExecutiveStructure:   STRUCTURE_SINGLE_COUNSIL,
				JudicialStructure:    STRUCTURE_SINGLE_COUNSIL,
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

func TestGovernment_String(t *testing.T) {
	tests := []struct {
		name string
		gv   *Government
		want string
	}{
		{
			name: "string 1",
			gv: &Government{
				Code:                 "4",
				Centralisation:       CENTRALISATION_FEDERAL,
				Authority:            AUTHORITY_BALANCED,
				LegislativeStructure: STRUCTURE_MULTIPLE_COUNSILS,
				ExecutiveStructure:   STRUCTURE_SINGLE_COUNSIL,
				JudicialStructure:    STRUCTURE_SINGLE_COUNSIL,
			},
			want: "4-FBB-ES-LM-JS",
		}, // TODO: Add test cases.
		{
			name: "string 2",
			gv: &Government{
				Code:                 "4",
				Centralisation:       CENTRALISATION_FEDERAL,
				Authority:            AUTHORITY_EXECUTIVE,
				LegislativeStructure: STRUCTURE_MULTIPLE_COUNSILS,
				ExecutiveStructure:   STRUCTURE_SINGLE_COUNSIL,
				JudicialStructure:    STRUCTURE_SINGLE_COUNSIL,
			},
			want: "4-FES-LM-JS",
		}, // TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gv.Profile(); got != tt.want {
				t.Errorf("Government.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
