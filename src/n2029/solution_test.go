package n2029

import "testing"

func Test_stoneGameIX(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[]int{2, 1}}, true},
		{"case2", args{[]int{2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameIX(tt.args.stones); got != tt.want {
				t.Errorf("stoneGameIX() = %v, want %v", got, tt.want)
			}
		})
	}
}
