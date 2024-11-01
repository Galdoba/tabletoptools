package starsystem

import (
	"reflect"
	"testing"
)

/*
{
			name: "1",
			args: args{"G7 V-0.929-0.967-0.738-6.336:A-2-3"},
			want: &Star{
				Type:         "G",
				SubType:      "7",
				Class:        "V",
				Mass:         0.929,
				Diameter:     0.967,
				Luminosity:   0.738,
				Age:          6.336,
				Designation:  "A",
				OrbitN:       2,
				Eccentricity: 3,
			},
			wantErr: false,
		}, // TODO: Add test cases.
		{
			name: "2",
			args: args{"G7 V-0.929-0.967-0.738-6.336"},
			want: &Star{
				Type:         "G",
				SubType:      "7",
				Class:        "V",
				Mass:         0.929,
				Diameter:     0.967,
				Luminosity:   0.738,
				Age:          6.336,
				Designation:  "",
				OrbitN:       0,
				Eccentricity: 0,
			},
			wantErr: false,
		}, // TODO: Add test cases.
	}
*/

func Test_parseStar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *Star
		wantErr bool
	}{
		{
			name: "1",
			args: args{"G7 V-0.929-0.967-0.738-6.336:A-2-3"},
			want: &Star{
				Type:         "G",
				SubType:      "7",
				Class:        "V",
				Mass:         0.929,
				Diameter:     0.967,
				Luminosity:   0.738,
				Age:          6.336,
				Designation:  "A",
				OrbitN:       2,
				Eccentricity: 3,
			},
			wantErr: false,
		}, // TODO: Add test cases.
		{
			name: "2",
			args: args{"G7 V-0.929-0.967-0.738-6.336"},
			want: &Star{
				Type:         "G",
				SubType:      "7",
				Class:        "V",
				Mass:         0.929,
				Diameter:     0.967,
				Luminosity:   0.738,
				Age:          6.336,
				Designation:  "",
				OrbitN:       0,
				Eccentricity: 0,
			},
			wantErr: false,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseStar(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseStar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseStar() = %v, want %v", got, tt.want)
			}
		})
	}
}
