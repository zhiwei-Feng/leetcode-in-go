package n0322

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]int{1}, 0}, 0},
		{"case2", args{[]int{1}, 1}, 1},
		{"case3", args{[]int{1}, 2}, 2},
		{"case4", args{[]int{1, 2, 5}, 11}, 3},
		{"case5", args{[]int{2, 5, 10, 1}, 27}, 4},
		{"case6", args{[]int{186, 419, 83, 408}, 6249}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
