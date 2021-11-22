package n1901

import (
	"reflect"
	"testing"
)

func Test_findPeakGrid(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[][]int{{1,4},{3,2}}}, []int{1,0}},
		{"case2", args{[][]int{{10,20,15},{21,30,14},{7,16,32}}}, []int{1,1}},
		{"case3", args{[][]int{{36,15,44,21,50},{50,4,15,3,21},{26,18,5,14,37}}}, []int{0,2}},
		{"case4", args{[][]int{{7, 2, 3, 1, 2}, {6, 5, 4, 2, 1}}}, []int{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakGrid(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPeakGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
