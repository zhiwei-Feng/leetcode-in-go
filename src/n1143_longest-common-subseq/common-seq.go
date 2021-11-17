package n1143

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	count := 0 // init dp[i][0], record if text2[0] is existed in text1 
	for i := 0; i < m; i++ {
		if text1[i] == text2[0] {
			count++
		}
		if count > 0 {
			dp[i][0] = 1
		}
	}
	count = 0 // init dp[0][j]
	for j := 0; j < n; j++ {
		if text1[0] == text2[j] {
			count++
		}
		if count > 0 {
			dp[0][j] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}

func max(x1, x2 int) int {
	if x1 > x2 {
		return x1
	}
	return x2
}
