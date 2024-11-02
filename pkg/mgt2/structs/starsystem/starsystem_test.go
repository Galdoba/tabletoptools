package starsystem

import (
	"fmt"
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
			args: args{"G8 V-0.907-0.957-0.681:B-6.1-0.08"},
			want: &Star{
				Type:         "G",
				SubType:      "8",
				Class:        "V",
				Mass:         0.907,
				Diameter:     0.957,
				Luminosity:   0.681,
				Age:          0,
				Designation:  "B",
				OrbitN:       6.1,
				Eccentricity: 0.08,
			},
			wantErr: false,
		}, // TODO: Add test cases.
		{
			name: "3",
			args: args{"D-0.490-0.017-0.000525"},
			want: &Star{
				Type:         "",
				SubType:      "",
				Class:        "D",
				Mass:         0.49,
				Diameter:     0.017,
				Luminosity:   0.000525,
				Age:          0,
				Designation:  "",
				OrbitN:       0,
				Eccentricity: 0,
			},
			wantErr: false,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StarFromProfile(tt.args.s)
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

func TestStar_Profile(t *testing.T) {
	tests := []struct {
		name string
		st   *Star
		want string
	}{
		{
			name: "",
			st: &Star{
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
			want: "G7 V-0.929-0.967-0.738-6.336",
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.st.Profile(); got != tt.want {
				t.Errorf("Star.Profile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromProfile(t *testing.T) {
	type args struct {
		from string
	}
	tests := []struct {
		name    string
		args    args
		want    *StarSystem
		wantErr bool
	}{
		{
			name: "----",
			args: args{"5-G7 V-0.929-0.967-0.738-6.336:Ab-0.09-0.11-G8 V-0.907-0.957-0.681:B-6.1-0.08-K8 V-0.626-0.777-0.136:Ca-12.1-0.47-M0 V-0.510-0.728-0.0895:Cb-0.21-0.24-D-0.490-0.017-0.000525"},
			want: &StarSystem{
				Stars: []*Star{
					{
						Type:         "G",
						SubType:      "7",
						Class:        "V",
						Mass:         0.929,
						Diameter:     0.967,
						Luminosity:   0.738,
						Age:          6.336,
						Designation:  "Ab",
						OrbitN:       0.09,
						Eccentricity: 0.11,
					},
					{
						Type:         "G",
						SubType:      "8",
						Class:        "V",
						Mass:         0.907,
						Diameter:     0.957,
						Luminosity:   0.681,
						Age:          0,
						Designation:  "B",
						OrbitN:       6.1,
						Eccentricity: 0.08,
					},
					{
						Type:         "K",
						SubType:      "8",
						Class:        "V",
						Mass:         0.626,
						Diameter:     0.777,
						Luminosity:   0.136,
						Age:          0,
						Designation:  "Ca",
						OrbitN:       12.1,
						Eccentricity: 0.47,
					},
					{
						Type:         "M",
						SubType:      "0",
						Class:        "V",
						Mass:         0.510,
						Diameter:     0.728,
						Luminosity:   0.0895,
						Age:          0,
						Designation:  "Cb",
						OrbitN:       0.21,
						Eccentricity: 0.24,
					},
					{
						Type:         "",
						SubType:      "",
						Class:        "D",
						Mass:         0.49,
						Diameter:     0.017,
						Luminosity:   0.000525,
						Age:          0,
						Designation:  "",
						OrbitN:       0,
						Eccentricity: 0,
					},
				},
				Age: 6.336,
			},
			wantErr: false,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromProfile(tt.args.from)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				for i, _ := range got.Stars {
					fmt.Println(tt.want.Stars[i].Profile())
					fmt.Println(got.Stars[i].Profile())
				}
				t.Errorf("FromProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
