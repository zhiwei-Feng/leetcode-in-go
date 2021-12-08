package n0394

import (
	"strings"
	"testing"
)

func Test_decodeString(t *testing.T) {

	var case5Res strings.Builder
	for i := 0; i < 100; i++ {
		case5Res.WriteString("leetcode")
	}
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"case1", args{"3[a]2[bc]"}, "aaabcbc"},
		{"case2", args{"3[a2[c]]"}, "accaccacc"},
		{"case3", args{"2[abc]3[cd]ef"}, "abcabccdcdcdef"},
		{"case4", args{"abc3[cd]xyz"}, "abccdcdcdxyz"},
		{"case5", args{"100[leetcode]"}, case5Res.String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
