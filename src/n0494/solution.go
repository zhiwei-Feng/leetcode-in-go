package n0494

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	diff := sum - target
	if diff < 0 || diff%2 != 0 {
		return 0
	}
	realTarget := diff / 2
	dp := make([]int, realTarget+1)
	if nums[0] == 0 {
		dp[0] = 2
	} else {
		dp[0] = 1
	}
	for i := 1; i <= realTarget; i++ {
		if nums[0] == i {
			dp[nums[0]] = 1
		}
	}
	for i := 1; i < len(nums); i++ {
		for j := realTarget; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[realTarget]
}
