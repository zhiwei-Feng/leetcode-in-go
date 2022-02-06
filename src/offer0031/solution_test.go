package offer0031

import "testing"

func Test_validateStackSequences(t *testing.T) {
	type args struct {
		pushed []int
		popped []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"c1", args{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}}, true},
		{"c2", args{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}}, false},
		{"c3", args{[]int{0,2,1}, []int{0,1,2}}, true},
		{"c4", args{[]int{1,3,2,0}, []int{1,2,0,3}}, true},
		{"c5", args{[]int{4,0,1,2,3}, []int{4,2,3,0,1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateStackSequences(tt.args.pushed, tt.args.popped); got != tt.want {
				t.Errorf("validateStackSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
