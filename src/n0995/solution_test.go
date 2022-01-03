package n0995

import "testing"

func Test_minKBitFlips(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case1", args{[]int{0, 1, 0}, 1}, 2},
		{"Case2", args{[]int{1, 1, 0}, 2}, -1},
		{"Case3", args{[]int{0,0,0,1,0,1,1,0}, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minKBitFlips(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minKBitFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
