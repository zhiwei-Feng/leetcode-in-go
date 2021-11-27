package n5923

import "testing"

func Test_minimumBuckets(t *testing.T) {
	type args struct {
		street string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case1", args{"H..H"}, 2},
		{"case1", args{"HH.H.."}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumBuckets(tt.args.street); got != tt.want {
				t.Errorf("minimumBuckets() = %v, want %v", got, tt.want)
			}
		})
	}
}
