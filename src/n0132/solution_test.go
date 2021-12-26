package n0132

import "testing"

func Test_minCut(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"cas1", args{"aab"}, 1},
		{"cas2", args{"ab"}, 1},
		{"cas3", args{"efe"}, 0},
		{"cas4", args{"abbaaa"}, 1},
		{"cas5", args{"aaabaa"}, 1},
		{"cas6", args{"ababababababababababababcbabababababababababababa"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
