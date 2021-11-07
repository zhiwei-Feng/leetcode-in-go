package n0055

func canJump(nums []int) bool {
	rightmost := 0
	for i := 0; i < len(nums); i++ {
		if i > rightmost {
			return false
		}
		rightmost = max(rightmost, i+nums[i])
	}
	return true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func canJumpv1(nums []int) bool {
	var (
		n  = len(nums)
		dp = make([]bool, n)
	)

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = true
		}
		if i == n-1 {
			break
		}
		if dp[i] {
			for j := 1; j <= nums[i] && i+j < n; j++ {
				dp[i+j] = true
			}
		}
	}
	return dp[n-1]
}
