package offerii0009

func numSubarrayProductLessThanK(nums []int, k int) int {
	winS := 0
	winProd := 1
	ans := 0
	for winE := 0; winE < len(nums); winE++ {
		winProd *= nums[winE]
		for winProd >= k && winS <= winE {
			winProd /= nums[winS]
			winS++
		}
		ans += winE - winS + 1
	}
	return ans
}
