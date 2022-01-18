package n0539

import "testing"

func Test_findMinDifference(t *testing.T) {
	type args struct {
		timePoints []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case1", args{[]string{"23:59", "00:00"}}, 1},
		{"Case2", args{[]string{"00:00", "23:59", "00:00"}}, 0},
		{"Case3", args{[]string{"01:01", "02:01", "03:00"}}, 59},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinDifference(tt.args.timePoints); got != tt.want {
				t.Errorf("findMinDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
