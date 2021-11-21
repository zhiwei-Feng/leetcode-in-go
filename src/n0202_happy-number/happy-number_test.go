package n0202

import (
	"testing"
)

func Test_happyChange(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{1}, 1},
		{"case2", args{19}, 82},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := happyChange(tt.args.n); got != tt.want {
				t.Errorf("happyChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{19}, true},
		{"case2", args{2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
