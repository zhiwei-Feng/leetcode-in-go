package n1182

import (
	"reflect"
	"testing"
)

func Test_shortestDistanceColor(t *testing.T) {
	type args struct {
		colors  []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{1, 1, 2, 1, 3, 2, 2, 3, 3}, [][]int{{1, 3}, {2, 2}, {6, 1}}}, []int{3, 0, 3}},
		{"case2", args{[]int{1, 2}, [][]int{{0, 3}}}, []int{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistanceColor(tt.args.colors, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestDistanceColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
