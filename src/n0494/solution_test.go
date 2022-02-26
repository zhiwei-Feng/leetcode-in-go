package n0494

import "testing"

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"1", args{[]int{1, 1, 1, 1, 1}, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findTargetSumWays(tt.args.nums, tt.args.target); gotAns != tt.wantAns {
				t.Errorf("findTargetSumWays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
