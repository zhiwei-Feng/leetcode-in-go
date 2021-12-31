package n1136

import "testing"

func Test_minimumSemesters(t *testing.T) {
	type args struct {
		n         int
		relations [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{3, [][]int{{1, 3}, {2, 3}}}, 2},
		{"case2", args{3, [][]int{{1, 2}, {2, 3}, {3, 1}}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSemesters(tt.args.n, tt.args.relations); got != tt.want {
				t.Errorf("minimumSemesters() = %v, want %v", got, tt.want)
			}
		})
	}
}
