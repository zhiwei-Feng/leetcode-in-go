package n1220

import "testing"

func Test_countVowelPermutation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"c1", args{1}, 5},
		{"c2", args{2}, 10},
		{"c3", args{3}, 19},
		{"c4", args{5}, 68},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVowelPermutation(tt.args.n); got != tt.want {
				t.Errorf("countVowelPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
