package n0201

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case1", args{5, 7}, 4},
		{"Case2", args{0, 0}, 0},
		{"Case3", args{1, 2147483647}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
