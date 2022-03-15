package bytedance0003

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		N       int
		M       int
		winTime []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{3,2,[]int{2,2,5}}, 2},
		{"2", args{3,3,[]int{2,2,11}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.N, tt.args.M, tt.args.winTime); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
