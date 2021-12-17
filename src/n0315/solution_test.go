package n0315

import (
	"reflect"
	"testing"
)

func Test_countSmaller(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[]int{5, 2, 6, 1}}, []int{2, 1, 1, 0}},
		{"case2", args{[]int{-1}}, []int{0}},
		{"case3", args{[]int{-1, -1}}, []int{0, 0}},
		{"case4", args{[]int{2, -1, -1}}, []int{2, 0, 0}},
		{"case5", args{[]int{2, 0, 1}}, []int{2, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSmaller(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmaller() = %v, want %v", got, tt.want)
			}
		})
	}
}
