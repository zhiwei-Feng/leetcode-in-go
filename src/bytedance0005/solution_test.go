package bytedance0005

import (
	"math"
	"testing"
)

func Test_sqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{5}, math.Sqrt(5)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sqrt(tt.args.x); got != tt.want {
				t.Errorf("sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
