package n0149

import "testing"

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[][]int{{1, 1}, {2, 2}, {3, 3}}}, 3},
		{"case2", args{[][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}}, 4},
		{"case3", args{[][]int{{0,0}}}, 1},
		{"case4", args{[][]int{{0,0},{4,5},{7,8},{8,9},{5,6},{3,4},{1,1}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
