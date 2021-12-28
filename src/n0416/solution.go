package n0416

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	total := 0
	for _, v := range nums {
		total += v
	}
	if total%2 != 0 {
		return false
	}
	half := total / 2
	dp := make([]bool, half+1)
	dp[0] = true
	for i := 1; i <= half; i++ {
		dp[i] = i == nums[0]
	}

	for i := 1; i < len(nums); i++ {
		for j := half; j >= 1; j-- {
			if j > nums[i] {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}

	return dp[half]
}

func canPartition_twodim_dp(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	total := 0
	for _, v := range nums {
		total += v
	}
	if total%2 != 0 {
		return false
	}
	half := total / 2
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, half+1)
		dp[i][0] = true
	}
	if nums[0] <= half {
		dp[0][nums[0]] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= half; j++ {
			if j > nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)-1][half]
}
