package interview0019

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{"aa","a"}, false},
		{"case2", args{"aa","a*"}, true},
		{"case3", args{"ab",".*"}, true},
		{"case4", args{"aab","c*a*b"}, true},
		{"case5", args{"mississippi","mis*is*p*."}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
