package n0152

func maxProduct(nums []int) int {
	var ans = nums[0]
	n := len(nums)
	fmax, fmin := nums[0], nums[0]
	for i := 1; i < n; i++ {
		x := nums[i]
		fmax, fmin = max(max(fmax*x, fmin*x), x), min(min(fmax*x, fmin*x), x)
		ans = max(ans, fmax)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
