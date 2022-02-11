package n0003

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"1", args{"abcabcbb"}, 3},
		{"2", args{"bbbbb"}, 1},
		{"3", args{"pwwkew"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := lengthOfLongestSubstring(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
