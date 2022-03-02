package n0162

func findPeakElement(nums []int) int {
	if len(nums) == 1 || nums[0] > nums[1] {
		return 0
	}
	n := len(nums)
	if nums[n-1] > nums[n-2] {
		return n - 1
	}
	l, r := 1, n-2
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[m-1] && nums[m] > nums[m+1] {
			return m
		}
		if nums[m] < nums[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
