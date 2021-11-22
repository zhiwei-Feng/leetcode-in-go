package n1060

import "testing"

func Test_missingElement(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{4, 7, 9, 10}, 1}, 5},
		{"case2", args{[]int{4, 7, 9, 10}, 3}, 8},
		{"case3", args{[]int{1, 2, 4}, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingElement(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("missingElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
