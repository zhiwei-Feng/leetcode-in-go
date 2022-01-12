package n0334

import "testing"

func Test_increasingTriplet(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1", args{[]int{1, 2, 3, 4, 5}}, true},
		{"case2", args{[]int{5, 4, 3, 2, 1}}, false},
		{"case3", args{[]int{2, 1, 5, 0, 4, 6}}, true},
		{"case4", args{[]int{20,100,10,12,5,13}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingTriplet(tt.args.nums); got != tt.want {
				t.Errorf("increasingTriplet() = %v, want %v", got, tt.want)
			}
		})
	}
}
