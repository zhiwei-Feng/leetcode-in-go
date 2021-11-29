package n1868

import (
	"reflect"
	"testing"
)

func Test_findRLEArray(t *testing.T) {
	type args struct {
		encoded1 [][]int
		encoded2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"case1", args{[][]int{{1, 3}, {2, 3}}, [][]int{{6, 3}, {3, 3}}}, [][]int{{6, 6}}},
		{"case2", args{[][]int{{1, 3}, {2, 1}, {3, 2}}, [][]int{{2, 3}, {3, 3}}}, [][]int{{2, 3}, {6, 1}, {9, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRLEArray(tt.args.encoded1, tt.args.encoded2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRLEArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
