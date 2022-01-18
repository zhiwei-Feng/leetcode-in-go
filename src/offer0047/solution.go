package offer0047

func maxValue(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)

	}
	// init
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	// iter
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}