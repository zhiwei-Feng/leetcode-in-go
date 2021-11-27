package n5925

import "testing"

func Test_countPyramids(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"Case1", args{[][]int{{0, 1, 1, 0}, {1, 1, 1, 1}}}, 2},
		// {"Case2", args{[][]int{{1, 1, 1}, {1, 1, 1}}}, 2},
		// {"Case3", args{[][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}}, 0},
		// {"Case4", args{[][]int{{1, 1, 1, 1, 0}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {0, 1, 0, 0, 1}}}, 13},
		{"Case5", args{[][]int{{0, 0, 0, 1, 0, 0, 0}, {0, 0, 1, 1, 1, 0, 0}, {0, 1, 1, 1, 1, 0, 0}, {1, 1, 1, 1, 1, 1, 1}}}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPyramids(tt.args.grid); got != tt.want {
				t.Errorf("countPyramids() = %v, want %v", got, tt.want)
			}
		})
	}
}
