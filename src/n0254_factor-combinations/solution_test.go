package n0254

import (
	"reflect"
	"testing"
)

func Test_getFactors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{12}, [][]int{{2, 2, 3}, {2, 6}, {3, 4}}},
		{"case2", args{37}, [][]int{}},
		{"case3", args{32}, [][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 4}, {2, 2, 8}, {2, 4, 4}, {2, 16}, {4, 8}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFactors(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFactors() = %v, want %v", got, tt.want)
			}
		})
	}
}
