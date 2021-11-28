package n1231

import "testing"

func Test_maximizeSweetness(t *testing.T) {
	type args struct {
		sweetness []int
		k         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5}, 6},
		{"case2", args{[]int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 8}, 1},
		{"case3", args{[]int{1, 2, 2, 1, 2, 2, 1, 2, 2}, 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximizeSweetness(tt.args.sweetness, tt.args.k); got != tt.want {
				t.Errorf("maximizeSweetness() = %v, want %v", got, tt.want)
			}
		})
	}
}
