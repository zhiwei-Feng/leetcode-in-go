package n0004

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"case1", args{[]int{1, 3}, []int{2}}, 2.00000},
		{"case2", args{[]int{1, 2}, []int{3, 4}}, 2.50000},
		{"case3", args{[]int{0, 0}, []int{0, 0}}, 0.00000},
		{"case4", args{[]int{}, []int{1}}, 1.00000},
		{"case5", args{[]int{2}, []int{}}, 2.00000},
		 {"case6", args{[]int{1,3}, []int{2,7}}, 2.50000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
