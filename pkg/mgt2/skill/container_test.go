package skill

import (
	"testing"
)

func TestSkillSet_Gain(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		sks     *SkillSet
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		//		{name: "gain no value", sks: &SkillSet{SklByCode: map[string]*skill{Admin: {name: Admin, parent: "", specialities: []string{}, effectiveScore: 1, maxScore: 4}}}, args: args{Admin}, wantErr: false},
		//		{name: "gain value 2", sks: &SkillSet{SklByCode: map[string]*skill{Admin: {name: Admin, parent: "", specialities: []string{}, effectiveScore: 2, maxScore: 4}}}, args: args{"Admin 3"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sks.Gain(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("SkillSet.Gain() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
