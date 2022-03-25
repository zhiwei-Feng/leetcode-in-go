package bytedance0006

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
		{"1", args{[]int{3,4,5,1,2}}, 5},
		{"2", args{[]int{4,5,6,7,0,1,2}}, 7},
		{"3", args{[]int{1,2,3}}, 3},
		{"4", args{[]int{3,2,1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.nums); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
