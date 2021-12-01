package n0286

import "testing"

func Test_wallsAndGates(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{"case1", args{[][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}}}},
		{"case2", args{[][]int{{-1}}}},
		{"case3", args{[][]int{{2147483647}}}},
		{"case4", args{[][]int{{0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallsAndGates(tt.args.rooms)
			t.Log(tt.args.rooms)
		})
	}
}
