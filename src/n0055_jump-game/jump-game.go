package n0055

func canJump(nums []int) bool {
	maxR := 0
	i:=0
	for ;i<len(nums)&&i<=maxR;i++{
		if i+nums[i]>maxR{
			maxR = i+nums[i]
		}
	}
	return i==len(nums)
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
