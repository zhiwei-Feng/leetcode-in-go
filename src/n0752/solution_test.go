package n0752

import "testing"

func Test_openLock(t *testing.T) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"}, -1},
		{"case2", args{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202"}, 6},
		{"case3", args{[]string{"8888"}, "0009"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openLock(tt.args.deadends, tt.args.target); got != tt.want {
				t.Errorf("openLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
