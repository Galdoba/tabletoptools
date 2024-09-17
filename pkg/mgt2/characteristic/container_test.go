package characteristic

import "testing"

func TestSet_ImportPreset(t *testing.T) {
	type args struct {
		presetName string
	}
	tests := []struct {
		name    string
		cs      *Set
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			cs:      &Set{},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cs.ImportPreset(tt.args.presetName); (err != nil) != tt.wantErr {
				t.Errorf("Set.ImportPreset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
