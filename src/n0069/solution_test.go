package n0069

import (
	"testing"
)

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"1", args{4}, 2},
		{"2", args{8}, 2},
		{"3", args{9}, 3},
		{"4", args{17}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySqrtFloat(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{8}, 2.8284271248},
		{"1", args{9}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrtFloat(tt.args.x); got != tt.want {
				t.Errorf("mySqrtFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}
