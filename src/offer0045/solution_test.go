package offer0045

import "testing"

func Test_minNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"c1", args{[]int{10, 2}}, "102"},
		{"c2", args{[]int{3, 30, 34, 5, 9}}, "3033459"},
		{"c3", args{[]int{0, 30, 34, 5, 9}}, "0303459"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
