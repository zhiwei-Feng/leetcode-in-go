package n0096

func numTrees(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		for root := 1; root <= i; root++ {
			leftn := root - 1
			rightn := i - root
			dp[i] += dp[leftn] * dp[rightn]
		}
	}

	return dp[n]
}
