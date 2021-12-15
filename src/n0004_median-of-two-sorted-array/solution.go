package n0004

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	total := m + n
	if total%2 == 1 {
		// only find total/2 + 1
		return float64(getKth(nums1, 0, m-1, nums2, 0, n-1, total/2+1))
	} else {
		// find total/2 and total/2 + 1
		r1 := float64(getKth(nums1, 0, m-1, nums2, 0, n-1, total/2))
		r2 := float64(getKth(nums1, 0, m-1, nums2, 0, n-1, total/2+1))
		return (r1 + r2) / 2
	}
}

func getKth(nums1 []int, st1, en1 int, nums2 []int, st2, en2 int, k int) int {
	len1, len2 := en1-st1+1, en2-st2+1
	if len1 > len2 {
		return getKth(nums2, st2, en2, nums1, st1, en1, k)
	}
	if len1 == 0 {
		return nums2[st2+k-1]
	}
	if k == 1 {
		return min(nums1[st1], nums2[st2])
	}

	median1 := st1 + min(k/2, len1) - 1
	median2 := st2 + min(k/2, len2) - 1
	if nums1[median1] < nums2[median2] {
		return getKth(nums1, median1+1, en1, nums2, st2, en2, k-(median1-st1+1))
	} else {
		return getKth(nums1, st1, en1, nums2, median2+1, en2, k-(median2-st2+1))
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
