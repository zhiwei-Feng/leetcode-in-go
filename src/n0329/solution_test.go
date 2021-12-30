package n0329

import "testing"

func Test_longestIncreasingPath(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}}, 4},
		{"case2", args{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}}, 4},
		{"case3", args{[][]int{{1}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath(tt.args.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
