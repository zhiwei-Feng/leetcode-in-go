package bytedance0002

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		n    int
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{23132, []int{2,4,9}}, 22999},
		{"2", args{21132, []int{2,4,9}}, 9999},
		{"3", args{24132, []int{2,4,9}}, 22999},
		{"4", args{24132, []int{1,4,9}}, 19999},
		{"5", args{14132, []int{2,4,9}}, 9999},
		{"5", args{14132, []int{1,4,9}}, 14119},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.nums); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
