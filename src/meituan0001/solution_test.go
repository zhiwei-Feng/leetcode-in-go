package meituan0001

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 2, 3, 5, 6, 7}}, 4},
		{"2", args{[]int{1, 2, 3, 4, 5, 6, 7, 8}}, 4},
		{"3", args{[]int{1, 2, 3, 4, 5, 6}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.nums); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
