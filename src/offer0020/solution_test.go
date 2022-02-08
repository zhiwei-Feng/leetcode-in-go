package offer0020

import "testing"

func Test_isNumber(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"c1", args{"+100"}, true},
		{"c2", args{"5e2"}, true},
		{"c3", args{"3.1416"}, true},
		{"c4", args{"-1E-16"}, true},
		{"c5", args{"0123"}, true},
		{"c6", args{"1a3.14"}, false},
		{"c7", args{"1.2.3"}, false},
		{"c8", args{"+-5"}, false},
		{"c9", args{"12e+5.4"}, false},
		{"c10", args{"e-5"}, false},
		{"c11", args{"1 "}, true},
		{"c12", args{" "}, false},
		{"c13", args{".1"}, true},
		{"c14", args{"."}, false},
		{"c15", args{".e1"}, false},
		{"c16", args{"0e"}, false},
		{"c17", args{"4e+"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumber(tt.args.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
