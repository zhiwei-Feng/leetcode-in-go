package bytedance0004

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {3, 4}, {1, 2}, {2, 4}, {2, 5}, {4, 5}}}, [][]int{
			{0, 1, 2, 3, 4, 5},
			{0, 1, 3, 2, 4, 5},
			{0, 3, 1, 2, 4, 5},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
