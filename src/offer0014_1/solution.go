package offer0014_1

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3

	for i := 4; i <= n; i++ {
		for j := 1; j < i; j++ {
			prod := dp[j] * dp[i-j]
			if prod > dp[i] {
				dp[i] = prod
			}
		}
	}
	return dp[n]
}
