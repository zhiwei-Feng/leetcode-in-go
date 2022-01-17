package offer0042

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	ans := nums[0]
	for i, v := range nums {
		if i == 0 {
			dp[i] = v
		} else {
			dp[i] = max(dp[i-1]+v, v)
		}
		ans = max(dp[i], ans)
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
