package skill

import "testing"

func Test_parseGain(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		// TODO: Add test cases.
		{name: "simple 1", args: args{"Admin 3"}, want: "Admin", want1: 3},
		{name: "simple 2", args: args{"Admin"}, want: "Admin", want1: 1},
		{name: "complex 1", args: args{"Admin (Drive)"}, want: "Drive", want1: 1},
		{name: "complex 2", args: args{"Admin (Drive) 4"}, want: "Drive", want1: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseGain(tt.args.value)
			if got != tt.want {
				t.Errorf("parseGain() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseGain() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
