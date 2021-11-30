package n0159

import "testing"

func Test_lengthOfLongestSubstringTwoDistinct(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case1", args{"eceba"}, 3},
		{"Case2", args{"ccaabbb"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringTwoDistinct(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringTwoDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
