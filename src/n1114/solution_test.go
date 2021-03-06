package n1114

import "testing"

func Test_print(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{[]int{1,2,3}}},
		{"2", args{[]int{1,3,2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			print(tt.args.nums)
		})
	}
}
