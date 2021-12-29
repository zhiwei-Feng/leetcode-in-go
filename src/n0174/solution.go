package n0174

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	if dungeon[m-1][n-1] > 0 {
		dp[m-1][n-1] = 1
	} else {
		dp[m-1][n-1] = 1 - dungeon[m-1][n-1]
	}

	for i := m - 2; i >= 0; i-- {
		dp[i][n-1] = max(dp[i+1][n-1]-dungeon[i][n-1], 1)
	}
	for i := n - 2; i >= 0; i-- {
		dp[m-1][i] = max(dp[m-1][i+1]-dungeon[m-1][i], 1)
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = max(min(dp[i+1][j], dp[i][j+1])-dungeon[i][j], 1)
		}
	}

	return dp[0][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
