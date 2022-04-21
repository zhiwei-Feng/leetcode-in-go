package offerii0001

import "testing"

func Test_divide(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{15, 2}, 7},
		{"2", args{7, -3}, -2},
		{"3", args{0, 1}, 0},
		{"4", args{1, 1}, 1},
		{"5", args{1, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
