package n1115

import "testing"

func Test_printAlternately(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintAlternately(tt.args.n)
		})
	}
}
