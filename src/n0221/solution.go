package n0221

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	ans := 0
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)

	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
			ans = 1
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ans = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
			}
		}
	}

	return ans * ans
}

func min(x1, x2, x3 int) int {
	minV := x1
	if x2 < minV {
		minV = x2
	}
	if x3 < minV {
		minV = x3
	}
	return minV
}
