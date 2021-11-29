package n0287findtheduplicatenumber

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 3, 4, 2, 2}}, 2},
		{"case2", args{[]int{3, 1, 3, 4, 2}}, 3},
		{"case3", args{[]int{1, 1}}, 1},
		{"case4", args{[]int{1, 1, 2}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
