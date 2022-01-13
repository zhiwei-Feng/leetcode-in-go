package offer0011

import "testing"

func Test_minArray(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"case1", args{[]int{3, 4, 5, 1, 2}}, 1},
		// {"case2", args{[]int{2, 2, 2, 0, 1}}, 0},
		// {"case3", args{[]int{1, 2, 3, 4, 5, 6}}, 1},
		{"case4", args{[]int{3, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minArray(tt.args.numbers); got != tt.want {
				t.Errorf("minArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
