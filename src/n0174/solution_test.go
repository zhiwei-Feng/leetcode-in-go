package n0174

import "testing"

func Test_calculateMinimumHP(t *testing.T) {
	type args struct {
		dungeon [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}}, 7},
		{"case2", args{[][]int{{100}}}, 1},
		{"case3", args{[][]int{{3, -20, 30}, {-3, 4, 0}}}, 1},
		{"case4", args{[][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMinimumHP(tt.args.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHP() = %v, want %v", got, tt.want)
			}
		})
	}
}
