package n5925

func countPyramids(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// from down to top
	dp[m-1] = grid[m-1]
	for i := m - 2; i >= 0; i-- {
		dp[i][0] = grid[i][0]
		dp[i][n-1] = grid[i][n-1]
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(min(dp[i+1][j-1]+1, dp[i+1][j+1]+1), dp[i+1][j]+1)
			}
			if dp[i][j] > 0 {
				ans += dp[i][j] - 1
			}
		}
	}

	// from top to down
	dp[0] = grid[0]
	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0]
		dp[i][n-1] = grid[i][n-1]
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = min(min(dp[i-1][j-1]+1, dp[i-1][j+1]+1), dp[i-1][j]+1)
			}
			if dp[i][j] > 0 {
				ans += dp[i][j] - 1
			}
		}
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
