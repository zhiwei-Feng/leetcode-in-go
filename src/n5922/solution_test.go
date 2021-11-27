package n5922

import "testing"

func Test_countWords(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]string{"leetcode", "is", "amazing", "as", "is"}, []string{"amazing", "leetcode", "is"}}, 2},
		{"case2", args{[]string{"b","bb","bbb"}, []string{"a","aa","aaa"}}, 0},
		{"case3", args{[]string{"a","ab"}, []string{"a","a","a","ab"}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countWords(tt.args.words1, tt.args.words2); got != tt.want {
				t.Errorf("countWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
