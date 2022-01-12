package n1629

import "testing"

func Test_slowestKey(t *testing.T) {
	type args struct {
		releaseTimes []int
		keysPressed  string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"case1", args{[]int{9,29,49,50}, "cbcd"}, 'c'},
		{"case2", args{[]int{2,4,10,16,25,26,28,33,75}, "duxwdgmgw"}, 'w'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slowestKey(tt.args.releaseTimes, tt.args.keysPressed); got != tt.want {
				t.Errorf("slowestKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
