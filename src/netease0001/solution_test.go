package netease0001

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{[]int{1, 6, 3, 4}}, []int{4, 1, 6, 3}},
		{"2", args{[]int{2, 1, 5, 6, 3, 4}}, []int{2, 5, 4, 1, 6, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
