package offer0033

import "testing"

func Test_verifyPostorder(t *testing.T) {
	type args struct {
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[]int{1,6,3,2,5}}, false},
		{"case2", args{[]int{1,3,2,6,5}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPostorder(tt.args.postorder); got != tt.want {
				t.Errorf("verifyPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
